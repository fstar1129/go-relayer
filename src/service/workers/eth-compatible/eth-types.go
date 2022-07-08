package eth

import (
	"fmt"
	"math/big"
	"strings"

	ethbr "github.com/LATOKEN/relayer-smart-contract.git/src/service/workers/eth-compatible/abi/bridge/eth"
	"github.com/LATOKEN/relayer-smart-contract.git/src/service/workers/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ParseDepositEvent ...
func (w *Erc20Worker) ParseEthDepositEvent(log *types.Log) (ContractEvent, error) {
	var ev DepositEvent
	abi, _ := abi.JSON(strings.NewReader(ethbr.EthbrABI))
	if err := abi.UnpackIntoInterface(&ev, DepositEventName, log.Data); err != nil {
		return nil, err
	}

	ev.DestinationChainID = utils.BytesToBytes8(log.Topics[1].Bytes())
	ev.ResourceID = utils.BytesToBytes32(log.Topics[2].Bytes())
	ev.DepositNonce = big.NewInt(0).SetBytes(log.Topics[3].Bytes()).Uint64()

	var SwapID = utils.CalcutateSwapID(common.Bytes2Hex(ev.OriginChainID[:]), common.Bytes2Hex(ev.DestinationChainID[:]), fmt.Sprint(ev.DepositNonce))
	w.logger.Infof("[%s] Deposited\n", SwapID)
	w.logger.Infof("[%s] destination chain ID: 0x%s\n", SwapID, common.Bytes2Hex(ev.DestinationChainID[:]))
	w.logger.Infof("[%s] resource ID: 0x%s\n", SwapID, common.Bytes2Hex(ev.ResourceID[:]))
	w.logger.Infof("[%s] deposit nonce: %d\n", SwapID, ev.DepositNonce)
	w.logger.Infof("[%s] depositor address: %s\n", SwapID, ev.Depositor.Hex())
	w.logger.Infof("[%s] recipient address: %s\n", SwapID, ev.RecipientAddress.Hex())
	w.logger.Infof("[%s] token address: %s\n", SwapID, ev.TokenAddress.Hex())
	w.logger.Infof("[%s] amount : %s\n", SwapID, ev.Amount.String())
	w.logger.Infof("[%s] DataHASH : %s\n", SwapID, common.Bytes2Hex(ev.DataHash[:]))

	return ev, nil
}

// ParseProposalEvent ...
func (w *Erc20Worker) ParseEthProposalEvent(log *types.Log) (ContractEvent, error) {
	var ev ProposalEvent
	abi, _ := abi.JSON(strings.NewReader(ethbr.EthbrABI))
	if err := abi.UnpackIntoInterface(&ev, ProposalEventName, log.Data); err != nil {
		return nil, err
	}
	ev.OriginChainID = utils.BytesToBytes8(log.Topics[1].Bytes())
	ev.DestinationChainID = utils.BytesToBytes8(log.Topics[2].Bytes())
	ev.RecipientAddress = common.BytesToAddress(log.Topics[3].Bytes())

	SwapID := utils.CalcutateSwapID(common.Bytes2Hex(ev.OriginChainID[:]), common.Bytes2Hex(ev.DestinationChainID[:]), fmt.Sprint(ev.DepositNonce))

	fmt.Printf("ProposalEvent\n")
	fmt.Printf("Swap ID: %s \n", SwapID)
	fmt.Printf("destination chain ID: 0x%s\n", common.Bytes2Hex(ev.DestinationChainID[:]))
	fmt.Printf("deposit nonce: %d\n", ev.DepositNonce)
	fmt.Printf("status: %d\n", ev.Status)
	fmt.Printf("resource ID: 0x%s\n", common.Bytes2Hex(ev.ResourceID[:]))
	fmt.Printf("amount: %s", ev.Amount.String())
	fmt.Printf("DataHash: 0x%s\n\n", common.Bytes2Hex(ev.DataHash[:]))

	w.setTxMonitor(SwapID, ev.Status)
	return ev, nil
}
