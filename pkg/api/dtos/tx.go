package dtos

type GetTxByTxIDDTO struct {
	TxID string `json:"txid"`
}

type GetTxsDTO struct {
	Page   uint32 `json:"page"`
	Count  uint32 `json:"count"`
	TxType string `json:"tx_type"`
}

type GetTxsByBlockHashDTO struct {
	BlockHash string `json:"hash"`
	Page      uint32 `json:"page"`
	Count     uint32 `json:"count"`
}

type GetTxsByAccountAddress struct {
	Address string `json:"address"`
	Page    uint32 `json:"page"`
	Count   uint32 `json:"count"`
}

type ReceiptsDTO []struct {
	FromAddr    string `json:"from_addr"`
	ToAddr      string `json:"to_addr"`
	CoinSymbol  string `json:"coin_symbol"`
	CoinAmount  uint64 `json:"coin_amount"`
	ReceiptCode uint32 `json:"receipt_code"`
	Memo        string `json:"memo"`
}

type CandidateVotesDTO []struct {
	VoteType     string `json:"vote_type"`
	VotedBcoins  uint64 `json:"voted_bcoins"`
	CandidateUID struct {
		IDType string `json:"id_type"`
		ID     string `json:"id"`
	} `json:"candidate_uid"`
	CandidateAddress string `json:"candidate_address"`
}

type TransfersDTO []struct {
	ToUID      string `json:"to_uid"`
	ToAddr     string `json:"to_addr"`
	CoinSymbol string `json:"coin_symbol"`
	CoinAmount uint64 `json:"coin_amount"`
}

type TxDTO struct {
	TxID            string      `json:"txid"`
	TxType          string      `json:"tx_type"`
	Version         uint8       `json:"version"`
	ValidHeight     uint32      `json:"valid_height"`
	Confirmations   uint32      `json:"confirmations"`
	ConfirmedHeight uint32      `json:"confirmed_height"`
	ConfirmedTime   uint32      `json:"confirmed_time"`
	BlockHash       string      `json:"block_hash"`
	Receipts        ReceiptsDTO `json:"receipts"`
	RawTx           string      `json:"rawtx"`

	TxUID          string            `json:"tx_uid,omitempty"`
	FromAddr       string            `json:"from_addr,omitempty"`
	ToUID          string            `json:"to_uid,omitempty"`
	ToAddr         string            `json:"to_addr,omitempty"`
	FeeSymbol      string            `json:"fee_symbol"`
	Fees           uint64            `json:"fees"`
	PubKey         string            `json:"pubkey,omitempty"`
	MinerPubKey    string            `json:"miner_pubkey,omitempty"`
	Signature      string            `json:"signature,omitempty"`
	CandidateVotes CandidateVotesDTO `json:"candidate_votes,omitempty"`
	Transfers      TransfersDTO      `json:"transfers,omitempty"`
	Memo           string            `json:"memo,omitempty"`
	RequiredSigs   uint8             `json:"required_sigs,omitempty"`
	Signatures     []string          `json:"signatures,omitempty"`
	StakeType      string            `json:"stake_type,omitempty"`
	CoinSymbol     string            `json:"coin_symbol"`
	CoinAmount     uint64            `json:"coin_amount"`
	VMType         uint8             `json:"vm_type,omitempty"`
	Upgradable     bool              `json:"upgradable,omitempty"`
	Code           string            `json:"code,omitempty"`
	Abi            string            `json:"abi,omitempty"`
	Arguments      string            `json:"arguments,omitempty"`
	AssetName      string            `json:"asset_name,omitempty"`
	AssetSymbol    string            `json:"asset_symbol,omitempty"`
	OwnerUID       string            `json:"owner_uid,omitempty"`
	OwnerAddr      string            `json:"owner_addr,omitempty"`
	TotalSupply    uint64            `json:"total_supply,omitempty"`
	Mintable       bool              `json:"mintable,omitempty"`
	UpdateType     string            `json:"update_type,omitempty"`
	UpdateValue    string            `json:"update_value,omitempty"`
	Domain         string            `json:"domain,omitempty"`
	Key            string            `json:"key,omitempty"`
	Value          string            `json:"value,omitempty"`
}

type TxListDTO struct {
	Data  []TxDTO `json:"data,omitempty"`
	Count uint32  `json:"count,omitempty"`
}
