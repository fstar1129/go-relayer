package rlr

// import (
// 	"fmt"
// 	"time"

// 	"latoken/relayer-smart-contract/src/service/storage"
// 	workers "latoken/relayer-smart-contract/src/service/workers"

// 	"github.com/ethereum/go-ethereum/common"
// )

// // emitCreateRequest ...
// func (r *RelayerSRV) emitCreateRequest(worker workers.IWorker) {
// 	swapType := storage.SwapTypeBind
// 	if worker.GetChain() == storage.LaChain {
// 		// for BTC
// 		for {
// 			swaps := r.storage.GetSwapsByTypeAndStatuses(swapType,
// 				[]storage.SwapStatus{storage.SwapStatusHTLTConfirmed, storage.SwapStatusCreateSent})

// 			for _, swap := range swaps {
// 				if swap.Status == storage.SwapStatusHTLTConfirmed {
// 					r.logger.Info(fmt.Sprintf("attempting to send create swap request | worker_addr=%s", worker.GetWorkerAddress()))
// 					if _, err := r.sendCreateRequest(r.laWorker, swap); err != nil {
// 						r.logger.Errorf("submit create swap request failed: %s", err)
// 					}
// 				} else {
// 					r.handleTxSent(swap, r.laWorker.GetChain(), storage.TxTypeCreate,
// 						storage.SwapStatusCreateConfirmed, storage.SwapStatusCreateSentFailed)
// 				}
// 			}

// 			time.Sleep(2 * time.Second)
// 		}
// 	}

// 	// for {
// 	// 	swaps := r.storage.GetSwapsByTypeAndStatuses(swapType,
// 	// 		[]storage.SwapStatus{storage.SwapStatusHTLTConfirmed, storage.SwapStatusCreateSent})

// 	// 	for _, swap := range swaps {
// 	// 		if swap.Status == storage.SwapStatusHTLTConfirmed {
// 	// 			r.logger.Info(fmt.Sprintf("attempting to send create swap request | worker_addr=%s", worker.GetWorkerAddress()))
// 	// 			if _, err := r.sendCreateRequest(r.laWorker, swap); err != nil {
// 	// 				r.logger.Errorf("submit create swap request failed: %s", err)
// 	// 			}
// 	// 		} else {
// 	// 			r.handleTxSent(swap, r.laWorker.GetChain(), storage.TxTypeCreate,
// 	// 				storage.SwapStatusCreateConfirmed, storage.SwapStatusCreateSentFailed)
// 	// 		}
// 	// 	}

// 	// 	time.Sleep(2 * time.Second)
// 	// }
// }

// func (r *RelayerSRV) sendCreateRequest(worker workers.IWorker, swap *storage.Swap) (string, error) {
// 	var txHash string
// 	workerChainSwapID := common.HexToHash(swap.WorkerChainSwapID)

// 	txSent := &storage.TxSent{
// 		Chain:      worker.GetChain(),
// 		SwapID:     swap.WorkerChainSwapID,
// 		CreateTime: time.Now().Unix(),
// 	}

// 	txSent.Type = storage.TxTypeCreate
// 	txHash, err := worker.CreateRequest(workerChainSwapID)
// 	if err != nil {
// 		txSent.ErrMsg = err.Error()
// 		txSent.Status = storage.TxSentStatusFailed
// 		r.storage.UpdateSwapStatus(swap, storage.SwapStatusCreateSentFailed, "")
// 		r.storage.CreateTxSent(txSent)
// 		return "", fmt.Errorf("could not send createRequest tx: %v", err)
// 	}
// 	txSent.TxHash = txHash
// 	r.storage.UpdateSwapStatus(swap, storage.SwapStatusCreateSent, "")

// 	r.logger.Infof("send chain %s create swap request tx success | worker_chain_swap_ID=%s | tx_hash=%s", worker.GetChain(),
// 		workerChainSwapID.String(), txHash)
// 	// create new tx(create request)
// 	r.storage.CreateTxSent(txSent)

// 	return txHash, nil
// }

// // // workerSendCreateRequest ...
// // func (r *RelayerSRV) workerSendCreateRequest(worker workers.IWorker) {
// // 	for {
// // 		swaps := r.storage.GetSwapsByTypeAndStatuses(storage.SwapTypeUnbind,
// // 			[]storage.SwapStatus{storage.SwapStatusHTLTConfirmed, storage.SwapStatusCreateSent})

// // 		for _, swap := range swaps {
// // 			if swap.Status == storage.SwapStatusHTLTConfirmed {
// // 				r.logger.Info(fmt.Sprintf("attempting to send create swap request | worker_addr=%s", worker.GetWorkerAddress()))
// // 				if _, err := r.sendCreateRequest(storage.SwapTypeUnbind, r.laWorker, swap); err != nil {
// // 					r.logger.Errorf("submit create swap request failed: %s", err)
// // 				}
// // 			} else {
// // 				r.handleTxSent(swap, r.laWorker.GetChain(), storage.TxTypeCreate,
// // 					storage.SwapStatusCreateConfirmed, storage.SwapStatusCreateSentFailed)
// // 			}
// // 		}

// // 		time.Sleep(2 * time.Second)
// // 	}
// // }
