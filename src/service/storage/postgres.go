package storage

import (
	"github.com/jinzhu/gorm"
)

// DataBase is a struct for work with postgres database via gorm
type DataBase struct {
	db *gorm.DB
}

// InitStorage ...
func InitStorage(db *gorm.DB) (*DataBase, error) {
	// create types if not exist
	// 'tx_logs'
	if err := db.Exec(createTxLogStatusIfNotExists).Error; err != nil {
		return nil, err
	}

	// 'tx_types'
	if err := db.Exec(createTxTypeIfNotExists).Error; err != nil {
		return nil, err
	}

	// 'tx_statuses'
	if err := db.Exec(createTxStatusIfNotExists).Error; err != nil {
		return nil, err
	}

	// 'swap_types'
	if err := db.Exec(createSwapTypeIfNotExists).Error; err != nil {
		return nil, err
	}

	// migrate table "block_log"
	if err := db.AutoMigrate(BlockLog{}).Error; err != nil {
		return nil, err
	}

	// migrate table "tx_log"
	if err := db.AutoMigrate(TxLog{}).Error; err != nil {
		return nil, err
	}

	// migrate table "swap"
	if err := db.AutoMigrate(Swap{}).Error; err != nil {
		return nil, err
	}

	// migrate table "tx_sent"
	if err := db.AutoMigrate(TxSent{}).Error; err != nil {
		return nil, err
	}

	// migrate table "resource_ids"
	if err := db.AutoMigrate(ResourceId{}).Error; err != nil {
		return nil, err
	}

	return &DataBase{db: db}, nil
}
