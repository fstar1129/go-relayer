package rlr

import (
	"sync"
	"time"

	watcher "github.com/LATOKEN/relayer-smart-contract.git/src/service/blockchains-watcher"
	"github.com/LATOKEN/relayer-smart-contract.git/src/service/storage"
	workers "github.com/LATOKEN/relayer-smart-contract.git/src/service/workers"
	"github.com/LATOKEN/relayer-smart-contract.git/src/service/workers/eth-compatible"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"github.com/LATOKEN/relayer-smart-contract.git/src/models"
)

// RelayerSRV ...
type RelayerSRV struct {
	sync.RWMutex
	logger   *logrus.Logger
	Watcher  *watcher.WatcherSRV
	laWorker workers.IWorker
	Workers  map[string]workers.IWorker
	storage  *storage.DataBase
}

// CreateNewRelayerSRV ...
func CreateNewRelayerSRV(logger *logrus.Logger, gormDB *gorm.DB, laConfig *models.WorkerConfig,
	chainCfgs []*models.WorkerConfig, resourceIDs []*storage.ResourceId) *RelayerSRV {
	// init database
	db, err := storage.InitStorage(gormDB)
	if err != nil {
		logger.Fatalf("Connect to DataBase: ", err)
	}

	// create Relayer instance
	inst := RelayerSRV{
		logger:   logger,
		storage:  db,
		laWorker: eth.NewErc20Worker(logger, laConfig, db),
		Workers:  make(map[string]workers.IWorker),
	}
	// create erc20 worker
	for _, cfg := range chainCfgs {
		inst.Workers[cfg.ChainName] = eth.NewErc20Worker(logger, cfg, db)
	}
	// // create la worker
	inst.Workers["LA"] = inst.laWorker

	// check rules for workers(>=2, different chainIDs...)
	if len(inst.Workers) < 1 {
		logger.Fatalf("Num of workers must be > 1, but = %d", len(inst.Workers))
		return nil
	}
	inst.Watcher = watcher.CreateNewWatcherSRV(logger, db, inst.Workers)
	db.SaveResourceIDs(resourceIDs)
	return &inst
}

// Run ...
func (r *RelayerSRV) Run() {
	// start watcher
	r.Watcher.Run()
	go r.emitChainSendClaim()
	go r.emitChainSendPass()
	go r.emitChainSendSpend()
	go r.emitChainSendExpire()
	// run Worker workers
	for _, worker := range r.Workers {
		go r.ConfirmWorkerTx(worker)
		go r.CheckTxSentRoutine(worker)
	}
}

func (r *RelayerSRV) GetSwapStatus(req *models.SwapStatus) (storage.SwapStatus, error) {
	swapType := storage.SwapTypeUnbind
	if req.Chain == "LA" {
		swapType = storage.SwapTypeBind
	}
	swap, err := r.storage.GetSwapByStatus(swapType, req.Sender, req.Receipt, req.Amount, req.TxHash)
	if err != nil {
		r.logger.Errorf("GetSwapByStatus type %s, req: %v, failed with error: %v", swapType, req, err)
		return "", err
	}
	if swap != nil {
		return swap.Status, nil
	}

	return "", nil
}

// Status ...
func (r *RelayerSRV) StatusOfWorkers() (map[string]*models.WorkerStatus, error) {
	// get blockchain heights from workers and from database
	workers := make(map[string]*models.WorkerStatus)
	for _, w := range r.Workers {
		status, err := w.GetStatus()
		if err != nil {
			r.logger.Errorf("While get status for worker = %s, err = %v", w.GetChainName(), err)
			return nil, err
		}
		workers[w.GetChainName()] = status
	}

	for name, w := range workers {
		blocks := r.storage.GetCurrentBlockLog(name)
		w.SyncHeight = blocks.Height
	}

	return workers, nil
}

// ConfirmWorkerTx ...
func (r *RelayerSRV) ConfirmWorkerTx(worker workers.IWorker) {
	for {
		txLogs, err := r.storage.FindTxLogs(worker.GetChainName(), worker.GetConfirmNum())
		if err != nil {
			r.logger.Errorf("ConfirmWorkerTx(), err = %s", err)
			time.Sleep(10 * time.Second)
			continue
		}

		txHashes := make([]string, 0, len(txLogs))
		newSwaps := make([]*storage.Swap, 0)

		// CREATE NEW SWAPS
		for _, txLog := range txLogs {
			if txLog.Status == storage.TxStatusInit {
				swapType := storage.SwapTypeUnbind
				if txLog.DestinationChainID == r.laWorker.GetDestinationID() {
					swapType = storage.SwapTypeBind
				}
				// reject swap request if receiver addr and worker chain addr both are r addr
				// if worker.IsSameAddress(txLog.ReceiverAddr, worker.GetWorkerAddress()) &&
				// 	!r.laWorker.IsSameAddress(txLog.WorkerChainAddr, r.laWorker.GetWorkerAddress()) {
				// 	r.logger.Warnln("THE SAME")
				// }

				r.logger.Infoln("NEW SWAP")
				newSwap := &storage.Swap{
					Type:               swapType,
					SwapID:             txLog.SwapID,
					SenderAddr:         txLog.SenderAddr,
					ReceiverAddr:       txLog.ReceiverAddr,
					InTokenAddr:        txLog.InTokenAddr,
					DepositNonce:       txLog.DepositNonce,
					ResourceID:         txLog.ResourceID,
					OutAmount:          txLog.OutAmount,
					DestinationChainID: txLog.DestinationChainID,
					OriginChainID:      txLog.OriginСhainID,
					Height:             txLog.Height,
					Status:             txLog.SwapStatus,
					CreateTime:         time.Now().Unix(),
				}
				if txLog.TxType == storage.TxTypeDeposit {
					newSwap.TxHash = txLog.TxHash
				}
				newSwaps = append(newSwaps, newSwap)
				txHashes = append(txHashes, txLog.TxHash)
				r.logger.Infof("compensate new swap tx coplete, tx %v, logs: %v, swaps: %v", txHashes, txLogs, newSwaps)
			}
		}

		//stores new txLogs in db
		if err := r.storage.ConfirmWorkerTx(worker.GetChainName(), txLogs, txHashes, newSwaps); err != nil {
			r.logger.Errorf("compensate new swap tx error, err=%s", err)
		}

		time.Sleep(2 * time.Second)
	}
}

// CheckTxSentRoutine ...
func (r *RelayerSRV) CheckTxSentRoutine(worker workers.IWorker) {
	for {
		r.CheckTxSent(worker)
		time.Sleep(time.Second)
	}
}

// CheckTxSent ...
func (r *RelayerSRV) CheckTxSent(worker workers.IWorker) {
	txsSent, err := r.storage.GetTxsSentByStatus(worker.GetChainName())
	if err != nil {
		r.logger.WithFields(logrus.Fields{"function": "CheckTxSent() | GetTxsSentByStatus()"}).Errorln(err)
		return
	}

	for _, txSent := range txsSent {
		// Get status of tx from chain
		status := worker.GetSentTxStatus(txSent.TxHash)
		if err := r.storage.UpdateTxSentStatus(txSent, status); err != nil {
			r.logger.WithFields(logrus.Fields{"function": "CheckTxSent() | UpdateTxSentStatus()"}).Errorln(err)
			return
		}
	}
}

func (r *RelayerSRV) handleTxSent(chain string, swap *storage.Swap, txType storage.TxType, backwardStatus storage.SwapStatus,
	failedStatus storage.SwapStatus) {
	txsSent := r.storage.GetTxsSentByType(chain, txType, swap)
	if len(txsSent) == 0 {
		r.storage.UpdateSwapStatus(swap, backwardStatus, "")
		return
	}
	latestTx := txsSent[0]
	timeElapsed := time.Now().Unix() - latestTx.CreateTime
	autoRetryTimeout, autoRetryNum := r.getAutoRetryConfig(chain)
	txStatus := latestTx.Status

	if timeElapsed > autoRetryTimeout &&
		(txStatus == storage.TxSentStatusNotFound ||
			txStatus == storage.TxSentStatusInit ||
			txStatus == storage.TxSentStatusPending) {

		r.logger.Warnln("timeElapsed(%d) | autoRetryTimeout(%d) \n", timeElapsed, autoRetryTimeout)
		if len(txsSent) >= autoRetryNum {
			r.storage.UpdateSwapStatus(swap, failedStatus, "")
		} else {
			r.storage.UpdateSwapStatus(swap, backwardStatus, "")
		}
		r.storage.UpdateTxSentStatus(latestTx, storage.TxSentStatusLost)
	} else if txStatus == storage.TxSentStatusFailed {
		r.storage.UpdateSwapStatus(swap, failedStatus, "")
	}
}

func (r *RelayerSRV) getAutoRetryConfig(chain string) (int64, int) {
	// if chain == "LA" {
	// 	autoRetryTimeout = r.Config.ChainConfig.BnbAutoRetryTimeout
	// 	autoRetryNum = r.Config.ChainConfig.BnbAutoRetryNum
	// } else {
	// 	autoRetryTimeout = r.Config.ChainConfig.WorkerChainAutoRetryTimeout
	// 	autoRetryNum = r.Config.ChainConfig.WorkerChainAutoRetryNum
	// }

	return 1000, 10
}
