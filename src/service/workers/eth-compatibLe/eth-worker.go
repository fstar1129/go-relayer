package eth

// import (
// 	"bytes"
// 	"context"
// 	"fmt"
// 	"math/big"
// 	"strings"
// 	"time"

// 	"github.com/ethereum/go-ethereum"
// 	"github.com/ethereum/go-ethereum/accounts/abi"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/sirupsen/logrus"

// 	"gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/models"
// 	"gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/storage"
// 	da "gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/workers/eth-compatible/abi/erc20Swap"
// 	"gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/workers/utils"
// )

// // Worker ...
// type Worker struct {
// 	provider         string
// 	chainID          string
// 	logger           *logrus.Logger // logger
// 	abi              abi.ABI
// 	Config           *models.WorkerConfig
// 	client           *ethclient.Client
// 	swapContractAddr common.Address
// }

// // NewEthWorker ...
// func NewEthWorker(logger *logrus.Logger, provider string, contractAddress common.Address, cfg *models.WorkerConfig) *Worker {
// 	contractabi, err := abi.JSON(strings.NewReader(da.ETHSwapABI))
// 	if err != nil {
// 		panic(fmt.Sprintf("marshal abi error, err=%s", err.Error()))
// 	}

// 	client, err := ethclient.Dial(provider)
// 	if err != nil {
// 		panic(fmt.Sprintf("new eth client error, err=%s", err.Error()))
// 	}

// 	return &Worker{
// 		logger:           logger.WithField("worker", cfg.ChainID).Logger,
// 		provider:         provider,
// 		abi:              contractabi,
// 		client:           client,
// 		swapContractAddr: contractAddress,
// 	}
// }

// // GetChain ...
// func (w *Worker) GetChain() string {
// 	return w.chainID
// }

// // GetStatus returns status of relayer account(balance eg)
// func (w *Worker) GetStatus() (interface{}, error) {
// 	ethStatus := &EthStatus{}

// 	ethBalance, err := w.EthBalance(w.Config.WorkerAddr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ethStatus.EthBalance = utils.QuoBigInt(ethBalance, utils.GetBigIntForDecimal(18)).String()

// 	return ethStatus, nil
// }

// // GetBlockAndTxs ...
// func (w *Worker) GetBlockAndTxs(height int64) (*models.BlockAndTxLogs, error) {
// 	header, err := w.client.HeaderByNumber(context.Background(), big.NewInt(height))
// 	if err != nil {
// 		return nil, err
// 	}

// 	txLogs, err := w.GetLogs(header.Hash())
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &models.BlockAndTxLogs{
// 		Height:          height,
// 		BlockHash:       header.Hash().String(),
// 		ParentBlockHash: header.ParentHash.String(),
// 		BlockTime:       int64(header.Time),
// 		TxLogs:          txLogs,
// 	}, nil
// }

// // GetLogs ...
// func (w *Worker) GetLogs(blockHash common.Hash) ([]*storage.TxLog, error) {
// 	topics := [][]common.Hash{{ClaimedEventHash, HTLTEventHash, RefundEventHash}}

// 	logs, err := w.client.FilterLogs(context.Background(), ethereum.FilterQuery{
// 		BlockHash: &blockHash,
// 		Topics:    topics,
// 		Addresses: []common.Address{w.swapContractAddr},
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	models := make([]*storage.TxLog, 0, len(logs))
// 	for _, log := range logs {
// 		event, err := ParseEvent(&w.abi, &log)
// 		if err != nil {
// 			w.logger.WithFields(logrus.Fields{"function": "GetLogs()"}).Errorln(err)
// 			continue
// 		}
// 		if event == nil {
// 			continue
// 		}

// 		txLog := event.ToTxLog()
// 		//	txLog.ContractAddr = log.Address.Hex()
// 		txLog.Height = int64(log.BlockNumber)
// 		txLog.BlockHash = log.BlockHash.Hex()
// 		txLog.TxHash = log.TxHash.Hex()
// 		txLog.Status = storage.TxStatusInit
// 		models = append(models, txLog)
// 	}
// 	return models, nil
// }

// // GetFetchInterval ...
// func (w *Worker) GetFetchInterval() time.Duration {
// 	return time.Duration(w.Config.FetchInterval) * time.Second
// }

// // GetHeight ...
// func (w *Worker) GetHeight() (int64, error) {
// 	header, err := w.client.HeaderByNumber(context.Background(), nil)
// 	if err != nil {
// 		return 0, nil
// 	}

// 	return header.Number.Int64(), nil
// }

// // HTLT ...
// func (w *Worker) HTLT(randomNumberHash common.Hash, timestamp int64, heightSpan int64, recipientAddr string,
// 	otherChainSenderAddr string, otherChainRecipientAddr string, outAmount *big.Int) (string, error) {
// 	auth, err := w.GetTransactor()
// 	if err != nil {
// 		return "", err
// 	}

// 	instance, err := da.NewETHSwap(w.swapContractAddr, w.client)
// 	if err != nil {
// 		return "", err
// 	}

// 	recvAddr := common.HexToAddress(recipientAddr)
// 	// bep2RecipientAddr, err := bt.AccAddressFromBech32(otherChainRecipientAddr)
// 	// if err != nil {
// 	// 	return "", err
// 	// }

// 	// var bep2SenderAddr bt.AccAddress
// 	// if otherChainSenderAddr != "" {
// 	// 	bep2SenderAddr, err = bt.AccAddressFromBech32(otherChainSenderAddr)
// 	// 	if err != nil {
// 	// 		return "", err
// 	// 	}
// 	// }

// 	//auth.From = w.Config.RelayerAddr
// 	auth.Value = outAmount

// 	tx, err := instance.Htlt(auth, randomNumberHash, uint64(timestamp), big.NewInt(heightSpan), recvAddr,
// 		common.BytesToAddress(nil), common.BytesToAddress(nil), big.NewInt(0))
// 	if err != nil {
// 		return "", err
// 	}
// 	w.logger.Debugf("init tx sent: %s", tx.Hash().Hex())

// 	return tx.Hash().String(), nil
// }

// // Claim ...
// func (w *Worker) Claim(swapID common.Hash, randomNumber common.Hash) (string, error) {
// 	auth, err := w.GetTransactor()
// 	if err != nil {
// 		return "", err
// 	}

// 	instance, err := da.NewETHSwap(w.swapContractAddr, w.client)
// 	if err != nil {
// 		return "", err
// 	}

// 	tx, err := instance.Claim(auth, swapID, randomNumber)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tx.Hash().String(), nil
// }

// // Refund ...
// func (w *Worker) Refund(swapID common.Hash) (string, error) {
// 	auth, err := w.GetTransactor()
// 	if err != nil {
// 		return "", err
// 	}

// 	instance, err := da.NewETHSwap(w.swapContractAddr, w.client)
// 	if err != nil {
// 		return "", err
// 	}

// 	tx, err := instance.Refund(auth, swapID)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tx.Hash().String(), nil
// }

// // GetSentTxStatus ...
// func (w *Worker) GetSentTxStatus(hash string) storage.TxStatus {
// 	_, isPending, err := w.client.TransactionByHash(context.Background(), common.HexToHash(hash))
// 	if err != nil {
// 		return storage.TxSentStatusNotFound
// 	}
// 	if isPending {
// 		return storage.TxSentStatusPending
// 	}

// 	txReceipt, err := w.client.TransactionReceipt(context.Background(), common.HexToHash(hash))
// 	if err != nil {
// 		return storage.TxSentStatusNotFound
// 	}

// 	if txReceipt.Status == types.ReceiptStatusFailed {
// 		return storage.TxSentStatusFailed
// 	}

// 	return storage.TxSentStatusSuccess
// }

// // HasSwap ...
// func (w *Worker) HasSwap(swapID common.Hash) (bool, error) {
// 	instance, err := da.NewETHSwap(w.swapContractAddr, w.client)
// 	if err != nil {
// 		return false, err
// 	}

// 	return instance.IsSwapExist(nil, swapID)
// }

// // // GetHTLTEvent ...
// // func (w *Worker) GetHTLTEvent(swapID common.Hash) (*HTLTEvent, error) {
// // 	topics := [][]common.Hash{{HTLTEventHash}, {}, {}, {swapID}}
// // 	logs, err := w.client.FilterLogs(context.Background(), ethereum.FilterQuery{
// // 		Topics: topics,
// // 	})
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	if len(logs) == 0 {
// // 		return nil, fmt.Errorf("swap id does not exist, swap_id=%s", swapID.String())
// // 	}

// // 	event, err := ParseHTLTEvent(&w.abi, &logs[0])
// // 	if err != nil {
// // 		w.logger.Errorf("parse event log error, er=%s", err.Error())
// // 		return nil, err
// // 	}

// // 	htltEvent := event.(HTLTEvent)
// // 	return &htltEvent, nil
// // }

// // // GetSwap ...
// // func (w *Worker) GetSwap(swapID common.Hash) (*models.SwapRequest, error) {
// // 	htltEvent, err := w.GetHTLTEvent(swapID)
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	return &models.SwapRequest{
// // 		ID:               swapID,
// // 		ExpireHeight:     htltEvent.ExpireHeight.Int64(),
// // 		SenderAddress:    htltEvent.MsgSender.String(),
// // 		RecipientAddress: htltEvent.LaRecipientAddr.String(),
// // 		OutAmount:        htltEvent.OutAmount,
// // 	}, nil
// // }

// // Refundable ...
// func (w *Worker) Refundable(swapID common.Hash) (bool, error) {
// 	instance, err := da.NewETHSwap(w.swapContractAddr, w.client)
// 	if err != nil {
// 		return false, err
// 	}

// 	refundable, err := instance.Refundable(nil, swapID)
// 	return refundable, err
// }

// // Claimable ...
// func (w *Worker) Claimable(swapID common.Hash) (bool, error) {
// 	instance, err := da.NewETHSwap(w.swapContractAddr, w.client)
// 	if err != nil {
// 		return false, err
// 	}

// 	claimable, err := instance.Claimable(nil, swapID)
// 	return claimable, err
// }

// // GetTransactor ...
// func (w *Worker) GetTransactor() (*bind.TransactOpts, error) {
// 	privateKey, err := utils.GetPrivateKey(w.Config)
// 	if err != nil {
// 		return nil, err
// 	}

// 	nonce, err := w.client.PendingNonceAt(context.Background(), w.Config.WorkerAddr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	auth := bind.NewKeyedTransactor(privateKey)
// 	auth.Nonce = big.NewInt(int64(nonce))
// 	auth.Value = big.NewInt(0)                // in wei
// 	auth.GasLimit = uint64(w.Config.GasLimit) // in units
// 	auth.GasPrice = w.Config.GasPrice
// 	return auth, nil
// }

// // GetBalance ...
// func (w *Worker) GetBalance(addressString string) (*big.Int, error) {
// 	address := common.HexToAddress(addressString)
// 	return w.EthBalance(address)
// }

// // EthBalance ...
// func (w *Worker) EthBalance(address common.Address) (*big.Int, error) {
// 	return w.client.BalanceAt(context.Background(), address, nil)
// }

// // GetRelayerAddress ...
// func (w *Worker) GetRelayerAddress() string {
// 	return w.Config.WorkerAddr.String()
// }

// // GetColdWalletAddress ...
// func (w *Worker) GetColdWalletAddress() string {
// 	return w.Config.ColdWalletAddr.String()
// }

// // CalcSwapID ...
// func (w *Worker) CalcSwapID(randomNumberHash common.Hash, sender string, senderOtherChain string) ([]byte, error) {
// 	// var bep2Addr = bt.AccAddress{}
// 	// if senderOtherChain != "" {
// 	// 	parsedAddr, err := bt.AccAddressFromBech32(senderOtherChain)
// 	// 	if err != nil {
// 	// 		return nil, err
// 	// 	}

// 	// 	bep2Addr = parsedAddr
// 	// }

// 	//return CalculateSwapID(randomNumberHash[:], common.FromHex(sender), nil), nil
// }

// // IsSameAddress ...
// func (w *Worker) IsSameAddress(addrA string, addrB string) bool {
// 	return bytes.Equal(common.FromHex(addrA), common.FromHex(addrB))
// }

// // GetBalanceAlertMsg ...
// func (w *Worker) GetBalanceAlertMsg() (string, error) {
// 	if w.Config.EthBalanceAlertThreshold.Cmp(big.NewInt(0)) == 0 {
// 		return "", nil
// 	}

// 	alertMsg := ""
// 	if w.Config.EthBalanceAlertThreshold.Cmp(big.NewInt(0)) > 0 {
// 		ethBalance, err := w.EthBalance(w.Config.WorkerAddr)
// 		if err != nil {
// 			return "", err
// 		}

// 		if ethBalance.Cmp(w.Config.EthBalanceAlertThreshold) < 0 {
// 			alertMsg = alertMsg + fmt.Sprintf("eth balance(%s) is less than %s",
// 				ethBalance.String(), w.Config.EthBalanceAlertThreshold.String())
// 		}
// 	}
// 	return alertMsg, nil
// }

// // SendAmount ...
// func (w *Worker) SendAmount(address string, amount *big.Int) (string, error) {
// 	return "", fmt.Errorf("not implemented") // TODO
// }
