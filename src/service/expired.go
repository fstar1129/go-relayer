package rlr

import (
	"time"

	"github.com/LATOKEN/relayer-smart-contract.git/src/service/storage"
)

// emitChainSendClaim ...
func (r *RelayerSRV) emitChainSendExpire() {
	for {
		swaps := r.storage.GetSwapsByTypeAndStatuses(
			[]storage.SwapStatus{storage.SwapStatusExpiredConfirmed})
		for _, swap := range swaps {
			r.logger.Info("updating expired status")
			r.storage.UpdateSwapStatus(swap, storage.SwapStatusExpiredSent, "")
			r.logger.Infof("send expired tx success | swap_ID=%s",
				swap.SwapID)
		}

		time.Sleep(2 * time.Second)
	}
}
