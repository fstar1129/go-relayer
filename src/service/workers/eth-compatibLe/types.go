package eth

import (
	"bytes"
	"fmt"
	"latoken/relayer-smart-contract/src/service/storage"
	ethbr "latoken/relayer-smart-contract/src/service/workers/eth-compatible/abi/bridge/eth"
	labr "latoken/relayer-smart-contract/src/service/workers/eth-compatible/abi/bridge/la"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const workerChainDecimal = 8

////
// EVENTS NAMES
////

const (
	// InitEventName ...
	InitEventName = "Init"

	// HTLTEventName ...
	HTLTEventName = "HTLT"

	// ClaimEventName ...
	ClaimEventName = "Claimed"

	// CreateUnbindEventName ...
	CreateUnbindEventName = "CreateUnbind"

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
	DepositEventHash = common.HexToHash("0xb17d269f66f091233e7970a3f50b6bed322eaccf27b39f83d94277be64cf4bea")

	// ProposalVoteHash
	ProposalVoteHash = common.HexToHash("0xcf0a87860fb1c3985a5ba36f56f0d37a1af050a78bad02ac50b51c29f3c2485c")

	// ProposalEventHash
	ProposalEventHash = common.HexToHash("0x1281a0584988174583df37b5112e6760e9a7d4845e65a832868aeffc0e547473")
)

// ContractEvent ...
type ContractEvent interface {
	ToTxLog() *storage.TxLog
}

// DepositEvent represents a Deposit event raised by the Bridge.sol contract.
type DepositEvent struct {
	DestinationChainID []byte
	ResourceID         []byte
	DepositNonce       uint64
	Depositor          common.Address
	RecipientAddress   []byte
	TokenAddress       common.Address
	Data               []byte
}

// ParseDepositEvent ...
func ParseDepositEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev DepositEvent
	if err := abi.UnpackIntoInterface(&ev, DepositEventName, log.Data); err != nil {
		return nil, err
	}

	ev.DestinationChainID = common.RightPadBytes(log.Topics[1].Bytes(), 8)
	ev.ResourceID = common.RightPadBytes(log.Topics[2].Bytes(), 32)
	ev.DepositNonce = big.NewInt(0).SetBytes(log.Topics[3].Bytes()).Uint64()

	fmt.Printf("Deposited\n")
	fmt.Printf("destination chain ID: 0x%s\n", common.Bytes2Hex(ev.DestinationChainID))
	fmt.Printf("resource ID: 0x%s\n", common.Bytes2Hex(ev.ResourceID))
	fmt.Printf("deposit nonce: %d\n", ev.DepositNonce)
	fmt.Printf("depositor address: %s\n", ev.Depositor.Hex())
	fmt.Printf("recipient address: %s\n", common.Bytes2Hex(ev.RecipientAddress))
	fmt.Printf("token address: %s\n", ev.TokenAddress.Hex())
	fmt.Printf("data: %s\n", common.Bytes2Hex(ev.Data))

	return ev, nil
}

// ToTxLog ...
func (ev DepositEvent) ToTxLog() *storage.TxLog {
	//	if
	return &storage.TxLog{
		Chain:              string(ev.ResourceID[:]),
		TxType:             storage.TxTypeDeposit,
		DestinationChainID: common.Bytes2Hex(ev.DestinationChainID),
		SwapID:             fmt.Sprintf("%d", ev.DepositNonce),
		ResourceID:         common.Bytes2Hex(ev.ResourceID),
		DepositNonce:       ev.DepositNonce,
		SenderAddr:         ev.Depositor.Hex(),
		ReceiverAddr:       common.Bytes2Hex(ev.RecipientAddress),
		ERC20TokenAddr:     ev.TokenAddress.Hex(),
		DataHash:           common.Bytes2Hex(ev.Data),
	}
}

// ProposalVoteEvent represents a Deposit event raised by the Bridge.sol contract.
type ProposalVoteEvent struct {
	OriginChainID [32]byte
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

	//	ev.OriginChainID = common.RightPadBytes(log.Topics[1].Bytes(), 32)
	ev.DepositNonce = big.NewInt(0).SetBytes(log.Topics[1].Bytes()).Uint64()
	//ev.Status = common.RightPadBytes(log.Topics[2].Bytes(), 8)
	ev.Status = uint8(big.NewInt(0).SetBytes(log.Topics[2].Bytes()).Uint64())

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
	OriginChainID [32]byte
	DepositNonce  uint64
	Status        uint8
	ResourceID    [32]byte
	DataHash      [32]byte
}

// ParseProposalEvent ...
func ParseProposalEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev ProposalEvent
	if err := abi.UnpackIntoInterface(&ev, ProposalEventName, log.Data); err != nil {
		return nil, err
	}

	//ev.OriginChainID = common.RightPadBytes(log.Topics[1].Bytes(), 32)
	ev.DepositNonce = big.NewInt(0).SetBytes(log.Topics[1].Bytes()).Uint64()
	ev.Status = uint8(big.NewInt(0).SetBytes(log.Topics[2].Bytes()).Uint64())

	fmt.Printf("ProposalEvent\n")
	fmt.Printf("origin chain ID: 0x%s\n", common.Bytes2Hex(ev.OriginChainID[:]))
	fmt.Printf("deposit nonce: %d\n", ev.DepositNonce)
	fmt.Printf("status: %d\n", ev.Status)
	fmt.Printf("resource ID: 0x%s\n", common.Bytes2Hex(ev.ResourceID[:]))
	fmt.Printf("DataHash: 0x%s\n\n", common.Bytes2Hex(ev.DataHash[:]))

	return ev, nil
}

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
		//DataHash:      common.Bytes2Hex(ev.DataHash[:]),
	}

	if ev.Status == uint8(2) {
		txlog.TxType = storage.TxTypePassed
	} else if ev.Status == uint8(3) {
		txlog.TxType = storage.TxTypeSpend
	}

	return txlog
}

// ParseEvent ...
func ParseEvent(log *types.Log) (ContractEvent, error) {
	if bytes.Equal(log.Topics[0][:], DepositEventHash[:]) {
		abi, _ := abi.JSON(strings.NewReader(ethbr.BridgeABI))
		return ParseDepositEvent(&abi, log)
	} else if bytes.Equal(log.Topics[0][:], ProposalEventHash[:]) {
		abi, _ := abi.JSON(strings.NewReader(labr.BridgeABI))
		return ParseProposalEvent(&abi, log)
	} else if bytes.Equal(log.Topics[0][:], ProposalVoteHash[:]) {
		abi, _ := abi.JSON(strings.NewReader(labr.BridgeABI))
		return ParseProposalVote(&abi, log)
	}

	return nil, nil
}

// // InitEvent ...
// type InitEvent struct {
// 	Timestamp        uint64
// 	MsgSender        common.Address
// 	RecipientAddr    common.Address
// 	SwapID           common.Hash
// 	RandomNumberHash common.Hash
// 	RandomNumber     common.Hash
// 	RefundAddr       common.Hash
// 	ExpireHeight     *big.Int
// 	OutAmount        *big.Int
// 	LaAmount         *big.Int
// }

// // ParseInitEvent ...
// func ParseInitEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
// 	var ev InitEvent
// 	if err := abi.UnpackIntoInterface(&ev, InitEventName, log.Data); err != nil {
// 		return nil, err
// 	}

// 	ev.MsgSender = common.BytesToAddress(log.Topics[1].Bytes())
// 	ev.RecipientAddr = common.BytesToAddress(log.Topics[2].Bytes())
// 	ev.SwapID = common.BytesToHash(log.Topics[3].Bytes())

// 	fmt.Printf("UNBIND\n")
// 	fmt.Printf("sender addr: %s\n", ev.MsgSender.String())
// 	fmt.Printf("receiver addr: %s\n", ev.RecipientAddr.String())
// 	fmt.Printf("swap id: %s \n", hex.EncodeToString(ev.SwapID[:]))
// 	fmt.Printf("timestamp: %d\n", ev.Timestamp)
// 	fmt.Printf("expire height: %d\n", ev.ExpireHeight)
// 	fmt.Printf("erc20 amount: %d\n", ev.OutAmount)
// 	fmt.Printf("la amount: %d\n", ev.LaAmount)
// 	return ev, nil
// }

// // ToTxLog ...
// func (ev InitEvent) ToTxLog() *storage.TxLog {
// 	return &storage.TxLog{
// 		Chain:        storage.EthChain,
// 		SwapType:     storage.SwapTypeUnbind,
// 		TxType:       storage.TxTypeCreate,
// 		SwapID:       ev.SwapID.String(),
// 		SenderAddr:   ev.MsgSender.Hex(),
// 		ReceiverAddr: ev.RecipientAddr.Hex(),
// 		Timestamp:    int64(ev.Timestamp),
// 		ExpireHeight: ev.ExpireHeight.Int64(),
// 		InAmount:     ev.LaAmount.String(),
// 		OutAmount:    ev.OutAmount.String(),
// 	}
// }

// // HTLTEvent ...
// type HTLTEvent struct {
// 	Timestamp        uint64
// 	MsgSender        common.Address
// 	LaRecipientAddr  common.Address
// 	SwapID           common.Hash
// 	RandomNumberHash common.Hash
// 	ERC20TokenAddr   common.Address
// 	LRC20TokenAddr   common.Address
// 	RefundAddr       common.Address
// 	ExpireHeight     *big.Int
// 	OutAmount        *big.Int
// 	LaAmount         *big.Int
// }

// // ParseHTLTEvent ...
// func ParseHTLTEvent(abix *abi.ABI, log *types.Log) (ContractEvent, error) {
// 	//
// 	var ev HTLTEvent
// 	err := abix.UnpackIntoInterface(&ev, HTLTEventName, log.Data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ev.MsgSender = common.BytesToAddress(log.Topics[1].Bytes())
// 	ev.LaRecipientAddr = common.BytesToAddress(log.Topics[2].Bytes())
// 	ev.SwapID = common.BytesToHash(log.Topics[3].Bytes())

// 	fmt.Printf("HTLT\n")
// 	fmt.Printf("sender addr: %s\n", ev.MsgSender.String())
// 	fmt.Printf("lachain receiver addr: %s\n", ev.LaRecipientAddr.String())
// 	fmt.Printf("swap ID: %s\n", ev.SwapID.String())
// 	fmt.Printf("Random num hash: %s\n", ev.RandomNumberHash.String())
// 	fmt.Printf("erc20 addr: %s\n", ev.ERC20TokenAddr.String())
// 	fmt.Printf("lrc20 addr: %s\n", ev.LRC20TokenAddr.String())
// 	fmt.Printf("timestamp: %d\n", ev.Timestamp)
// 	fmt.Printf("expire height: %d\n", ev.ExpireHeight)
// 	fmt.Printf("erc20 amount: %d\n", ev.OutAmount)
// 	fmt.Printf("la amount: %d\n", ev.LaAmount)
// 	return ev, nil
// }

// // ToTxLog ...
// func (ev HTLTEvent) ToTxLog() *storage.TxLog {
// 	swapID := calculateSwapID(ev.RandomNumberHash.Bytes(), ev.MsgSender.Bytes(), ev.LaRecipientAddr.Bytes())

// 	fmt.Println(hex.EncodeToString(swapID), " | ", ev.SwapID.String())

// 	return &storage.TxLog{
// 		Chain:          storage.EthChain,
// 		SwapType:       storage.SwapTypeBind,
// 		TxType:         storage.TxTypeHTLT,
// 		SwapID:         ev.SwapID.String(),
// 		ERC20TokenAddr: ev.ERC20TokenAddr.Hex(),
// 		LRC20TokenAddr: ev.LRC20TokenAddr.Hex(),
// 		SenderAddr:     ev.MsgSender.Hex(),
// 		ReceiverAddr:   ev.LaRecipientAddr.Hex(),
// 		Timestamp:      int64(ev.Timestamp),
// 		ExpireHeight:   ev.ExpireHeight.Int64(),
// 		InAmount:       ev.LaAmount.String(),
// 		OutAmount:      ev.OutAmount.String(),
// 	}
// }

// // ClaimedEvent ...
// type ClaimedEvent struct {
// 	MsgSender        common.Address
// 	RecipientAddr    common.Address
// 	SwapID           common.Hash
// 	RandomNumberHash common.Hash
// 	RandomNumber     common.Hash
// }

// // ParseClaimEvent ...
// func ParseClaimEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
// 	var ev ClaimedEvent
// 	if err := abi.UnpackIntoInterface(&ev, ClaimEventName, log.Data); err != nil {
// 		return nil, err
// 	}

// 	ev.MsgSender = common.BytesToAddress(log.Topics[1].Bytes())
// 	ev.RecipientAddr = common.BytesToAddress(log.Topics[2].Bytes())
// 	ev.SwapID = common.BytesToHash(log.Topics[3].Bytes())

// 	fmt.Printf("CLAIMED\n")
// 	fmt.Printf("sender addr: %s | ", ev.MsgSender.String())
// 	fmt.Printf("receiver addr: %s | ", ev.RecipientAddr.String())
// 	fmt.Printf("swap id: %s | ", hex.EncodeToString(ev.SwapID[:]))
// 	fmt.Printf("random number hash: %s | ", hex.EncodeToString(ev.RandomNumberHash[:]))
// 	fmt.Printf("random number: %s | \n", hex.EncodeToString(ev.RandomNumber[:]))
// 	return ev, nil
// }

// // ToTxLog ...
// func (ev ClaimedEvent) ToTxLog() *storage.TxLog {
// 	return &storage.TxLog{
// 		Chain:        storage.EthChain,
// 		TxType:       storage.TxTypeClaim,
// 		SwapID:       hex.EncodeToString(ev.SwapID[:]),
// 		SenderAddr:   ev.MsgSender.Hex(),
// 		ReceiverAddr: ev.RecipientAddr.Hex(),
// 	}
// }

// // CreateUnbindEvent ...
// type CreateUnbindEvent struct {
// 	Timestamp        uint64
// 	MsgSender        common.Address
// 	RecipientAddr    common.Address
// 	SwapID           common.Hash
// 	RandomNumberHash common.Hash
// 	RandomNumber     common.Hash
// 	RefundAddr       common.Address
// 	ExpireHeight     *big.Int
// 	OutAmount        *big.Int
// 	LaAmount         *big.Int
// }

// // ParseCreateUnbindEvent ...
// func ParseCreateUnbindEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
// 	var ev CreateUnbindEvent
// 	if err := abi.UnpackIntoInterface(&ev, CreateUnbindEventName, log.Data); err != nil {
// 		return nil, err
// 	}

// 	ev.MsgSender = common.BytesToAddress(log.Topics[1].Bytes())
// 	ev.RecipientAddr = common.BytesToAddress(log.Topics[2].Bytes())
// 	ev.SwapID = common.BytesToHash(log.Topics[3].Bytes())

// 	fmt.Printf("UNBIND\n")
// 	fmt.Printf("sender addr: %s\n", ev.MsgSender.String())
// 	fmt.Printf("receiver addr: %s\n", ev.RecipientAddr.String())
// 	fmt.Printf("swap id: %s \n", hex.EncodeToString(ev.SwapID[:]))
// 	fmt.Printf("timestamp: %d\n", ev.Timestamp)
// 	fmt.Printf("expire height: %d\n", ev.ExpireHeight)
// 	fmt.Printf("erc20 amount: %d\n", ev.OutAmount)
// 	fmt.Printf("la amount: %d\n", ev.LaAmount)
// 	return ev, nil
// }

// // ToTxLog ...
// func (ev CreateUnbindEvent) ToTxLog() *storage.TxLog {
// 	return &storage.TxLog{
// 		Chain:        storage.EthChain,
// 		SwapType:     storage.SwapTypeUnbind,
// 		TxType:       storage.TxTypeCreate,
// 		SwapID:       ev.SwapID.String(),
// 		SenderAddr:   ev.MsgSender.Hex(),
// 		ReceiverAddr: ev.RecipientAddr.Hex(),
// 		Timestamp:    int64(ev.Timestamp),
// 		ExpireHeight: ev.ExpireHeight.Int64(),
// 		InAmount:     ev.LaAmount.String(),
// 		OutAmount:    ev.OutAmount.String(),
// 	}
// }

// // RefundEvent ...
// type RefundEvent struct {
// 	MsgSender        common.Address
// 	RecipientAddr    common.Address
// 	SwapID           common.Hash
// 	RandomNumberHash common.Hash
// }

// // ParseRefundEvent ...
// func ParseRefundEvent(log *types.Log) (ContractEvent, error) {
// 	var ev RefundEvent
// 	ev.MsgSender = common.BytesToAddress(log.Topics[1].Bytes())
// 	ev.RecipientAddr = common.BytesToAddress(log.Topics[2].Bytes())
// 	ev.SwapID = common.BytesToHash(log.Topics[3].Bytes())

// 	fmt.Printf("sender addr: %s", ev.MsgSender.String())
// 	fmt.Printf("swap id: %s", hex.EncodeToString(ev.SwapID[:]))
// 	fmt.Printf("receiver addr: %s", ev.RecipientAddr.String())
// 	fmt.Printf("random number hash: %s", hex.EncodeToString(ev.RandomNumberHash[:]))

// 	return ev, nil
// }

// // ToTxLog ...
// func (ev RefundEvent) ToTxLog() *storage.TxLog {
// 	return &storage.TxLog{
// 		TxType:       storage.TxTypeWorkerRefund,
// 		SwapID:       hex.EncodeToString(ev.SwapID[:]),
// 		SenderAddr:   ev.MsgSender.Hex(),
// 		ReceiverAddr: ev.RecipientAddr.Hex(),
// 	}
// }
