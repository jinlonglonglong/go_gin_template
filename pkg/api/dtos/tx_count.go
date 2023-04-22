package dtos

type GetLatestTxCountsDTO struct {
	Count uint32 `json:"count"`
}

type TxCountDTO struct {
	TxCount uint32 `json:"tx_count"`
	Date    string `json:"date"`
}
