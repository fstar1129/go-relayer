package storage

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// GetConfirmedTxsLog ...
func (d *DataBase) GetConfirmedTxsLog(tx *gorm.DB, chain string, swap *Swap) (txLogs []*TxLog, err error) {
	query := tx.Where("chain = ? and status = ? and swap_id = ?", chain, TxStatusConfirmed, swap.SwapID)
	if err := query.Order("id desc").Find(&txLogs).Error; err != nil {
		return txLogs, err
	}

	return txLogs, nil
}

// FindTxLogs ...
func (d *DataBase) FindTxLogs(chainID string, confirmNum int64) (txLogs []*TxLog, err error) {
	if err := d.db.Where("chain = ? and status = ? and confirmed_num >= ?",
		chainID, TxStatusInit, confirmNum).Find(&txLogs).Error; err != nil {
		return nil, err
	}

	return txLogs, nil
}

// ConfirmWorkerTx ...
func (d *DataBase) ConfirmWorkerTx(chainID string, txLogs []*TxLog, txHashes []string, newSwaps []*Swap) error {
	tx := d.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	err := tx.Model(TxLog{}).Where("tx_hash in (?)", txHashes).Updates(
		map[string]interface{}{
			"status":      TxStatusConfirmed,
			"update_time": time.Now().Unix(),
		}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// create swap
	for _, swap := range newSwaps {
		if err := tx.Create(swap).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, txLog := range txLogs {
		if err := d.ConfirmTx(tx, txLog); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := d.CompensateNewSwap(tx, chainID, newSwaps); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

//  !!! TODO !!!

// ConfirmTx ...
func (d *DataBase) ConfirmTx(tx *gorm.DB, txLog *TxLog) error {
	switch txLog.TxType {
	case TxTypeDeposit:
		// BIND tokens(mainchain(e.g. btc, eth) -> lachain)
		if err := d.UpdateSwapStatusWhenConfirmTx(tx, SwapTypeBind, txLog, []SwapStatus{
			SwapStatusDepositConfirmed, SwapStatusDepositSent, SwapStatusDepositSentFailed},
			nil, SwapStatusDepositConfirmed); err != nil {
			return err
		}

		// UNBIND tokens(lachain -> mainchain(e.g. btc, eth)
		if err := d.UpdateSwapStatusWhenConfirmTx(tx, SwapTypeUnbind, txLog, []SwapStatus{
			SwapStatusDepositConfirmed, SwapStatusDepositSent, SwapStatusDepositSentFailed},
			nil, SwapStatusDepositConfirmed); err != nil {
			return err
		}
	case TxTypeClaim:
		fmt.Println("CLAIM")
		// BIND tokens(mainchain(e.g. btc, eth) -> lachain)
		if err := d.UpdateSwapStatusWhenConfirmTx(tx, SwapTypeBind, txLog, []SwapStatus{
			SwapStatusClaimSent}, nil, SwapStatusClaimConfirmed); err != nil {
			return err
		}

		// UNBIND tokens(lachain -> mainchain(e.g. btc, eth)
		if err := d.UpdateSwapStatusWhenConfirmTx(tx, SwapTypeUnbind, txLog, []SwapStatus{
			SwapStatusClaimSent}, nil, SwapStatusClaimConfirmed); err != nil {
			return err
		}

	case TxTypePassed:
		fmt.Println("PASSSED")
		// BIND tokens(mainchain(e.g. btc, eth) -> lachain)
		if err := d.UpdateSwapStatusWhenConfirmTx(tx, SwapTypeBind, txLog, []SwapStatus{
			SwapStatusClaimConfirmed}, nil, SwapStatusPassedConfirmed); err != nil {
			fmt.Println("XXXXX: ", err)
			return err
		}

		// UNBIND tokens(lachain -> mainchain(e.g. btc, eth)
		if err := d.UpdateSwapStatusWhenConfirmTx(tx, SwapTypeUnbind, txLog, []SwapStatus{
			SwapStatusClaimConfirmed}, nil, SwapStatusPassedConfirmed); err != nil {
			return err
		}

		// // BIND tokens(mainchain(e.g. btc, eth) -> lachain)
		// if err := d.UpdateSwapStatusWhenConfirmTx(tx, SwapTypeBind, txLog, []SwapStatus{
		// 	SwapStatusPassedReceived}, nil, SwapStatusPassedConfirmed); err != nil {
		// 	return err
		// }

		// // UNBIND tokens(lachain -> mainchain(e.g. btc, eth)
		// if err := d.UpdateSwapStatusWhenConfirmTx(tx, SwapTypeUnbind, txLog, []SwapStatus{
		// 	SwapStatusPassedReceived}, nil, SwapStatusPassedConfirmed); err != nil {
		// 	return err
		// }
	case TxTypeSpend:
		// BIND tokens(mainchain(e.g. btc, eth) -> lachain)
		if err := d.UpdateSwapStatusWhenConfirmTx(tx, SwapTypeBind, txLog, []SwapStatus{
			SwapStatusSpendSent}, nil, SwapStatusSpendConfirmed); err != nil {
			return err
		}
		// UNBIND tokens(lachain -> mainchain(e.g. btc, eth)
		if err := d.UpdateSwapStatusWhenConfirmTx(tx, SwapTypeUnbind, txLog, []SwapStatus{
			SwapStatusSpendSent}, nil, SwapStatusSpendConfirmed); err != nil {
			return err
		}
		// case TxTypeLaRefund:
		// 	if err := d.UpdateSwapStatusWhenConfirmTx(tx, SwapTypeUnbind, txLog, nil,
		// 		nil, SwapStatusLaRefundConfirmed); err != nil {
		// 		return err
		// 	}

		// 	if err := d.UpdateSwapStatusWhenConfirmTx(tx, SwapTypeBind, txLog, nil,
		// 		[]SwapStatus{SwapStatusHTLTExpired, SwapStatusWorkerRefundConfirmed},
		// 		SwapStatusLaRefundConfirmed); err != nil {
		// 		return err
		// 	}
		// case TxTypeWorkerRefund:
		// 	if err := d.UpdateSwapStatusWhenConfirmTx(tx, SwapTypeUnbind, txLog, nil,
		// 		nil, SwapStatusWorkerRefundConfirmed); err != nil {
		// 		return err
		// 	}

		// 	if err := d.UpdateSwapStatusWhenConfirmTx(tx, SwapTypeUnbind, txLog, nil,
		// 		[]SwapStatus{SwapStatusHTLTExpired, SwapStatusLaRefundConfirmed},
		// 		SwapStatusWorkerRefundConfirmed); err != nil {
		// 		return err
		// 	}
	}

	return nil
}

// ------ TXSENT ------

// CreateTxSent ...
func (d *DataBase) CreateTxSent(txSent *TxSent) error {
	if txSent.Status == "" {
		txSent.Status = TxSentStatusInit
	}

	return d.db.Create(txSent).Error
}

// UpdateTxSentStatus ...
func (d *DataBase) UpdateTxSentStatus(txSent *TxSent, status TxStatus) error {
	return d.db.Model(txSent).Update(
		map[string]interface{}{
			"status":      status,
			"update_time": time.Now().Unix(),
		}).Error
}

// GetTxsSentByStatus ...
func (d *DataBase) GetTxsSentByStatus(chain string) ([]*TxSent, error) {
	txsSent := make([]*TxSent, 0)
	status := []TxStatus{TxSentStatusInit, TxSentStatusNotFound, TxSentStatusPending}
	if err := d.db.Where("chain = ? and status in (?)", chain, status).Find(&txsSent).Error; err != nil {
		return nil, err
	}

	return txsSent, nil
}

// GetTxsSentByType ...
func (d *DataBase) GetTxsSentByType(chain string, txType TxType, swap *Swap) []*TxSent {
	txsSent := make([]*TxSent, 0)
	query := d.db.Where("chain = ? and type = ?", chain, txType)
	query = query.Where("swap_id = ?", swap.SwapID)
	query.Order("id desc").Find(&txsSent)

	return txsSent
}
