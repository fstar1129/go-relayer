package eth

import (
	"fmt"
	"strings"
	"time"

	labr "github.com/LATOKEN/relayer-smart-contract.git/src/service/workers/eth-compatible/abi/bridge/la"
	"github.com/LATOKEN/relayer-smart-contract.git/src/service/workers/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var txStatus = make(map[string]uint8)

func (w *Erc20Worker) setTxMonitor(SwapID string, Status uint8) {
	if Status == 1 {
		txStatus[SwapID] = Status
		go func(SwapID string, Status uint8) {
			time.Sleep(5 * 60 * time.Second)
			if NewStatus, ok := txStatus[SwapID]; ok {
				if NewStatus != 3 {
					w.logger.Errorf("[%s] stuck in status %d\n\n", SwapID, NewStatus)
				}
				delete(txStatus, SwapID)
			}
		}(SwapID, Status)
	} else {
		if HashStatus, ok := txStatus[SwapID]; ok {
			if HashStatus < Status {
				txStatus[SwapID] = Status
			}
		}
	}
}

// ParseLaDepositEvent ...
func (w *Erc20Worker) ParseLaDepositEvent(log *types.Log) (ContractEvent, error) {
	var ev DepositEvent
	abi, _ := abi.JSON(strings.NewReader(labr.LabrABI))
	if err := abi.UnpackIntoInterface(&ev, DepositEventName, log.Data); err != nil {
		return nil, err
	}

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

// ParseLaProposalEvent ...
func (w *Erc20Worker) ParseLaProposalEvent(log *types.Log) (ContractEvent, error) {
	var ev ProposalEvent
	abi, _ := abi.JSON(strings.NewReader(labr.LabrABI))
	if err := abi.UnpackIntoInterface(&ev, ProposalEventName, log.Data); err != nil {
		return nil, err
	}
	// ev.OriginChainID = log.Topics[1].Bytes()
	// ev.DepositNonce = big.NewInt(0).SetBytes(log.Topics[2].Bytes()).Uint64()
	// ev.Status = uint8(big.NewInt(0).SetBytes(log.Topics[3].Bytes()).Uint64())

	fmt.Printf("ProposalEvent\n")
	fmt.Printf("origin chain ID: 0x%s\n", common.Bytes2Hex(ev.OriginChainID[:]))
	fmt.Printf("Swap ID:%s\n", utils.CalcutateSwapID(common.Bytes2Hex(ev.OriginChainID[:]), common.Bytes2Hex(ev.DestinationChainID[:]), fmt.Sprint(ev.DepositNonce)))
	fmt.Printf("destination chain ID: 0x%s\n", common.Bytes2Hex(ev.DestinationChainID[:]))
	fmt.Printf("deposit nonce: %d\n", ev.DepositNonce)
	fmt.Printf("status: %d\n", ev.Status)
	fmt.Printf("resource ID: 0x%s\n", common.Bytes2Hex(ev.ResourceID[:]))
	fmt.Printf("DataHash: 0x%s\n\n", common.Bytes2Hex(ev.DataHash[:]))


	w.setTxMonitor(SwapID, ev.Status)
	return ev, nil
}
