package storage

// BlockLog ...
type BlockLog struct {
	Chain      string    `gorm:"type:TEXT"`
	BlockHash  string    `gorm:"type:TEXT"`
	ParentHash string    `gorm:"type:TEXT"`
	Height     int64     `gorm:"type:BIGINT"`
	BlockTime  int64     `gorm:"type:BIGINT"`
	Type       BlockType `gorm:"block_type"`
	CreateTime int64     `gorm:"type:BIGINT"`
}

// TxLog ...
type TxLog struct {
	SwapID string `gorm:"primaryKey"`
	Chain  string `gorm:"type:TEXT"`
	// swap id should be hex encoded bytes without '0x' prefix
	TxType       TxType `gorm:"type:tx_types"`
	TxHash       string `gorm:"type:TEXT"`
	InTokenAddr  string `gorm:"type:TEXT"`
	OutTokenAddr string `gorm:"type:TEXT"`
	// sender address should be encoded by each relayer
	SenderAddr string `gorm:"type:TEXT"`
	// receiver address should be encoded by each relayer
	ReceiverAddr       string      `gorm:"type:TEXT"`
	SenderWorkerChain  string      `gorm:"type:TEXT"`
	WorkerChainAddr    string      `gorm:"type:TEXT"`
	InAmount           string      `gorm:"type:TEXT"`
	OutAmount          string      `gorm:"type:TEXT"`
	OriginСhainID      string      `gorm:"type:TEXT"`
	DestinationChainID string      `gorm:"type:TEXT"`
	DepositNonce       uint64      `gorm:"type:BIGINT"`
	ResourceID         string      `gorm:"type:TEXT"`
	SwapStatus         SwapStatus  `gorm:"type:TEXT"`
	ExpireHeight       int64       `gorm:"type:BIGINT"`
	Timestamp          int64       `gorm:"type:BIGINT"`
	BlockHash          string      `gorm:"type:TEXT"`
	Height             int64       `gorm:"type:BIGINT"`
	Status             TxLogStatus `gorm:"type:tx_log_statuses"`
	ConfirmedNum       int64       `gorm:"type:BIGINT"`
	CreateTime         int64       `gorm:"type:BIGINT"`
	UpdateTime         int64       `gorm:"type:BIGINT"`
}

// Swap ...
type Swap struct {
	SwapID             string `gorm:"primaryKey"`
	Type               SwapType
	DestinationChainID string
	OriginChainID      string
	SenderAddr         string
	ReceiverAddr       string
	//	WorkerChainAddr   string
	InTokenAddr      string
	OutTokenAddr     string
	InAmount         string
	OutAmount        string
	RelayerOutAmount string
	DepositNonce     uint64
	ResourceID       string
	// ExpireHeight     int64
	Height int64
	//	Timestamp  int64
	Status SwapStatus
	//	SwapStatus uint8
	CreateTime int64
	UpdateTime int64
	TxHash     string
}

// TxSent ...
type TxSent struct {
	ID         int64    `json:"id"`
	Chain      string   `json:"chain" gorm:"type:TEXT"`
	SwapID     string   `json:"swap_id" gorm:"type:TEXT"`
	Type       TxType   `json:"type" gorm:"type:tx_types"`
	TxHash     string   `json:"tx_hash" gorm:"type:TEXT"`
	ErrMsg     string   `json:"err_msg" gorm:"type:TEXT"`
	Status     TxStatus `json:"status" gorm:"type:tx_statuses"`
	CreateTime int64    `json:"create_time" gorm:"type:BIGINT"`
	UpdateTime int64    `json:"update_time" gorm:"type:BIGINT"`
}

type ResourceId struct {
	Name string `gorm:"primaryKey"`
	ID   string `gorm:"type:TEXT"`
}
