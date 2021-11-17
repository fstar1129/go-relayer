package rlr

import (
	"fmt"
	"time"

	"gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/storage"
	workers "gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/workers"
	"gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/workers/utils"
)

// emitChainSendClaim ...
func (r *RelayerSRV) emitChainSendClaim(swapType storage.SwapType) {
	for {
		swaps := r.storage.GetSwapsByTypeAndStatuses(swapType,
			[]storage.SwapStatus{storage.SwapStatusDepositConfirmed, storage.SwapStatusClaimSent})

		for _, swap := range swaps {
			if swap.Status == storage.SwapStatusDepositConfirmed {
				r.logger.Info("attempting to send claim for swap")
				if _, err := r.sendClaim(swapType, r.laWorker, swap); err != nil {
					r.logger.Errorf("submit claim failed: %s", err)
				}
			} else {
				r.handleTxSent(r.laWorker.GetChainName(), swap, storage.TxTypeClaim,
					storage.SwapStatusClaimConfirmed, storage.SwapStatusClaimSentFailed)
			}
		}

		time.Sleep(2 * time.Second)
	}
}

// ethSendClaim ...
func (r *RelayerSRV) sendClaim(direction storage.SwapType, worker workers.IWorker, swap *storage.Swap) (string, error) {
	txSent := &storage.TxSent{
		Chain:      worker.GetChainName(),
		Type:       storage.TxTypeClaim,
		SwapID:     swap.SwapID,
		CreateTime: time.Now().Unix(),
	}

	r.logger.Infof("claim parameters: depositNonce(%d) | sender(%s) | outAmount(%s) | resourceID(%s)\n",
		swap.DepositNonce, swap.SenderAddr, swap.OutAmount, swap.ResourceID)

	txHash, err := worker.Vote(swap.DepositNonce, utils.StringToBytes8(swap.ChainID),
		utils.StringToBytes32(swap.ResourceID), swap.ReceiverAddr, swap.OutAmount)
	if err != nil {
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusFailed
		r.storage.UpdateSwapStatus(swap, storage.SwapStatusClaimSentFailed, "")
		r.storage.CreateTxSent(txSent)
		return "", fmt.Errorf("could not send claim tx: %w", err)
	}
	txSent.TxHash = txHash
	r.storage.UpdateSwapStatus(swap, storage.SwapStatusClaimSent, "")

	r.logger.Infof("send claim tx success | chain=%s, swap_ID=%s, tx_hash=%s", worker.GetChainName(),
		swap.SwapID, txSent.TxHash)
	// create new tx(claimed)
	r.storage.CreateTxSent(txSent)

	return txSent.TxHash, nil
}
