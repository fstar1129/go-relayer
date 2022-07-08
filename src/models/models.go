package models

import (
	"math/big"
	"time"

	"github.com/LATOKEN/relayer-smart-contract.git/src/service/storage"

	"github.com/ethereum/go-ethereum/common"
)

// ServiceConfig contains configurations for rest-api service.
type ServiceConfig struct {
	Host string // Service Host
	Port string // Service port
}

// RelayerStatus ...
type RelayerStatus struct {
	Mode    string                  `json:"mode"`
	Workers map[string]WorkerStatus `json:"workers"`
}

// BlockAndTxLogs ...
type BlockAndTxLogs struct {
	Height          int64
	BlockHash       string
	ParentBlockHash string
	BlockTime       int64
	TxLogs          []*storage.TxLog
}

// StorageConfig contains configurations for storage, postgreSQL
type StorageConfig struct {
	URL        string // DataBase URL for connection
	DBDriver   string // DataBase driver
	DBHOST     string // DataBase host
	DBPORT     int64
	DBSSL      string // DataBase sslmode
	DBName     string // DataBase name
	DBUser     string // DataBase's user
	DBPassword string // User's password
}

// WorkerConfig ...
type WorkerConfig struct {
	NetworkType                    string         `json:"type"`
	ChainName                      string         `json:"chain_name"`
	User                           string         `json:"user"`
	Password                       string         `json:"password"`
	SwapType                       string         `json:"swap_type"`
	PrivateKey                     string         `json:"private_key"`
	Provider                       string         `json:"provider"`
	ContractAddr                   common.Address `json:"contract_addr"`
	NativeResourceID               string         `json:"native_resource_id"`
	WorkerAddr                     common.Address `json:"worker_addr"`
	ColdWalletAddr                 common.Address `json:"cold_wallet_addr"`
	TokenBalanceAlertThreshold     *big.Int       `json:"token_balance_alert_threshold"`
	EthBalanceAlertThreshold       *big.Int       `json:"eth_balance_alert_threshold"`
	AllowanceBalanceAlertThreshold *big.Int       `json:"allowance_balance_alert_threshold"`
	FetchInterval                  int64          `json:"fetch_interval"`
	GasLimit                       int64          `json:"gas_limit"`
	GasPrice                       *big.Int       `json:"gas_price"`
	ChainDecimal                   int            `json:"chain_decimal"`
	ConfirmNum                     int64          `json:"confirm_num"`
	StartBlockHeight               int64          `json:"start_block_height"`
	DestinationChainID             string         `json:"dest_id"`
}

// SwapStatus ...
type SwapStatus struct {
	Chain   string `json:"chain"`
	Sender  string `json:"sender"`
	Receipt string `json:"receipt"`
	Amount  string `json:"amount"`
	TxHash  string `json:"tx_hash"`
}

// StatusResponce ...
type StatusResponce struct {
	Workers []WorkerStatus `json:"workers"`
}

// WorkerStatus ...
type WorkerStatus struct {
	Height             int64         `json:"height"`
	SyncHeight         int64         `json:"sync_height"`
	LastBlockFetchedAt time.Time     `json:"last_block_fetched_at"`
	Account            WorkerAccount `json:"account"`
}

// WorkerAccount ...
type WorkerAccount struct {
	Address string `json:"address"`
}
