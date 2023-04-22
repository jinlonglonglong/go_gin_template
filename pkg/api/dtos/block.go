package dtos

type GetBlockByHeightDTO struct {
	Height uint32 `json:"height"`
}

type GetBlockByHashDTO struct {
	Hash string `json:"hash"`
}

type GetBlocksDTO struct {
	Page  uint32 `json:"page"`
	Count uint32 `json:"count"`
}

type BlockDTO struct {
	BlockHash         string   `json:"block_hash"`
	MinerUID          string   `json:"miner_uid"`
	MinerAddress      string   `json:"miner_address"`
	Confirmations     uint32   `json:"confirmations"`
	Size              uint32   `json:"size"`
	Height            uint32   `json:"height"`
	Version           uint8    `json:"version"`
	MerkleRoot        string   `json:"merkle_root"`
	TxCount           uint32   `json:"tx_count"`
	Tx                []string `json:"tx"`
	Time              uint32   `json:"time"`
	PreviousBlockHash string   `json:"previous_block_hash",omitempty`
	NextBlockHash     string   `json:"next_block_hash,omitempty"`
}

type BlockListDTO struct {
	Data  []BlockDTO `json:"data,omitempty"`
	Count uint32     `json:"count,omitempty"`
}
