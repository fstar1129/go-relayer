package rlr

import (
	"time"

	"gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/storage"
)

// emitChainSendClaim ...
func (r *RelayerSRV) emitChainSendSpend() {
	for {
		swaps := r.storage.GetSwapsByTypeAndStatuses(
			[]storage.SwapStatus{storage.SwapStatusSpendConfirmed})

		for _, swap := range swaps {
			r.logger.Info("updating spend status")
			println()
			r.storage.UpdateSwapStatus(swap, storage.SwapStatusSpendSent, "")
		}

		time.Sleep(2 * time.Second)
	}
}
