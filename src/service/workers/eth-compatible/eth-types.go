package eth

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ethbr "github.com/LATOKEN/relayer-smart-contract.git/src/service/workers/eth-compatible/abi/bridge/eth"
	"github.com/LATOKEN/relayer-smart-contract.git/src/service/workers/utils"
)

// ParseDepositEvent ...
func ParseEthDepositEvent(log *types.Log) (ContractEvent, error) {
	var ev DepositEvent
	abi, _ := abi.JSON(strings.NewReader(ethbr.EthbrABI))
	if err := abi.UnpackIntoInterface(&ev, DepositEventName, log.Data); err != nil {
		return nil, err
	}

	ev.DestinationChainID = utils.BytesToBytes8(log.Topics[1].Bytes())
	ev.ResourceID = utils.BytesToBytes32(log.Topics[2].Bytes())
	ev.DepositNonce = big.NewInt(0).SetBytes(log.Topics[3].Bytes()).Uint64()

	fmt.Printf("Deposited\n")
	fmt.Printf("destination chain ID: 0x%s\n", common.Bytes2Hex(ev.DestinationChainID[:]))
	fmt.Printf("resource ID: 0x%s\n", common.Bytes2Hex(ev.ResourceID[:]))
	fmt.Printf("deposit nonce: %d\n", ev.DepositNonce)
	fmt.Printf("depositor address: %s\n", ev.Depositor.Hex())
	fmt.Printf("recipient address: %s\n", ev.RecipientAddress.Hex())
	fmt.Printf("token address: %s\n", ev.TokenAddress.Hex())
	fmt.Printf("amount : %s\n", ev.Amount.String())
	fmt.Printf("DataHASH : %s\n", common.Bytes2Hex(ev.DataHash[:]))

	return ev, nil
}

// ParseProposalEvent ...
func ParseEthProposalEvent(log *types.Log) (ContractEvent, error) {
	var ev ProposalEvent
	abi, _ := abi.JSON(strings.NewReader(ethbr.EthbrABI))
	if err := abi.UnpackIntoInterface(&ev, ProposalEventName, log.Data); err != nil {
		return nil, err
	}
	ev.OriginChainID = utils.BytesToBytes8(log.Topics[1].Bytes())
	ev.DestinationChainID = utils.BytesToBytes8(log.Topics[2].Bytes())
	ev.RecipientAddress = common.BytesToAddress(log.Topics[3].Bytes())

	fmt.Printf("ProposalEvent\n")
	fmt.Printf("destination chain ID: 0x%s\n", common.Bytes2Hex(ev.DestinationChainID[:]))
	fmt.Printf("deposit nonce: %d\n", ev.DepositNonce)
	fmt.Printf("status: %d\n", ev.Status)
	fmt.Printf("resource ID: 0x%s\n", common.Bytes2Hex(ev.ResourceID[:]))
	fmt.Printf("DataHash: 0x%s\n\n", common.Bytes2Hex(ev.DataHash[:]))

	return ev, nil
}
