package rlr

import (
	"fmt"
	watcher "latoken/relayer-smart-contract/src/service/blockchains-watcher"
	"latoken/relayer-smart-contract/src/service/storage"
	workers "latoken/relayer-smart-contract/src/service/workers"
	"latoken/relayer-smart-contract/src/service/workers/eth-compatibLe"
	"math/big"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"latoken/relayer-smart-contract/src/models"
)

var laChainURL = "http://localhost:7545"

var (
	WorkerChainRatio            *big.Float = &big.Float{}
	BnbMinRemainHeight          int64      = 6
	WorkerChainExpireHeightSpan int64      = 7
	WorkerChainMaxrOutAmount    *big.Int   = &big.Int{}
	WorkerChainDecimal                     = 2

	WorkerChainConfirmNum = 8
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
func CreateNewRelayerSRV(logger *logrus.Logger, gormDB *gorm.DB, laConfig, ethConfig *models.WorkerConfig) *RelayerSRV {
	// init database
	db, err := storage.InitStorage(gormDB)
	if err != nil {
		logger.Fatalf("Connect to DataBase: ", err)
	}

	// create Relayer instance
	inst := RelayerSRV{
		logger:   logger,
		storage:  db,
		laWorker: eth.NewErc20Worker(logrus.StandardLogger(), laConfig),
		Workers:  make(map[string]workers.IWorker),
	}
	// create erc20 worker
	inst.Workers["ERC20"] = eth.NewErc20Worker(logrus.StandardLogger(), ethConfig)
	// // create la worker
	inst.Workers[storage.LaChain] = inst.laWorker

	// check rules for workers(>=2, different chainIDs...)
	if len(inst.Workers) < 1 {
		logger.Fatalf("Num of workers must be > 1, but = %d", len(inst.Workers))
		return nil
	}
	inst.Watcher = watcher.CreateNewWatcherSRV(logger, db, inst.Workers)

	return &inst
}

// Run ...
func (r *RelayerSRV) Run() {
	// start watcher
	r.Watcher.Run()
	go r.emitChainSendClaim(storage.SwapTypeUnbind)
	go r.emitChainSendClaim(storage.SwapTypeBind)
	// // run Worker workers
	for _, worker := range r.Workers {
		go r.ConfirmWorkerTx(worker)
		go r.emitChainSendSpend(worker)
		go r.CheckTxSentRoutine(worker)
	}
}

// ConfirmWorkerTx ...
func (r *RelayerSRV) ConfirmWorkerTx(worker workers.IWorker) {
	for {
		txLogs, err := r.storage.FindTxLogs(worker.GetChain(), worker.GetConfirmNum())
		if err != nil {
			r.logger.Errorf("ConfirmWorkerTx(), err = %s", err)
			time.Sleep(10 * time.Second)
			continue
		}

		txHashes := make([]string, 0, len(txLogs))
		newSwaps := make([]*storage.Swap, 0)

		// CREATE NEW SWAPS
		for _, txLog := range txLogs {
			if txLog.TxType == storage.TxTypeDeposit {
				swapType := storage.SwapTypeBind
				if txLog.Chain == storage.LaChain {
					swapType = storage.SwapTypeUnbind
				}
				// reject swap request if receiver addr and worker chain addr both are r addr
				if worker.IsSameAddress(txLog.ReceiverAddr, worker.GetWorkerAddress()) &&
					!r.laWorker.IsSameAddress(txLog.WorkerChainAddr, r.laWorker.GetWorkerAddress()) {

					fmt.Println("THE SAME")
				}

				// CHECK DATA HASH

				// randomNumberHash := common.HexToHash(txLog.RandomNumberHash)
				// laChainSwapID, err := r.laWorker.CalcSwapID(randomNumberHash, r.laWorker.GetWorkerAddress(), txLog.SenderAddr)
				// if err != nil {
				// 	r.logger.Errorf("calculate swap id parse error, random_number_hash=%s, sender_addr=%s, Worker_swap_id=%s", randomNumberHash.String(), txLog.SenderAddr, txLog.SwapID)
				// 	continue
				// }
				fmt.Println("NEW!!!")
				newSwap := &storage.Swap{
					Type:           swapType,
					SwapID:         txLog.SwapID,
					SenderAddr:     txLog.SenderAddr,
					ReceiverAddr:   txLog.ReceiverAddr,
					ERC20TokenAddr: txLog.ERC20TokenAddr,
					DepositNonce:   txLog.DepositNonce,
					ResourceID:     txLog.ResourceID,
					DataHash:       txLog.DataHash,
					СhainID:        txLog.DestinationChainID,
					Height:         txLog.Height,
					Status:         storage.SwapStatusDepositConfirmed,
					CreateTime:     time.Now().Unix(),
				}
				newSwaps = append(newSwaps, newSwap)
			}
			txHashes = append(txHashes, txLog.TxHash)
		}

		//
		if err := r.storage.ConfirmWorkerTx(worker.GetChain(), txLogs, txHashes, newSwaps); err != nil {
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
	txsSent, err := r.storage.GetTxsSentByStatus(worker.GetChain())
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
		r.storage.UpdateSwapStatus(swap, backwardStatus, swap.DataHash, "")
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

		fmt.Printf("timeElapsed(%d) | autoRetryTimeout(%d) \n", timeElapsed, autoRetryTimeout)
		if len(txsSent) >= autoRetryNum {
			r.storage.UpdateSwapStatus(swap, failedStatus, "", "")
		} else {
			r.storage.UpdateSwapStatus(swap, backwardStatus, "", "")
		}
		r.storage.UpdateTxSentStatus(latestTx, storage.TxSentStatusLost)
	} else if txStatus == storage.TxSentStatusFailed {
		r.storage.UpdateSwapStatus(swap, failedStatus, "", "")
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

	return 100000, 1000
}
