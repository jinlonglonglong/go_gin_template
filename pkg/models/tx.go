package models

/*
+------------------+---------------------+------+-----+---------+----------------+
| Field            | Type                | Null | Key | Default | Extra          |
+------------------+---------------------+------+-----+---------+----------------+
| id               | int(10) unsigned    | NO   | PRI | NULL    | auto_increment |
| txid             | varchar(64)         | NO   | UNI | NULL    |                |
| tx_type          | varchar(30)         | NO   |     | NULL    |                |
| version          | tinyint(3) unsigned | NO   |     | NULL    |                |
| valid_height     | int(5) unsigned     | NO   |     | NULL    |                |
| confirmed_height | int(10) unsigned    | NO   |     | NULL    |                |
| confirmed_time   | int(10) unsigned    | NO   |     | NULL    |                |
| block_hash       | varchar(64)         | NO   | MUL | NULL    |                |
| receipts         | mediumtext          | YES  |     | NULL    |                |
| rawtx            | mediumtext          | YES  |     | NULL    |                |
| tx_uid           | varbinary(66)       | YES  |     | NULL    |                |
| from_addr        | varbinary(34)       | YES  | MUL | NULL    |                |
| to_uid           | varbinary(66)       | YES  |     | NULL    |                |
| to_addr          | varbinary(34)       | YES  | MUL | NULL    |                |
| fee_symbol       | varchar(30)         | YES  |     | NULL    |                |
| fees             | bigint(20) unsigned | YES  |     | NULL    |                |
| pubkey           | varbinary(66)       | YES  |     | NULL    |                |
| miner_pubkey     | varbinary(66)       | YES  |     | NULL    |                |
| signature        | varchar(200)        | YES  |     | NULL    |                |
| candidate_votes  | mediumtext          | YES  |     | NULL    |                |
| transfers        | mediumtext          | YES  |     | NULL    |                |
| memo             | varchar(200)        | YES  |     | NULL    |                |
| required_sigs    | tinyint(3) unsigned | YES  |     | NULL    |                |
| signatures       | mediumtext          | YES  |     | NULL    |                |
| stake_type       | varchar(30)         | YES  |     | NULL    |                |
| coin_symbol      | varchar(30)         | YES  |     | NULL    |                |
| coin_amount      | bigint(20) unsigned | YES  |     | NULL    |                |
| vm_type          | tinyint(3) unsigned | YES  |     | NULL    |                |
| upgradable       | tinyint(1) unsigned | YES  |     | NULL    |                |
| code             | mediumtext          | YES  |     | NULL    |                |
| abi              | mediumtext          | YES  |     | NULL    |                |
| arguments        | mediumtext          | YES  |     | NULL    |                |
| asset_name       | varchar(60)         | YES  |     | NULL    |                |
| asset_symbol     | varchar(30)         | YES  |     | NULL    |                |
| owner_uid        | varbinary(66)       | YES  |     | NULL    |                |
| owner_addr       | varbinary(34)       | YES  |     | NULL    |                |
| total_supply     | bigint(20) unsigned | YES  |     | NULL    |                |
| mintable         | tinyint(1) unsigned | YES  |     | NULL    |                |
| update_type      | varchar(30)         | YES  |     | NULL    |                |
| update_value     | varchar(30)         | YES  |     | NULL    |                |
| domain           | varchar(100)        | YES  |     | NULL    |                |
| key              | varchar(100)        | YES  |     | NULL    |                |
| value            | text                | YES  |     | NULL    |                |
+------------------+---------------------+------+-----+---------+----------------+
*/

// Tx 交易信息
type Tx struct {
	ID              uint64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	TxID            string `gorm:"column:txid"`
	TxType          string `gorm:"column:tx_type"`
	Version         uint8  `gorm:"column:version"`
	ValidHeight     uint32 `gorm:"column:valid_height"`
	ConfirmedHeight uint32 `gorm:"column:confirmed_height"`
	ConfirmedTime   uint32 `gorm:"column:confirmed_time"`
	BlockHash       string `gorm:"column:block_hash;UNIQUE"`
	Receipts        string `gorm:"column:receipts"`
	RawTx           string `gorm:"column:rawtx"`

	TxUID          string `gorm:"column:tx_uid"`
	FromAddr       string `gorm:"column:from_addr"`
	ToUID          string `gorm:"column:to_uid"`
	ToAddr         string `gorm:"column:to_addr"`
	FeeSymbol      string `gorm:"column:fee_symbol"`
	Fees           uint64 `gorm:"column:fees"`
	Pubkey         string `gorm:"column:pubkey"`
	MinerPubkey    string `gorm:"column:miner_pubkey"`
	Signature      string `gorm:"column:signature"`
	CandidateVotes string `gorm:"column:candidate_votes"`
	Transfers      string `gorm:"column:transfers"`
	Memo           string `gorm:"column:memo"`
	RequiredSigs   uint8  `gorm:"column:required_sigs"`
	Signatures     string `gorm:"column:signatures"`
	StakeType      string `gorm:"column:stake_type"`
	CoinSymbol     string `gorm:"column:coin_symbol"`
	CoinAmount     uint64 `gorm:"column:coin_amount"`
	VMType         uint8  `gorm:"column:vm_type"`
	Upgradable     bool   `gorm:"column:upgradable"`
	Code           string `gorm:"column:code"`
	Abi            string `gorm:"column:abi"`
	Arguments      string `gorm:"column:arguments"`
	AssetName      string `gorm:"column:asset_name"`
	AssetSymbol    string `gorm:"column:asset_symbol"`
	OwnerUID       string `gorm:"column:owner_uid"`
	OwnerAddr      string `gorm:"column:owner_addr"`
	TotalSupply    uint64 `gorm:"column:total_supply"`
	Mintable       bool   `gorm:"column:mintable"`
	UpdateType     string `gorm:"column:update_type"`
	UpdateValue    string `gorm:"column:update_value"`
	Domain         string `gorm:"column:domain"`
	Key            string `gorm:"column:key"`
	Value          string `gorm:"column:value"`
}

// TableName 表名
func (tx Tx) TableName() string {
	return "tx_tbl"
}
