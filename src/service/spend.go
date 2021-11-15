package rlr

import (
	"fmt"
	"time"

	"gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/storage"
	workers "gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/workers"
	"gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/workers/utils"
)

// emitChainSendSpend ...
func (r *RelayerSRV) emitChainSendSpend(worker workers.IWorker) {
	swapType := storage.SwapTypeUnbind
	if worker.GetChain() == storage.LaChain {
		swapType = storage.SwapTypeBind
	}

	for {
		swaps := r.storage.GetSwapsByTypeAndStatuses(swapType,
			[]storage.SwapStatus{storage.SwapStatusPassedConfirmed, storage.SwapStatusSpendSent})

		for _, swap := range swaps {
			if swap.Status == storage.SwapStatusPassedConfirmed {
				r.logger.Infof("attempting to send SPEND for swap | worker_addr=%s\n", worker.GetWorkerAddress())
				if _, err := r.sendSpend(swapType, worker, swap); err != nil {
					r.logger.Errorf("submit spend failed: %s", err)
				}
			} else {
				r.handleTxSent(worker.GetChain(), swap, storage.TxTypeSpend,
					storage.SwapStatusSpendConfirmed, storage.SwapStatusSpendSentFailed)
			}
		}
		time.Sleep(2 * time.Second)
	}
}

func (r *RelayerSRV) sendSpend(direction storage.SwapType, worker workers.IWorker, swap *storage.Swap) (txHash string, err error) {
	txSent := &storage.TxSent{
		Chain:      worker.GetChain(),
		Type:       storage.TxTypeSpend,
		SwapID:     swap.SwapID,
		CreateTime: time.Now().Unix(),
	}

	r.logger.Infof("spend parameters: depositNonce(%d) | sender(%s) | outAmount(%s) | resourceID(%s) | chainID(%s)\n",
		swap.DepositNonce, swap.SenderAddr, swap.OutAmount, swap.ResourceID, worker.GetChain())

	if direction == storage.SwapTypeBind {
		txHash, err = worker.SpendBind(swap.DepositNonce, utils.StringToBytes8(swap.ChainID), utils.StringToBytes32(swap.ResourceID),
			swap.ReceiverAddr, swap.OutAmount)
	} else {
		txHash, err = worker.SpendUnbind(swap.DepositNonce, utils.StringToBytes8(swap.ChainID), utils.StringToBytes32(swap.ResourceID),
			swap.ReceiverAddr, swap.OutAmount)
	}
	if err != nil {
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusFailed
		r.storage.UpdateSwapStatus(swap, storage.SwapStatusSpendSentFailed, "")
		r.storage.CreateTxSent(txSent)
		return "", fmt.Errorf("could not send spendUnbind tx: %w", err)
	}
	txSent.TxHash = txHash
	r.storage.UpdateSwapStatus(swap, storage.SwapStatusSpendSent, "")
	r.logger.Infof("send spend tx success, chain=%s, swap_ID=%s, tx_hash=%s", worker.GetChain(), swap.SwapID, txSent.TxHash)
	// create new tx(spend)
	r.storage.CreateTxSent(txSent)

	return txSent.TxHash, nil
}
