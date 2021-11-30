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
	abi, _ := abi.JSON(strings.NewReader(labr.LabrABI))
	if err := abi.UnpackIntoInterface(&ev, DepositEventName, log.Data); err != nil {
		return nil, err
	}

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
