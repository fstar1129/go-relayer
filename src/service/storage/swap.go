package storage

import (
	"time"

	"github.com/jinzhu/gorm"
)

/*
1. CREATE swap IN ./tx.go|ConfirmWorkerTx(...)
2. GET swap
3. UPDATE swap
*/

// GetSwapsByTypeAndStatuses ...
func (d *DataBase) GetSwapsByTypeAndStatuses(statuses []SwapStatus) []*Swap {
	swaps := make([]*Swap, 0)
	if err := d.db.Where("status in (?)", statuses).
		Find(&swaps).Error; err != nil {
		return nil
	}

	return swaps
}

// GetSwapByStatus ...
// func (d *DataBase) GetSwapByStatus(swapType SwapType, sender, receipt string, amount string) (*Swap, error) {
// 	swap := &Swap{}
// 	if err := d.db.Where("type = ? and sender_addr = ? and receiver_addr = ? and out_amount = ? and status in (?)",
// 		swapType, sender, receipt, amount,
// 		[]SwapStatus{SwapStatusClaimSent, SwapStatusClaimConfirmed, SwapStatusClaimSentFailed, SwapStatusDepositConfirmed, SwapStatusDepositFailed, SwapStatusRejected}).
// 		Find(&swap).Error; err != nil {
// 		return nil, err
// 	}

// 	return swap, nil
// }

// UpdateSwapStatus ...
func (d *DataBase) UpdateSwapStatus(swap *Swap, status SwapStatus, rOutAmount string) {
	toUpdate := map[string]interface{}{
		"status":      status,
		"update_time": time.Now().Unix(),
	}

	if rOutAmount != "" {
		toUpdate["r_out_amount"] = rOutAmount
	}

	d.db.Model(swap).Update(toUpdate)
}

// CompensateNewSwap ...
func (d *DataBase) CompensateNewSwap(tx *gorm.DB, chain string, newSwaps []*Swap) error {
	for _, swap := range newSwaps {
		txLogs, err := d.GetConfirmedTxsLog(tx, chain, swap)
		if err != nil {
			continue
		}

		if len(txLogs) == 0 {
			continue
		}

		if err = d.ConfirmTx(tx, txLogs[0]); err != nil {
			return err
		}
	}

	return nil
}

// UpdateSwapStatusWhenConfirmTx ...
func (d *DataBase) UpdateSwapStatusWhenConfirmTx(tx *gorm.DB, swapType SwapType, txLog *TxLog,
	inStatuses []SwapStatus, notInStatuses []SwapStatus, updateStatus SwapStatus) error {
	query := tx.Model(Swap{})

	query = query.Where("swap_id = ?", txLog.SwapID)

	if swapType != "" {
		query = query.Where("type = ?", swapType)
	}

	if len(inStatuses) != 0 {
		query = query.Where("status in (?)", inStatuses)
	}

	if len(notInStatuses) != 0 {
		query = query.Where("status not in (?)", notInStatuses)
	}

	toUpdate := map[string]interface{}{
		"status":      updateStatus,
		"update_time": time.Now().Unix(),
	}

	return query.Updates(toUpdate).Error
}
