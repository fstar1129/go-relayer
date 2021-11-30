package storage

var (
	LaChain string = "LA"
	// EthChain - Ethereum chain
	EthChain string = "ETH"
	// BtcChain - Bitcoin chain
	BtcChain string = "BTC"
)

// BlockType ...
type BlockType string

const (
	BlockTypeCurrent BlockType = "CURRENT"
	BlockTypeParent  BlockType = "PARENT"
)

// TxStatus ...
type TxStatus string

const (
	TxSentStatusInit     TxStatus = "INIT"
	TxSentStatusNotFound TxStatus = "NOT_FOUND"
	TxSentStatusPending  TxStatus = "PENDING"
	TxSentStatusFailed   TxStatus = "FAILED"
	TxSentStatusSuccess  TxStatus = "SUCCESS"
	TxSentStatusLost     TxStatus = "LOST"
)

//
type TxType string

const (
	TxTypeDeposit TxType = "DEPOSIT"
)

type SwapType string

const (
	SwapTypeBind   SwapType = "BIND"
	SwapTypeUnbind SwapType = "UNBIND"
)

type SwapStatus string

const (
	// DEPOSIT
	SwapStatusDepositSent       SwapStatus = "DEPOSIT_SENT"
	SwapStatusDepositConfirmed  SwapStatus = "DEPOSIT_CONFIRMED"
	SwapStatusDepositFailed     SwapStatus = "DEPOSIT_FAILED"
	SwapStatusDepositSentFailed SwapStatus = "DEPOSIT_SENT_FAILED"

	SwapStatusRejected SwapStatus = "REJECTED"
)

// TxLogStatus ...
type TxLogStatus string

const (
	TxStatusInit      TxLogStatus = "INIT"
	TxStatusConfirmed TxLogStatus = "CONFIRMED"
)
