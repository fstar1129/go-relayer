package eth

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"latoken/relayer-smart-contract/src/service/workers/utils"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"

	"latoken/relayer-smart-contract/src/models"
	"latoken/relayer-smart-contract/src/service/storage"
	ethbr "latoken/relayer-smart-contract/src/service/workers/eth-compatible/abi/bridge/eth"
	labr "latoken/relayer-smart-contract/src/service/workers/eth-compatible/abi/bridge/la"
	es "latoken/relayer-smart-contract/src/service/workers/eth-compatible/abi/erc20Swap"
)

// Erc20Worker ...
type Erc20Worker struct {
	Provider         string
	chainID          string
	logger           *logrus.Entry // logger
	Config           *models.WorkerConfig
	Client           *ethclient.Client
	swapContractAddr common.Address
}

// NewErc20Worker ...
func NewErc20Worker(logger *logrus.Logger, cfg *models.WorkerConfig) *Erc20Worker {
	client, err := ethclient.Dial(cfg.Provider)
	if err != nil {
		panic("new eth client error")
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

	// init token addresses
	return &Erc20Worker{
		chainID:          cfg.ChainID,
		logger:           logger.WithField("worker", cfg.ChainID),
		Provider:         cfg.Provider,
		Config:           cfg,
		Client:           client,
		swapContractAddr: cfg.ContractAddr,
	}
}

// GetChain returns chain ID
func (w *Erc20Worker) GetChain() string {
	return string(w.chainID)
}

// GetStartHeight returns start blockchain height from config
func (w *Erc20Worker) GetStartHeight() int64 {
	return w.Config.StartBlockHeight
}

// GetConfirmNum returns numbers of blocks after them tx will be confirmed
func (w *Erc20Worker) GetConfirmNum() int64 {
	return w.Config.ConfirmNum
}

// // GetStatus returns status of relayer account(balance eg)
// func (w *Erc20Worker) GetStatus(symbol string) (interface{}, error) {
// 	ethStatus := &EthStatus{}

// 	allowance, err := w.Allowance(symbol)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ethStatus.Allowance = QuoBigInt(allowance, GetBigIntForDecimal(w.Config.ChainDecimal)).String()

// 	balance, err := w.Erc20Balance(w.Config.WorkerAddr, symbol)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ethStatus.Erc20Balance = QuoBigInt(balance, GetBigIntForDecimal(w.Config.ChainDecimal)).String()

// 	ethBalance, err := w.EthBalance(w.Config.WorkerAddr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ethStatus.EthBalance = QuoBigInt(ethBalance, GetBigIntForDecimal(18)).String()

// 	return ethStatus, nil
// }

// GetBlockAndTxs ...
func (w *Erc20Worker) GetBlockAndTxs(height int64) (*models.BlockAndTxLogs, error) {
	header, err := w.Client.HeaderByNumber(context.Background(), big.NewInt(height))
	if err != nil {
		return nil, err
	}

	txLogs, err := w.getLogs(header.Hash())
	if err != nil {
		return nil, err
	}

	return &models.BlockAndTxLogs{
		Height:          height,
		BlockHash:       header.Hash().String(),
		ParentBlockHash: header.ParentHash.String(),
		BlockTime:       int64(header.Time),
		TxLogs:          txLogs,
	}, nil
}

// GetFetchInterval ...
func (w *Erc20Worker) GetFetchInterval() time.Duration {
	return time.Duration(w.Config.FetchInterval) * time.Second
}

// getLogs ...
func (w *Erc20Worker) getLogs(blockHash common.Hash) ([]*storage.TxLog, error) {
	//	topics := [][]common.Hash{{DepositEventHash, ProposalEventHash, ProposalVoteHash}}
	logs, err := w.Client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: &blockHash,
		//	Topics:    topics,
		Addresses: []common.Address{w.swapContractAddr},
	})
	if err != nil {
		return nil, err
	}

	models := make([]*storage.TxLog, 0, len(logs))
	for _, log := range logs {
		fmt.Printf("WORKER(%s) NEW EVENT: %v\n\n", w.chainID, log.Topics)
		event, err := ParseEvent(&log)
		if err != nil {
			w.logger.WithFields(logrus.Fields{"function": "GetLogs()"}).Errorf("parse event log error, err=%s", err)
			continue
		}
		if event == nil {
			continue
		}

		txLog := event.ToTxLog()
		// if txLog.TxType == storage.TxTypeDeposit {
		// 	txLog.SwapType = storage.SwapTypeBind
		// 	if w.chainID == storage.LaChain {
		// 		txLog.SwapType = storage.SwapTypeUnbind
		// 	}
		// } else if txLog.TxType == storage.TxTypeSpend {
		// 	txLog.SwapType = storage.SwapTypeUnbind
		// 	if w.chainID == storage.LaChain {
		// 		txLog.SwapType = storage.SwapTypeBind
		// 	}
		// }
		txLog.Chain = w.chainID
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
	header, err := w.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, nil
	}
	return header.Number.Int64(), nil
}

// Vote ...
func (w *Erc20Worker) Vote(depositNonce uint64, resourceID [32]byte, data []byte) (string, error) {
	auth, err := w.GetTransactor()
	if err != nil {
		return "", err
	}

	instance, err := labr.NewBridge(w.swapContractAddr, w.Client)
	if err != nil {
		return "", err
	}

	fmt.Println(utils.StringToBytes8(w.chainID))

	tx, err := instance.VoteProposal(auth, utils.StringToBytes8(w.chainID), depositNonce, resourceID, data)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

// Spend spends bind request(==mints token erc-20 on Lachain)
// Requires:
// - swapID- swap's id
// - erc20Addr - erc20 token address, laAddr - lrc-20 token address(may to bind on pair of tokens(erc-20 and lrc-20) later on)
// - toAddress - receiver address, which will receive minted coins
// - amount - value of minted coins
func (w *Erc20Worker) SpendBind(depositNonce uint64, chainID [8]byte, data []byte, resourceID [32]byte) (string, error) {
	auth, err := w.GetTransactor()
	if err != nil {
		return "", err
	}

	instance, err := labr.NewBridge(w.swapContractAddr, w.Client)
	if err != nil {
		return "", err
	}

	fmt.Println(chainID, depositNonce, data, resourceID)

	tx, err := instance.ExecuteProposal(auth, chainID, depositNonce, data, resourceID)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

// SpendUnbind spends unbind request(==sends erc-20 token to receiver on Ethereum)
// Requires:
// - swapID- swap's id
// - toAddress - receiver address, which will receive minted coins
// - amount - value of minted coins
func (w *Erc20Worker) SpendUnbind(depositNonce uint64, data []byte, resourceID [32]byte) (string, error) {
	auth, err := w.GetTransactor()
	if err != nil {
		return "", err
	}

	instance, err := ethbr.NewBridge(w.swapContractAddr, w.Client)
	if err != nil {
		return "", err
	}

	tx, err := instance.ExecuteProposal(auth, utils.StringToBytes8(w.chainID), depositNonce, data, resourceID)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

// GetSentTxStatus ...
func (w *Erc20Worker) GetSentTxStatus(hash string) storage.TxStatus {
	_, isPending, err := w.Client.TransactionByHash(context.Background(), common.HexToHash(hash))
	if err != nil {
		w.logger.Errorln("GetSentTxStatus, err = ", err)
		return storage.TxSentStatusNotFound
	}

	if isPending {
		return storage.TxSentStatusPending
	}

	txReceipt, err := w.Client.TransactionReceipt(context.Background(), common.HexToHash(hash))
	if err != nil {
		w.logger.Errorln("GetSentTxStatus, err = ", err)
		return storage.TxSentStatusNotFound
	}

	if txReceipt.Status == types.ReceiptStatusFailed {
		return storage.TxSentStatusFailed
	}

	return storage.TxSentStatusSuccess
}

// HasSwap ...
func (w *Erc20Worker) HasSwap(swapID common.Hash) (bool, error) {
	instance, err := es.NewERC20Swap(w.swapContractAddr, w.Client)
	if err != nil {
		return false, err
	}

	return instance.IsSwapExist(nil, swapID)
}

// Refundable ...
func (w *Erc20Worker) Refundable(swapID common.Hash) (bool, error) {
	instance, err := es.NewERC20Swap(w.swapContractAddr, w.Client)
	if err != nil {
		return false, err
	}

	refundable, err := instance.Refundable(nil, swapID)
	return refundable, err
}

// Claimable ...
func (w *Erc20Worker) Claimable(swapID common.Hash) (bool, error) {
	instance, err := es.NewERC20Swap(w.swapContractAddr, w.Client)
	if err != nil {
		return false, err
	}

	claimable, err := instance.Claimable(nil, swapID)
	return claimable, err
}

// GetTransactor ...
func (w *Erc20Worker) GetTransactor() (*bind.TransactOpts, error) {
	privateKey, err := utils.GetPrivateKey(w.Config)
	if err != nil {
		return nil, err
	}

	nonce, err := w.Client.PendingNonceAt(context.Background(), w.Config.WorkerAddr)
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)                // in wei
	auth.GasLimit = uint64(w.Config.GasLimit) // in units
	auth.GasPrice = w.Config.GasPrice

	return auth, nil
}

// EthBalance ...
func (w *Erc20Worker) EthBalance(address common.Address) (*big.Int, error) {
	return w.Client.BalanceAt(context.Background(), address, nil)
}

// GetWorkerAddress ...
func (w *Erc20Worker) GetWorkerAddress() string {
	return w.Config.WorkerAddr.String()
}

// GetColdWalletAddress ...
func (w *Erc20Worker) GetColdWalletAddress() string {
	return w.Config.ColdWalletAddr.String()
}

// IsSameAddress ...
func (w *Erc20Worker) IsSameAddress(addrA string, addrB string) bool {
	return bytes.Equal(common.FromHex(addrA), common.FromHex(addrB))
}

// SendAmount ...
func (w *Erc20Worker) SendAmount(address string, amount *big.Int) (string, error) {
	return "", fmt.Errorf("not implemented") // TODO
}

// // GetHTLTEvent ...
// func (w *Erc20Worker) GetHTLTEvent(swapID common.Hash) (*HTLTEvent, error) {
// 	topics := [][]common.Hash{{HTLTEventHash}, {}, {}, {swapID}}
// 	logs, err := w.Client.FilterLogs(context.Background(), ethereum.FilterQuery{
// 		Topics: topics,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(logs) == 0 {
// 		return nil, fmt.Errorf("swap id does not exist, swap_id=%s", swapID.String())
// 	}

// 	event, err := ParseHTLTEvent(&w.abi, &logs[0])
// 	if err != nil {
// 		w.logger.WithFields(logrus.Fields{"function": "GetLogs()"}).Errorln(err)
// 		return nil, err
// 	}
// 	htltEvent := event.(HTLTEvent)

// 	return &htltEvent, nil
// }

// // GetSwap ...
// func (w *Erc20Worker) GetSwap(swapID common.Hash) (*models.SwapRequest, error) {
// 	htltEvent, err := w.GetHTLTEvent(swapID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &models.SwapRequest{
// 		ID:               swapID,
// 		ExpireHeight:     htltEvent.ExpireHeight.Int64(),
// 		SenderAddress:    htltEvent.MsgSender.String(),
// 		RecipientAddress: htltEvent.LaRecipientAddr.String(),
// 		OutAmount:        htltEvent.OutAmount,
// 	}, nil
// }

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

// 	instance, err := es.NewERC20Swap(w.swapContractAddr, w.Client)
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

// 	instance, err := bt.NewManager(w.swapContractAddr, w.Client)
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

// 	instance, err := bt.NewManager(w.swapContractAddr, w.Client)
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
// 	if w.Config.EthBalanceAlertThreshold.Cmp(big.NewInt(0)) == 0 &&
// 		w.Config.TokenBalanceAlertThreshold.Cmp(big.NewInt(0)) == 0 &&
// 		w.Config.AllowanceBalanceAlertThreshold.Cmp(big.NewInt(0)) == 0 {
// 		return "", nil
// 	}

// 	alertMsg := ""
// 	if w.Config.EthBalanceAlertThreshold.Cmp(big.NewInt(0)) > 0 {
// 		ethBalance, err := w.EthBalance(w.Config.WorkerAddr)
// 		if err != nil {
// 			return "", err
// 		}

// 		if ethBalance.Cmp(w.Config.EthBalanceAlertThreshold) < 0 {
// 			alertMsg = alertMsg + fmt.Sprintf("eth balance(%s) is less than %s\n",
// 				ethBalance.String(), w.Config.EthBalanceAlertThreshold.String())
// 		}
// 	}

// 	if w.Config.AllowanceBalanceAlertThreshold.Cmp(big.NewInt(0)) > 0 {
// 		allowance, err := w.Allowance(tokenSymbol)
// 		if err != nil {
// 			return "", err
// 		}
// 		if allowance.Cmp(w.Config.AllowanceBalanceAlertThreshold) < 0 {
// 			alertMsg = alertMsg + fmt.Sprintf("token allowance balance(%s) is less than %s\n",
// 				allowance.String(), w.Config.AllowanceBalanceAlertThreshold.String())
// 		}
// 	}

// 	if w.Config.TokenBalanceAlertThreshold.Cmp(big.NewInt(0)) > 0 {
// 		tokenBalance, err := w.Erc20Balance(w.Config.WorkerAddr, "") // TODO
// 		if err != nil {
// 			return "", err
// 		}
// 		if tokenBalance.Cmp(w.Config.TokenBalanceAlertThreshold) < 0 {
// 			alertMsg = alertMsg + fmt.Sprintf("token balance(%s) is less than %s",
// 				tokenBalance.String(), w.Config.TokenBalanceAlertThreshold.String())
// 		}
// 	}

// 	return alertMsg, nil
// }
