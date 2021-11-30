package eth

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/storage"
	"gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/workers/utils"
)

////
// EVENTS NAMES
////

const (
	// DepositEventName ...
	DepositEventName = "Deposit"
)

////
// EVENTS HASHES  | web3.utils.sha3('HTLT(types,...)');
////

var (
	// DepositEventHash
	DepositEventHash = common.HexToHash("0x5d633e6cc5b698cdfbcf6dffac28414bfb25b893e0cc792c53d17ce1f46ccb5c")
)

// ContractEvent ...
type ContractEvent interface {
	ToTxLog() *storage.TxLog
}

// DepositEvent represents a Deposit event raised by the Bridge.sol contract.
type DepositEvent struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	ResourceID         [32]byte
	DepositNonce       uint64
	Depositor          common.Address
	RecipientAddress   common.Address
	TokenAddress       common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// ToTxLog ...
func (ev DepositEvent) ToTxLog() *storage.TxLog {
	return &storage.TxLog{
		Chain:              string(ev.ResourceID[:]),
		TxType:             storage.TxTypeDeposit,
		DestinationChainID: common.Bytes2Hex(ev.DestinationChainID[:]),
		SwapID:             utils.CalcutateSwapID(ev.Amount.String(), ev.RecipientAddress.Hex(), "0x8D960cbDc18eE3D10169EC27058B33eae55A7C35"),
		ResourceID:         common.Bytes2Hex(ev.ResourceID[:]),
		DepositNonce:       ev.DepositNonce,
		SenderAddr:         ev.Depositor.Hex(),
		ReceiverAddr:       ev.RecipientAddress.Hex(),
		InTokenAddr:        ev.TokenAddress.Hex(),
		OutAmount:          ev.Amount.String(),
	}
}

// ParseEvent ...
func (w *Erc20Worker) parseEvent(log *types.Log) (ContractEvent, error) {
	if bytes.Equal(log.Topics[0][:], DepositEventHash[:]) {
		if w.chainName == storage.EthChain {
			return ParseEthDepositEvent(log)
		} else if w.chainName == storage.LaChain {
			return ParseLaDepositEvent(log)
		}
	}

	return nil, nil
}

type Header struct {
	Hash       common.Hash    `json:"hash"`
	ParentHash common.Hash    `json:"parentHash"       gencodec:"required"`
	Time       hexutil.Uint64 `json:"timestamp"        gencodec:"required"`
}
