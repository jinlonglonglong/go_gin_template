package dtos

type TxTypeDTO struct {
	Enum          uint32 `json:"enum"`
	TxType        string `json:"tx_type"`
	CNDescription string `json:"cn_description"`
	ENDescription string `json:"en_description"`
}
