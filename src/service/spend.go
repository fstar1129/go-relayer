package rlr

import (
	"time"

	"github.com/LATOKEN/relayer-smart-contract.git/src/service/storage"
)

// emitChainSendClaim ...
func (r *RelayerSRV) emitChainSendSpend() {
	for {
		swaps := r.storage.GetSwapsByTypeAndStatuses(
			[]storage.SwapStatus{storage.SwapStatusSpendConfirmed})

		for _, swap := range swaps {
			r.logger.Info("updating spend status")
			r.storage.UpdateSwapStatus(swap, storage.SwapStatusSpendSent, "")
			r.logger.Infof("send spend tx success | swap_ID=%s",
				swap.SwapID)
		}

		time.Sleep(2 * time.Second)
	}
}
