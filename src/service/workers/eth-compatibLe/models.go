package eth

import (
	"math"
	"math/big"

	"latoken/relayer-smart-contract/src/service/storage"
)

// Fixed8Decimals ...
var Fixed8Decimals = big.NewInt(int64(math.Pow10(8)))

// DeputyMode ...
type DeputyMode int

// Error ...
type Error struct {
	err       error
	retryable bool
}

// NewError ...
func NewError(err error, retryable bool) *Error {
	return &Error{
		err:       err,
		retryable: retryable,
	}
}

// FailedSwaps ...
type FailedSwaps struct {
	TotalCount int                   `json:"total_count"`
	CurPage    int                   `json:"cur_page"`
	NumPerPage int                   `json:"num_per_page"`
	Swaps      []*storage.SwapStatus `json:"swaps"`
}

// // ReconciliationStatus ...
// type ReconciliationStatus struct {
// 	Bep2TokenBalance          *big.Float
// 	Bep2TokenOutPending       *big.Float
// 	OtherChainTokenBalance    *big.Float
// 	OtherChainTokenOutPending *big.Float
// }

// // DeputyStatus ...
// type DeputyStatus struct {
// 	Mode                         DeputyMode `json:"mode"`
// 	BnbChainHeight               int64      `json:"bnb_chain_height"`
// 	BnbSyncHeight                int64      `json:"bnb_sync_height"`
// 	OtherChainHeight             int64      `json:"other_chain_height"`
// 	OtherChainSyncHeight         int64      `json:"other_chain_sync_height"`
// 	BnbChainLastBlockFetchedAt   time.Time  `json:"bnb_chain_last_block_fetched_at"`
// 	OtherChainLastBlockFetchedAt time.Time  `json:"other_chain_last_block_fetched_at"`

// 	BnbStatus        interface{} `json:"bnb_status"`
// 	OtherChainStatus interface{} `json:"other_chain_status"`
// }

// EthStatus ...
type EthStatus struct {
	Allowance    string `json:"allowance"`
	Erc20Balance string `json:"erc20_balance"`
	EthBalance   string `json:"eth_balance"`
}
