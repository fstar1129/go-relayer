package rlr

import (
	"fmt"
	"strconv"
	"time"

	"gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/storage"
	workers "gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/workers"
	"gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/workers/utils"
)

// emitChainSendClaim ...
func (r *RelayerSRV) emitChainSendClaim() {
	for {
		swaps := r.storage.GetSwapsByTypeAndStatuses(
			[]storage.SwapStatus{storage.SwapStatusDepositConfirmed, storage.SwapStatusClaimSentFailed})

		for _, swap := range swaps {
			if swap.Status == storage.SwapStatusDepositConfirmed {
				r.logger.Info("attempting to send claim for swap")
				if _, err := r.sendClaim(r.laWorker, swap); err != nil {
					r.logger.Errorf("submit claim failed: %s", err)
				}
			} else {
				r.handleTxSent(r.laWorker.GetChainName(), swap, storage.TxTypeDeposit,
					storage.SwapStatusClaimConfirmed, storage.SwapStatusDepositConfirmed)
			}
		}

		time.Sleep(2 * time.Second)
	}
}

// ethSendClaim ...
func (r *RelayerSRV) sendClaim(worker workers.IWorker, swap *storage.Swap) (string, error) {
	txSent := &storage.TxSent{
		Chain:      worker.GetChainName(),
		Type:       storage.TxTypeDeposit,
		SwapID:     swap.SwapID,
		CreateTime: time.Now().Unix(),
	}

	var amount int64
	println(swap.OriginChainID, r.Workers[storage.BSCChain].GetDestinationID())
	println(swap.ResourceID, r.storage.FetchResourceID("tether").Name)
	if swap.OriginChainID == r.Workers[storage.BSCChain].GetDestinationID() && swap.ResourceID == r.storage.FetchResourceID("tether").Name {
		println("in origin")
		amount = utils.ConvertDecimals(swap.OutAmount, 18, 6)
	} else if swap.DestinationChainID == r.Workers[storage.BSCChain].GetDestinationID() && swap.ResourceID == r.storage.FetchResourceID("tether").Name {
		println("in dest")
		amount = utils.ConvertDecimals(swap.OutAmount, 6, 18)
	} else {
		amount, _ = strconv.ParseInt(swap.OutAmount, 10, 0)
	}

	r.logger.Infof("claim parameters: depositNonce(%d) | sender(%s) | outAmount(%d) | resourceID(%s)\n",
		swap.DepositNonce, swap.SenderAddr, amount, swap.ResourceID)

	txHash, err := worker.Vote(swap.DepositNonce, utils.StringToBytes8(swap.OriginChainID), utils.StringToBytes8(swap.DestinationChainID),
		utils.StringToBytes32(swap.ResourceID), swap.ReceiverAddr, fmt.Sprint(amount))
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
