package eth

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/LATOKEN/relayer-smart-contract.git/src/service/workers/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	jsonrpc "github.com/ybbus/jsonrpc/v2"

	"github.com/LATOKEN/relayer-smart-contract.git/src/models"
	"github.com/LATOKEN/relayer-smart-contract.git/src/service/storage"
	labr "github.com/LATOKEN/relayer-smart-contract.git/src/service/workers/eth-compatible/abi/bridge/la"
	es "github.com/LATOKEN/relayer-smart-contract.git/src/service/workers/eth-compatible/abi/erc20Swap"
)

// Erc20Worker ...
type Erc20Worker struct {
	provider           string
	chainName          string
	chainID            int64
	destinationChainID string
	logger             *logrus.Entry // logger
	config             *models.WorkerConfig
	client             *ethclient.Client
	swapContractAddr   common.Address
	db                 *storage.DataBase
}

// NewErc20Worker ...
func NewErc20Worker(logger *logrus.Logger, cfg *models.WorkerConfig, db *storage.DataBase) *Erc20Worker {
	client, err := ethclient.Dial(cfg.Provider)
	if err != nil {
		panic(fmt.Sprintf("rpc err for chain %s: %s", cfg.ChainName, err.Error()))
	}

	privKey, err := utils.GetPrivateKey(cfg)
	if err != nil {
		panic(fmt.Sprintf("generate private key error, err=%s", err.Error()))
	}

	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic(fmt.Sprintf("get public key error, err=%s", err.Error()))
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	if !bytes.Equal(cfg.WorkerAddr.Bytes(), fromAddress.Bytes()) {
		panic(fmt.Sprintf(
			"relayer address supplied in config (%s) does not match mnemonic (%s)",
			cfg.WorkerAddr, fromAddress,
		))
	}
	chainid, err := client.ChainID(context.Background())
	if err != nil {
		panic("cannot get chain id from rpc")
	}

	// init token addresses
	return &Erc20Worker{
		chainName:          cfg.ChainName,
		chainID:            chainid.Int64(),
		destinationChainID: cfg.DestinationChainID,
		logger:             logger.WithField("worker", cfg.ChainName),
		provider:           cfg.Provider,
		config:             cfg,
		client:             client,
		swapContractAddr:   cfg.ContractAddr,
		db:                 db,
	}
}

// GetChainName returns chain name
func (w *Erc20Worker) GetChainName() string {
	return w.chainName
}

// GetChainID returns chain id
func (w *Erc20Worker) GetChainID() int64 {
	return w.chainID
}

func (w *Erc20Worker) GetDestinationID() string {
	return w.destinationChainID
}

// // GetChain returns chain name
// func (w *Erc20Worker) GetChain() string {
// 	return w.chainName
// }
// GetStartHeight returns start blockchain height from config
func (w *Erc20Worker) GetStartHeight() (int64, error) {
	if w.config.StartBlockHeight == 0 {
		return w.GetHeight()
	}
	return w.config.StartBlockHeight, nil
}

// GetConfirmNum returns numbers of blocks after them tx will be confirmed
func (w *Erc20Worker) GetConfirmNum() int64 {
	return w.config.ConfirmNum
}

// GetStatus returns status of relayer: blockchain; account(address, balance ...)
func (w *Erc20Worker) GetStatus() (*models.WorkerStatus, error) {
	status := &models.WorkerStatus{}
	// set current block height
	height, err := w.GetHeight()
	if err != nil {
		return nil, err
	}
	status.Height = height
	// set worker address
	status.Account.Address = w.config.WorkerAddr.Hex()

	return status, nil
}

// GetBlockAndTxs ...
func (w *Erc20Worker) GetBlockAndTxs(height int64) (*models.BlockAndTxLogs, error) {
	// var head *Header
	// rpcClient := jsonrpc.NewClient(w.provider)

	// resp, err := rpcClient.Call("eth_getBlockByNumber", "latest", false)
	// if err != nil {
	// 	w.logger.Errorln("while call eth_getBlockByNumber, err = ", err)
	// 	return nil, err
	// }

	// if err := resp.GetObject(&head); err != nil {
	// 	w.logger.Errorln("while GetObject, err = ", err)
	// 	return nil, err
	// }

	// if head == nil {
	// 	return nil, fmt.Errorf("not found")
	// }

	// if height >= int64(head.Number) {
	// 	return nil, fmt.Errorf("not found")
	// }

	// logs, err := w.getLogs(height, int64(head.Number))
	// if err != nil {
	// 	w.logger.Errorf("while getEvents(from = %d, to = %d), err = %v", height, int64(head.Number), err)
	// 	return nil, err
	// }

	// return &models.BlockAndTxLogs{
	// 	Height:          int64(head.Number),
	// 	BlockHash:       head.Hash.String(),
	// 	ParentBlockHash: head.ParentHash.Hex(),
	// 	BlockTime:       int64(head.Time),
	// 	TxLogs:          logs,
	// }, nil

	client, err := ethclient.Dial(w.provider)
	if err != nil {
		w.logger.Errorln("Error while dialing the client = ", err)
		return nil, err
	}

	clientResp, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		w.logger.Errorln("Error while fetching the block header = ", err)
		return nil, err

	}

	if height >= clientResp.Number.Int64() {
		return nil, fmt.Errorf("not found")
	}

	logs, err := w.getLogs(height, clientResp.Number.Int64())
	if err != nil {
		w.logger.Errorf("while getEvents(block number from %d to %d), err = %v", height, clientResp.Number, err)
		return nil, err
	}

	client.Close()

	return &models.BlockAndTxLogs{
		Height:          clientResp.Number.Int64(),
		BlockHash:       clientResp.Hash().String(),
		ParentBlockHash: clientResp.ParentHash.Hex(),
		BlockTime:       int64(clientResp.Time),
		TxLogs:          logs,
	}, nil
}

// GetFetchInterval ...
func (w *Erc20Worker) GetFetchInterval() time.Duration {
	return time.Duration(w.config.FetchInterval) * time.Second
}

// getLogs ...
func (w *Erc20Worker) getLogs(curHeight, nextHeight int64) ([]*storage.TxLog, error) {
	if curHeight == 0 || nextHeight-curHeight > 3500 {
		curHeight = nextHeight - 1
	}
	logs, err := w.client.FilterLogs(context.Background(), ethereum.FilterQuery{
		// BlockHash: &blockHash,
		FromBlock: big.NewInt(curHeight + 1),
		ToBlock:   big.NewInt(nextHeight),
		Addresses: []common.Address{w.swapContractAddr},
		// Topics:    [][]common.Hash{},
	})
	if err != nil {
		w.logger.WithFields(logrus.Fields{"function": "GetLogs()"}).Errorf("get event log error, err=%s", err)
		return nil, err
	}

	models := make([]*storage.TxLog, 0, len(logs))
	for _, log := range logs {
		w.logger.Infof("WORKER(%s) NEW EVENT: %v\n\n", w.chainID, log)
		event, err := w.parseEvent(&log)
		if err != nil {
			w.logger.WithFields(logrus.Fields{"function": "GetLogs()"}).Errorf("parse event log error, err=%s", err)
			continue
		}
		if event == nil {
			continue
		}

		txLog := event.ToTxLog()

		// 	w.logger.Infof("Deposited\ndestination chain ID: 0x%s\nresource ID: 0x%s\ndeposit nonce: %d\nrecipient address: %s\namount address: %s\ntoken address: %s\n",
		// 		txLog.DestinationChainID, txLog.ResourceID, txLog.DepositNonce, txLog.SenderAddr, txLog.ReceiverAddr, txLog.OutAmount, txLog.InTokenAddr)
		// } else if txLog.TxType == storage.TxTypeClaim {
		// 	w.logger.Infof("\nProposalEvent:\nOrigin chain ID: 0x%s\ndeposit nonce: %d\ndeposit nonce: %v\nstatus: %v\nresource ID: 0x%s\n",
		// 		txLog.OriginСhainID, txLog.DepositNonce, txLog.SwapStatus, txLog.ResourceID, txLog.ReceiverAddr)
		// }
		txLog.Chain = w.chainName
		txLog.Height = int64(log.BlockNumber)
		txLog.BlockHash = log.BlockHash.Hex()
		txLog.TxHash = log.TxHash.Hex()
		txLog.Status = storage.TxStatusInit

		models = append(models, txLog)
	}
	return models, nil
}

// GetHeight ..
func (w *Erc20Worker) GetHeight() (int64, error) {
	header, err := w.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, nil
	}
	return header.Number.Int64(), nil
}

// Vote ...
func (w *Erc20Worker) Vote(depositNonce uint64, originchainID [8]byte, destinationChainID [8]byte, resourceID [32]byte, receiptAddr string, amount string) (string, error) {
	auth, err := w.getTransactor()
	if err != nil {
		return "", err
	}

	instance, err := labr.NewLabr(w.swapContractAddr, w.client)
	if err != nil {
		return "", err
	}

	value, _ := new(big.Int).SetString(amount, 10)
	tx, err := instance.VoteProposal(auth, originchainID, destinationChainID, depositNonce, resourceID, common.HexToAddress(receiptAddr), value)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

// HasSwap ...
func (w *Erc20Worker) HasSwap(swapID common.Hash) (bool, error) {
	instance, err := es.NewERC20Swap(w.swapContractAddr, w.client)
	if err != nil {
		return false, err
	}

	return instance.IsSwapExist(nil, swapID)
}

// Refundable ...
func (w *Erc20Worker) Refundable(swapID common.Hash) (bool, error) {
	instance, err := es.NewERC20Swap(w.swapContractAddr, w.client)
	if err != nil {
		return false, err
	}

	refundable, err := instance.Refundable(nil, swapID)
	return refundable, err
}

// Claimable ...
func (w *Erc20Worker) Claimable(swapID common.Hash) (bool, error) {
	instance, err := es.NewERC20Swap(w.swapContractAddr, w.client)
	if err != nil {
		return false, err
	}

	claimable, err := instance.Claimable(nil, swapID)
	return claimable, err
}

func (w *Erc20Worker) GetTxCountLatest() (uint64, error) {
	var result hexutil.Uint64
	rpcClient := jsonrpc.NewClient(w.provider)

	resp, err := rpcClient.Call("eth_getTransactionCount", w.config.WorkerAddr.Hex(), "pending")
	if err != nil {
		return 0, err
	}
	if err := resp.GetObject(&result); err != nil {
		return 0, err
	}
	return uint64(result), nil
}

// GetTransactor ...
func (w *Erc20Worker) getTransactor() (auth *bind.TransactOpts, err error) {
	privateKey, err := utils.GetPrivateKey(w.config)
	if err != nil {
		return nil, err
	}

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(w.chainID))
	if err != nil {
		return nil, err
	}

	var nonce uint64
	// if w.chainName == "LA" {
	// 	nonce, err = w.GetTxCountLatest()
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// } else {
	nonce, err = w.client.PendingNonceAt(context.Background(), w.config.WorkerAddr)
	if err != nil {
		return nil, err
	}
	// }

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)                // in wei
	auth.GasLimit = uint64(w.config.GasLimit) // in units
	auth.GasPrice = w.config.GasPrice

	return auth, nil
}

// EthBalance ...
func (w *Erc20Worker) EthBalance(address common.Address) (*big.Int, error) {
	return w.client.BalanceAt(context.Background(), address, nil)
}

// GetWorkerAddress ...
func (w *Erc20Worker) GetWorkerAddress() string {
	return w.config.WorkerAddr.String()
}

// GetSentTxStatus ...
func (w *Erc20Worker) GetSentTxStatus(hash string) storage.TxStatus {
	// _, isPending, err := w.client.TransactionByHash(context.Background(), common.HexToHash(hash))
	// if err != nil {
	// 	w.logger.Errorln("GetSentTxStatus, err = ", err)
	// 	return storage.TxSentStatusNotFound
	// }

	// if isPending {
	// 	return storage.TxSentStatusPending
	// }

	txReceipt, err := w.client.TransactionReceipt(context.Background(), common.HexToHash(hash))
	if err != nil {
		return storage.TxSentStatusNotFound
	}

	if txReceipt.Status == types.ReceiptStatusFailed {
		return storage.TxSentStatusFailed
	}

	return storage.TxSentStatusSuccess
}

// GetColdWalletAddress ...
func (w *Erc20Worker) GetColdWalletAddress() string {
	return w.config.ColdWalletAddr.String()
}

// IsSameAddress ...
func (w *Erc20Worker) IsSameAddress(addrA string, addrB string) bool {
	return bytes.Equal(common.FromHex(addrA), common.FromHex(addrB))
}

// SendAmount ...
func (w *Erc20Worker) SendAmount(address string, amount *big.Int) (string, error) {
	return "", fmt.Errorf("not implemented") // TODO
}

func initTokenAddresses(client *ethclient.Client, addresses []common.Address) (tokenAddresses map[string]common.Address) {
	// for _, address := range addresses {
	// 	instance, err := es.NewERC20(address, client)
	// 	if err != nil {
	// 		log.Fatalf("Create new ERC20 instance, err=%v\n", err)
	// 		return
	// 	}

	// 	symbol, err := instance.Symbol(nil)
	// 	if err != nil {
	// 		log.Fatalf("Get token symbol, err=%v\n", err)
	// 		return
	// 	}
	// 	tokenAddresses[symbol] = address

	// 	time.Sleep(200 * time.Millisecond)
	// }

	return tokenAddresses
}

// func (w *Erc20Worker) CreateOrClaimRequest(swapID common.Hash, tokenAddr, erc20Addr, laAddr, recipientAddr, otherChainRecipientAddr string,
// 	timestamp int64, heightSpan int64, outAmount *big.Int) (string, error) {
// 	auth, err := w.GetTransactor()
// 	if err != nil {
// 		return "", err
// 	}

// 	instance, err := es.NewERC20Swap(w.swapContractAddr, w.client)
// 	if err != nil {
// 		return "", err
// 	}

// 	// convert address hex to ethereum like address
// 	tknAddr := common.HexToAddress(tokenAddr)
// 	recvAddr := common.HexToAddress(recipientAddr)

// 	tx, err := instance.CreateOrClaim(auth, swapID, outAmount, big.NewInt(heightSpan), [32]byte{}, uint64(timestamp), tknAddr, recvAddr)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tx.Hash().String(), nil
// }

// // CreateRequest ...
// func (w *Erc20Worker) CreateRequest(swapID common.Hash) (string, error) {
// 	auth, err := w.GetTransactor()
// 	if err != nil {
// 		return "", err
// 	}

// 	instance, err := bt.NewManager(w.swapContractAddr, w.client)
// 	if err != nil {
// 		return "", err
// 	}

// 	tx, err := instance.CreateSwapRequest(auth, swapID)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tx.Hash().String(), nil
// }

// // SendClaim ...
// func (w *Erc20Worker) SendClaim(swapID common.Hash) (string, error) {
// 	auth, err := w.GetTransactor()
// 	if err != nil {
// 		return "", err
// 	}

// 	instance, err := bt.NewManager(w.swapContractAddr, w.client)
// 	if err != nil {
// 		return "", err
// 	}

// 	tx, err := instance.ApproveSwap(auth, swapID)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tx.Hash().String(), nil
// }

// // GetBalanceAlertMsg ...
// func (w *Erc20Worker) GetBalanceAlertMsg(tokenSymbol string) (string, error) {
// 	if w.config.EthBalanceAlertThreshold.Cmp(big.NewInt(0)) == 0 &&
// 		w.config.TokenBalanceAlertThreshold.Cmp(big.NewInt(0)) == 0 &&
// 		w.config.AllowanceBalanceAlertThreshold.Cmp(big.NewInt(0)) == 0 {
// 		return "", nil
// 	}

// 	alertMsg := ""
// 	if w.config.EthBalanceAlertThreshold.Cmp(big.NewInt(0)) > 0 {
// 		ethBalance, err := w.EthBalance(w.config.WorkerAddr)
// 		if err != nil {
// 			return "", err
// 		}

// 		if ethBalance.Cmp(w.config.EthBalanceAlertThreshold) < 0 {
// 			alertMsg = alertMsg + fmt.Sprintf("eth balance(%s) is less than %s\n",
// 				ethBalance.String(), w.config.EthBalanceAlertThreshold.String())
// 		}
// 	}

// 	if w.config.AllowanceBalanceAlertThreshold.Cmp(big.NewInt(0)) > 0 {
// 		allowance, err := w.Allowance(tokenSymbol)
// 		if err != nil {
// 			return "", err
// 		}
// 		if allowance.Cmp(w.config.AllowanceBalanceAlertThreshold) < 0 {
// 			alertMsg = alertMsg + fmt.Sprintf("token allowance balance(%s) is less than %s\n",
// 				allowance.String(), w.config.AllowanceBalanceAlertThreshold.String())
// 		}
// 	}

// 	if w.config.TokenBalanceAlertThreshold.Cmp(big.NewInt(0)) > 0 {
// 		tokenBalance, err := w.Erc20Balance(w.config.WorkerAddr, "") // TODO
// 		if err != nil {
// 			return "", err
// 		}
// 		if tokenBalance.Cmp(w.config.TokenBalanceAlertThreshold) < 0 {
// 			alertMsg = alertMsg + fmt.Sprintf("token balance(%s) is less than %s",
// 				tokenBalance.String(), w.config.TokenBalanceAlertThreshold.String())
// 		}
// 	}

// 	return alertMsg, nil
// }
