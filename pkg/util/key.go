package util

import "fmt"

// complete, unique key
const (
	InfoKey = "node_info_key"

	LatestBlocksKey = "latest_blocks_key"
	BlockCountKey   = "block_count_key"

	LatestTxsKey = "latest_txs_key"
	TxCountKey   = "tx_count_key"

	TxTypesKey = "tx_types_key"
)

// unique, prefix key
const (
	AccountPrefix      = "account"
	AccountCountPrefix = "account_count"
	BlockHeightPrefix  = "block_height"
	BlockHashPrefix    = "block_hash"
	TxCountPrefix      = "tx_count"
	TxIDPrefix         = "tx_txid"
)

func BuildAccountKey(address string) string {
	return fmt.Sprintf("%s_%s", AccountPrefix, address)
}

func BuildBlockHeightKey(height uint32) string {
	return fmt.Sprintf("%s_%d", BlockHeightPrefix, height)
}

func BuildBlockHashKey(hash string) string {
	return fmt.Sprintf("%s_%s", BlockHashPrefix, hash)
}

func BuildTxCountKey(limit uint32) string {
	return fmt.Sprintf("%s_%d", TxCountPrefix, limit)
}

func BuildAccountCountKey(limit uint32) string {
	return fmt.Sprintf("%s_%d", AccountCountPrefix, limit)
}

func BuildTxIDKey(txid string) string {
	return fmt.Sprintf("%s_%s", TxIDPrefix, txid)
}
