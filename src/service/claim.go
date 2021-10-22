package rlr

import (
	"fmt"
	"time"

	"latoken/relayer-smart-contract/src/service/storage"
	workers "latoken/relayer-smart-contract/src/service/workers"
	"latoken/relayer-smart-contract/src/service/workers/utils"

	"github.com/ethereum/go-ethereum/common"
)

// emitChainSendClaim ...
func (r *RelayerSRV) emitChainSendClaim(swapType storage.SwapType) {
	for {
		swaps := r.storage.GetSwapsByTypeAndStatuses(swapType,
			[]storage.SwapStatus{storage.SwapStatusDepositConfirmed, storage.SwapStatusClaimSent})

		for _, swap := range swaps {
			if swap.Status == storage.SwapStatusDepositConfirmed {
				r.logger.Info(fmt.Sprintf("attempting to send claim for swap"))
				if _, err := r.sendClaim(swapType, r.laWorker, swap); err != nil {
					r.logger.Errorf("submit claim failed: %s", err)
				}
			} else {
				r.handleTxSent(r.laWorker.GetChain(), swap, storage.TxTypeClaim,
					storage.SwapStatusClaimConfirmed, storage.SwapStatusClaimSentFailed)
			}
		}

		time.Sleep(2 * time.Second)
	}
}

// ethSendClaim ...
func (r *RelayerSRV) sendClaim(direction storage.SwapType, worker workers.IWorker, swap *storage.Swap) (string, error) {
	txSent := &storage.TxSent{
		Chain:      worker.GetChain(),
		Type:       storage.TxTypeClaim,
		SwapID:     swap.SwapID,
		CreateTime: time.Now().Unix(),
	}

	fmt.Printf("Vote parameters: depositNonce(%d) | datahash(%s) | resourceID(%s)\n", swap.DepositNonce, swap.DataHash, swap.ResourceID)

	txHash, err := worker.Vote(swap.DepositNonce, utils.StringToBytes32(swap.ResourceID), common.Hex2Bytes(swap.DataHash))
	if err != nil {
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusFailed
		r.storage.UpdateSwapStatus(swap, storage.SwapStatusClaimSentFailed, "", "")
		r.storage.CreateTxSent(txSent)
		return "", fmt.Errorf("could not send claim tx: %w", err)
	}
	txSent.TxHash = txHash
	r.storage.UpdateSwapStatus(swap, storage.SwapStatusClaimSent, "", "")

	r.logger.Infof("send claim tx success | chain=%s, swap_ID=%s, tx_hash=%s", worker.GetChain(),
		swap.SwapID, txSent.TxHash)
	// create new tx(claimed)
	r.storage.CreateTxSent(txSent)

	return txSent.TxHash, nil
}
