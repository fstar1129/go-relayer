package rlr

import (
	"fmt"
	"strings"
	"time"

	"github.com/LATOKEN/relayer-smart-contract.git/src/service/storage"
	workers "github.com/LATOKEN/relayer-smart-contract.git/src/service/workers"
	"github.com/LATOKEN/relayer-smart-contract.git/src/service/workers/utils"
)

// emitChainSendClaim ...
func (r *RelayerSRV) emitChainSendClaim() {
	for {
		swaps := r.storage.GetSwapsByTypeAndStatuses(
			[]storage.SwapStatus{storage.SwapStatusDepositConfirmed, storage.SwapStatusClaimConfirmed, storage.SwapStatusDepositInit, storage.SwapStatusClaimSentFailed})
		for _, swap := range swaps {
			if swap.Status == storage.SwapStatusDepositConfirmed {
				r.logger.Info("attempting to send claim for swap")
				if _, err := r.sendClaim(r.laWorker, swap); err != nil {
					r.logger.Errorf("submit claim failed: %s", err)
				}
				time.Sleep(2 * time.Second)
			} else {
				r.handleTxSent(r.laWorker.GetChainName(), swap, storage.TxTypeDeposit,
					storage.SwapStatusDepositConfirmed, storage.SwapStatusDepositFailed)
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
	var originWorker workers.IWorker
	var destWorker workers.IWorker
	for _, wrkr := range r.Workers {
		if strings.ToLower(wrkr.GetDestinationID()) == strings.ToLower(swap.OriginChainID) {
			originWorker = wrkr
		}
		if strings.ToLower(wrkr.GetDestinationID()) == strings.ToLower(swap.DestinationChainID) {
			destWorker = wrkr
		}
	}

	if originWorker == nil || destWorker == nil {
		err := "Missing worker"
		println(err)
		txSent.ErrMsg = err
		txSent.Status = storage.TxSentStatusFailed
		r.storage.UpdateSwapStatus(swap, storage.SwapStatusDepositFailed, "")
		return "", fmt.Errorf("could not send claim tx: %s", err)
	}

	originDecimals, err := originWorker.GetDecimalsFromResourceID(swap.ResourceID)
	if err != nil {
		println("error in decimals", err.Error())
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusFailed
		r.storage.UpdateSwapStatus(swap, storage.SwapStatusDepositFailed, "")
		return "", fmt.Errorf("could not send claim tx: %w", err)
	}

	destDecimals, err := destWorker.GetDecimalsFromResourceID(swap.ResourceID)
	if err != nil {
		println("error in decimals", err.Error())
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusFailed
		r.storage.UpdateSwapStatus(swap, storage.SwapStatusDepositFailed, "")
		return "", fmt.Errorf("could not send claim tx: %w", err)
	}
	var amount string
	if originDecimals == destDecimals {
		amount = swap.OutAmount
	} else if originDecimals == 0 || destDecimals == 0 || originDecimals > 63 || destDecimals > 63 {
		err = fmt.Errorf("One of decimals is zero or greater than 63")
		println("error in decimals", err.Error())
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusFailed
		r.storage.UpdateSwapStatus(swap, storage.SwapStatusDepositFailed, "")
		return "", fmt.Errorf("could not send claim tx: %w", err)
	} else {
		amount = utils.ConvertDecimals(originDecimals, destDecimals, swap.OutAmount)
	}
	r.logger.Infof("claim parameters: depositNonce(%d) | sender(%s) | outAmount(%d) | resourceID(%s)\n",
		swap.DepositNonce, swap.SenderAddr, amount, swap.ResourceID)

	txHash, err := worker.Vote(swap.DepositNonce, utils.StringToBytes8(swap.OriginChainID), utils.StringToBytes8(swap.DestinationChainID),
		utils.StringToBytes32(swap.ResourceID), swap.ReceiverAddr, amount)
	if err != nil {
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusNotFound
		r.storage.UpdateSwapStatus(swap, storage.SwapStatusClaimSentFailed, "")
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
