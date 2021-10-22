package rlr

// import (
// 	"fmt"
// 	"time"

// 	"latoken/relayer-smart-contract/src/service/storage"
// )

// // emitChainSendClaim ...
// func (r *RelayerSRV) emitChainPassed(swapType storage.SwapType) {
// 	for {
// 		swaps := r.storage.GetSwapsByTypeAndStatuses(swapType,
// 			[]storage.SwapStatus{storage.SwapStatusClaimConfirmed, storage.SwapStatusClaimSent})

// 		for _, swap := range swaps {
// 			if swap.Status == storage.SwapStatusDepositConfirmed {
// 				r.logger.Info(fmt.Sprintf("attempting to send claim for swap"))
// 				if _, err := r.sendClaim(swapType, r.laWorker, swap); err != nil {
// 					r.logger.Errorf("submit claim failed: %s", err)
// 				}
// 			} else {
// 				r.handleTxSent(r.laWorker.GetChain(), swap, storage.Tx,
// 					storage.SwapStatusClaimConfirmed, storage.SwapStatusClaimSentFailed)
// 			}
// 		}

// 		time.Sleep(2 * time.Second)
// 	}
// }
