package eth

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/storage"
	labr "gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/workers/eth-compatible/abi/bridge/la"
)

////
// EVENTS NAMES
////

const (
	// DepositEventName ...
	DepositEventName = "Deposit"

	ProposalVoteName = "ProposalVote"

	// ProposalEventName ...
	ProposalEventName = "ProposalEvent"
)

////
// EVENTS HASHES  | web3.utils.sha3('HTLT(types,...)');
////

var (
	// HTLTEventHash ...
	HTLTEventHash = common.HexToHash("0xa9d611af8ea77eff0d427e2ae8375a23513129a25c615378cff5594b03c36ce9")

	// ClaimedEventHash ...
	ClaimedEventHash = common.HexToHash("0x91d697238e9aa9f3172d17522c9be529b94a892481554e1ea619369b5b12f39a")

	// CreateUnbindEventHash ...
	CreateUnbindEventHash = common.HexToHash("0xf3f48b99c0f6ef1f16092aed9ebd7d552e8b7647e0332990550d7626d7857ba5")

	// ConfirmedEventHash ...
	ConfirmedEventHash = common.HexToHash("0xf3f48b99c0f6ef1f16092aed9ebd7d552e8b7647e0332990550d7626d7857ba5")

	// InitEventHash ...
	InitEventHash = common.HexToHash("0xf3f48b99c0f6ef1f16092aed9ebd7d552e8b7647e0332990550d7626d7857ba5")

	// RefundEventHash ...
	RefundEventHash = common.HexToHash("0x04eb8ae268f23cfe2f9d72fa12367b104af16959f6a93530a4cc0f50688124f9")

	// DepositEventHash
	DepositEventHash = common.HexToHash("0x6a8f7ffca1e3ab54554c709a9eb23e11fc5063f0b24e40087c40ab3e7027f189")

	// ProposalVoteHash
	ProposalVoteHash = common.HexToHash("0x85f41114efc645854a10eef33ef4dec54341cb3ec3ab32386c92c881f3b1b505")

	// ProposalEventHash
	ProposalEventHash = common.HexToHash("0x8dc49847a011c3b316cd0f50cf982e0fd5b3ddb7fdf970fc81a25557f0923a73")
)

// ContractEvent ...
type ContractEvent interface {
	ToTxLog() *storage.TxLog
}

// DepositEvent represents a Deposit event raised by the Bridge.sol contract.
type DepositEvent struct {
	DestinationChainID [8]byte
	ResourceID         [32]byte
	DepositNonce       uint64
	Depositor          common.Address
	RecipientAddress   common.Address
	TokenAddress       common.Address
	Amount             *big.Int
}

// resource id: 000000000000000000000000000000c76ebe4a02bbc34786d860b355f5a5ce00

// resource ID: 0x000000000000000000000000000000c76ebe4a02bbc34786d860b355f5a5ce00

// ParseDepositEvent ...
func ParseDepositEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev DepositEvent
	if err := abi.UnpackIntoInterface(&ev, DepositEventName, log.Data); err != nil {
		return nil, err
	}

	// ev.DestinationChainID = common.RightPadBytes(log.Topics[1].Bytes(), 8)
	// ev.ResourceID = common.RightPadBytes(log.Topics[2].Bytes(), 32)
	// ev.DepositNonce = big.NewInt(0).SetBytes(log.Topics[3].Bytes()).Uint64()

	fmt.Printf("Deposited\n")
	fmt.Printf("destination chain ID: 0x%s\n", common.Bytes2Hex(ev.DestinationChainID[:]))
	fmt.Printf("resource ID: 0x%s\n", common.Bytes2Hex(ev.ResourceID[:]))
	fmt.Printf("deposit nonce: %d\n", ev.DepositNonce)
	fmt.Printf("depositor address: %s\n", ev.Depositor.Hex())
	fmt.Printf("recipient address: %s\n", ev.RecipientAddress.Hex())
	fmt.Printf("token address: %s\n", ev.TokenAddress.Hex())
	fmt.Printf("amount : %s\n", ev.Amount.String())

	return ev, nil
}

// ToTxLog ...
func (ev DepositEvent) ToTxLog() *storage.TxLog {
	//	if
	return &storage.TxLog{
		Chain:              string(ev.ResourceID[:]),
		TxType:             storage.TxTypeDeposit,
		DestinationChainID: common.Bytes2Hex(ev.DestinationChainID[:]),
		SwapID:             fmt.Sprintf("%d", ev.DepositNonce),
		ResourceID:         common.Bytes2Hex(ev.ResourceID[:]),
		DepositNonce:       ev.DepositNonce,
		SenderAddr:         ev.Depositor.Hex(),
		ReceiverAddr:       ev.RecipientAddress.Hex(),
		InTokenAddr:        ev.TokenAddress.Hex(),
		OutAmount:          ev.Amount.String(),
	}
}

// ProposalVoteEvent represents a Deposit event raised by the Bridge.sol contract.
type ProposalVoteEvent struct {
	OriginChainID [8]byte
	DepositNonce  uint64
	Status        uint8
	ResourceID    [32]byte
}

// ParseProposalVote ...
func ParseProposalVote(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev ProposalVoteEvent
	if err := abi.UnpackIntoInterface(&ev, ProposalVoteName, log.Data); err != nil {
		return nil, err
	}

	// ev.OriginChainID = log.Topics[1].Bytes()
	// ev.DepositNonce = big.NewInt(0).SetBytes(log.Topics[2].Bytes()).Uint64()
	// ev.Status = uint8(big.NewInt(0).SetBytes(log.Topics[3].Bytes()).Uint64())

	fmt.Printf("ProposalVote\n")
	fmt.Printf("origin chain ID: 0x%s\n", common.Bytes2Hex(ev.OriginChainID[:]))
	fmt.Printf("deposit nonce: %d\n", ev.DepositNonce)
	fmt.Printf("status: %d\n", ev.Status)
	fmt.Printf("resource id: %s\n\n", common.Bytes2Hex(ev.ResourceID[:]))

	return ev, nil
}

// ToTxLog ...
func (ev ProposalVoteEvent) ToTxLog() *storage.TxLog {
	return &storage.TxLog{
		Chain:         storage.EthChain,
		TxType:        storage.TxTypeClaim,
		SwapID:        fmt.Sprintf("%d", ev.DepositNonce),
		OriginСhainID: common.Bytes2Hex(ev.OriginChainID[:]),
		DepositNonce:  ev.DepositNonce,
		SwapStatus:    ev.Status,
		ResourceID:    common.Bytes2Hex(ev.ResourceID[:]),
	}
}

// ProposalEvent represents a Deposit event raised by the Bridge.sol contract.
type ProposalEvent struct {
	OriginChainID [8]byte
	DepositNonce  uint64
	Status        uint8
	ResourceID    [32]byte
	DataHash      [32]byte
}

// // ParseProposalEvent ...
// func ParseProposalEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
// 	var ev ProposalEvent
// 	if err := abi.UnpackIntoInterface(&ev, ProposalEventName, log.Data); err != nil {
// 		return nil, err
// 	}
// 	ev.OriginChainID = log.Topics[1].Bytes()
// 	ev.DepositNonce = big.NewInt(0).SetBytes(log.Topics[2].Bytes()).Uint64()
// 	ev.Status = uint8(big.NewInt(0).SetBytes(log.Topics[3].Bytes()).Uint64())

// 	fmt.Printf("ProposalEvent\n")
// 	fmt.Printf("origin chain ID: 0x%s\n", common.Bytes2Hex(ev.OriginChainID[:]))
// 	fmt.Printf("deposit nonce: %d\n", ev.DepositNonce)
// 	fmt.Printf("status: %d\n", ev.Status)
// 	fmt.Printf("resource ID: 0x%s\n", common.Bytes2Hex(ev.ResourceID[:]))
// 	fmt.Printf("DataHash: 0x%s\n\n", common.Bytes2Hex(ev.DataHash[:]))

// 	return ev, nil
// }

// ToTxLog ...
func (ev ProposalEvent) ToTxLog() *storage.TxLog {
	// if status == 2 -> already claimed -> mint
	// if status == 3-> already minted(executed)
	txlog := &storage.TxLog{
		Chain:         storage.EthChain,
		TxType:        storage.TxTypeClaim,
		SwapID:        fmt.Sprintf("%d", ev.DepositNonce),
		OriginСhainID: common.Bytes2Hex(ev.OriginChainID[:]),
		DepositNonce:  ev.DepositNonce,
		SwapStatus:    ev.Status,
		ResourceID:    common.Bytes2Hex(ev.ResourceID[:]),
	}

	if ev.Status == uint8(2) {
		txlog.TxType = storage.TxTypePassed
	} else if ev.Status == uint8(3) {
		txlog.TxType = storage.TxTypeSpend
	}

	return txlog
}

// ParseEvent ...
func (w *Erc20Worker) parseEvent(log *types.Log) (ContractEvent, error) {
	if bytes.Equal(log.Topics[0][:], DepositEventHash[:]) {
		fmt.Println(w.chainName)
		if w.chainName == storage.EthChain {
			return ParseEthDepositEvent(log)
		} else if w.chainName == storage.LaChain {
			return ParseLaDepositEvent(log)
		}
	} else if bytes.Equal(log.Topics[0][:], ProposalEventHash[:]) {
		fmt.Println(w.chainName)
		if w.chainName == storage.EthChain {
			return ParseEthProposalEvent(log)
		} else if w.chainName == storage.LaChain {
			return ParseLaProposalEvent(log)
		}
	} else if bytes.Equal(log.Topics[0][:], ProposalVoteHash[:]) {
		abi, _ := abi.JSON(strings.NewReader(labr.BridgeABI))
		return ParseProposalVote(&abi, log)
	}

	return nil, nil
}

type Header struct {
	Hash       common.Hash    `json:"hash"`
	ParentHash common.Hash    `json:"parentHash"       gencodec:"required"`
	Time       hexutil.Uint64 `json:"timestamp"        gencodec:"required"`
}
