package dtos

type InfoDTO struct {
	TotalAccounts uint32 `json:"total_accounts"`
	TotalTokens   uint32 `json:"total_tokens"`
	TotalBlocks   uint32 `json:"total_blocks"`
	TotalTxs      uint32 `json:"total_txs"`
}

type InfoDTOInternal struct {
	Version         string `json:"version"`
	ProtocolVersion uint32 `json:"protocol_version"`
	NetType         string `json:"net_type"`
	Proxy           string `json:"proxy"`
	ExtIp           string `json:"ext_ip"`
	ConfDir         string `json:"conf_dir"`
	DataDir         string `json:"data_dir"`
	BlockInterval   uint32 `json:"block_interval"`
	GenBlock        uint8  `json:"genblock"`
	CoinName        string `json:"coin_name"`
	//TotalCoins       uint64  `json:"total_coins"`
	//TotalAccounts    uint32  `json:"total_accounts"`
	//TotalContracts   uint32  `json:"total_contracts"`
	WalletBalance    float32 `json:"wallet_balance"`
	RelayFeePerKb    float32 `json:"relay_fee_perkb"`
	TipBlockFuelRate uint32  `json:"tipblock_fuel_rate"`
	TipBlockFuel     uint32  `json:"tipblock_fuel"`
	TipBlockTime     uint32  `json:"tipblock_time"`
	TipBlockHash     string  `json:"tipblock_hash"`
	TipBlockHeight   uint32  `json:"tipblock_height"`
	SynBlockHeight   uint32  `json:"synblock_height"`
	Connections      uint32  `json:"connections"`
	Errors           string  `json:"errors"`
}
