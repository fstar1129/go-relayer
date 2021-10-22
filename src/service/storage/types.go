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
	TxTypeCreate  TxType = "CREATE"
	TxTypeClaim   TxType = "CLAIM"
	TxTypePassed  TxType = "PASSED"
	// SPENDING BIND/UNBIND
	TxTypeSpend TxType = "SPEND"
	// REFUND
	TxTypeWorkerRefund TxType = "WORKER_REFUND"
	TxTypeLaRefund     TxType = "LA_REFUND"
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
	// CLAIM
	SwapStatusClaimSent       SwapStatus = "CLAIM_SENT"
	SwapStatusClaimConfirmed  SwapStatus = "CLAIM_CONFIRMED"
	SwapStatusClaimSentFailed SwapStatus = "CLAIM_SENT_FAILED"
	// PASSED
	// SwapStatusPassedReceived  SwapStatus = "PASSED_RECEIVED"
	SwapStatusPassedConfirmed SwapStatus = "PASSED_CONFIRMED"
	// WORKER SPEND
	SwapStatusSpendSent       SwapStatus = "SPEND_SENT"
	SwapStatusSpendConfirmed  SwapStatus = "SPEND_CONFIRMED"
	SwapStatusSpendSentFailed SwapStatus = "SPEND_SENT_FAILED"
	// // BIND
	// SwapStatusBindSent       SwapStatus = "BIND_SENT"
	// SwapStatusBindConfirmed  SwapStatus = "BIND_CONFIRMED"
	// SwapStatusBindSentFailed SwapStatus = "BIND_SENT_FAILED"
	// // UNBIND
	// SwapStatusUnbindSent       SwapStatus = "UNBIND_SENT"
	// SwapStatusLUnbindConfirmed  SwapStatus = "UNBIND_CONFIRMED"
	// SwapStatusUnbindSentFailed SwapStatus = "UNBIND_SENT_FAILED"
	// // REFUND
	SwapStatusWorkerRefundSent       SwapStatus = "WORKER_REFUND_SENT"
	SwapStatusWorkerRefundConfirmed  SwapStatus = "WORKER_REFUND_CONFIRMED"
	SwapStatusWorkerRefundSentFailed SwapStatus = "WORKER_REFUND_SENT_FAILED"
	// REFUND
	SwapStatusLaRefundSent       SwapStatus = "LA_REFUND_SENT"
	SwapStatusLaRefundConfirmed  SwapStatus = "LA_REFUND_CONFIRMED"
	SwapStatusLaRefundSentFailed SwapStatus = "LA_REFUND_SENT_FAILED"

	SwapStatusRejected SwapStatus = "REJECTED"
)

// TxLogStatus ...
type TxLogStatus string

const (
	TxStatusInit      TxLogStatus = "INIT"
	TxStatusConfirmed TxLogStatus = "CONFIRMED"
)
