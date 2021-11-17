package eth

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	labr "gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/workers/eth-compatible/abi/bridge/la"
)

// ParseLaDepositEvent ...
func ParseLaDepositEvent(log *types.Log) (ContractEvent, error) {
	var ev DepositEvent
	abi, _ := abi.JSON(strings.NewReader(labr.BridgeABI))
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

// ParseLaProposalVote ...
func ParseLaProposalVote(log *types.Log) (ContractEvent, error) {
	var ev ProposalVoteEvent
	abi, _ := abi.JSON(strings.NewReader(labr.BridgeABI))
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

//
// ParseLaProposalEvent ...
func ParseLaProposalEvent(log *types.Log) (ContractEvent, error) {
	var ev ProposalEvent
	abi, _ := abi.JSON(strings.NewReader(labr.BridgeABI))
	if err := abi.UnpackIntoInterface(&ev, ProposalEventName, log.Data); err != nil {
		return nil, err
	}
	// ev.OriginChainID = log.Topics[1].Bytes()
	// ev.DepositNonce = big.NewInt(0).SetBytes(log.Topics[2].Bytes()).Uint64()
	// ev.Status = uint8(big.NewInt(0).SetBytes(log.Topics[3].Bytes()).Uint64())

	fmt.Printf("ProposalEvent\n")
	fmt.Printf("origin chain ID: 0x%s\n", common.Bytes2Hex(ev.OriginChainID[:]))
	fmt.Printf("deposit nonce: %d\n", ev.DepositNonce)
	fmt.Printf("status: %d\n", ev.Status)
	fmt.Printf("resource ID: 0x%s\n", common.Bytes2Hex(ev.ResourceID[:]))
	fmt.Printf("DataHash: 0x%s\n\n", common.Bytes2Hex(ev.DataHash[:]))

	return ev, nil
}
