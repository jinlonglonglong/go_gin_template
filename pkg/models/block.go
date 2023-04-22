package models

/*
+---------------------+---------------------+------+-----+---------+----------------+
| Field               | Type                | Null | Key | Default | Extra          |
+---------------------+---------------------+------+-----+---------+----------------+
| id                  | int(10) unsigned    | NO   | PRI | NULL    | auto_increment |
| height              | int(10) unsigned    | NO   | UNI | NULL    |                |
| block_hash          | varchar(64)         | NO   | UNI | NULL    |                |
| miner_uid           | varchar(66)         | NO   |     | NULL    |                |
| miner_address       | varchar(34)         | NO   |     | NULL    |                |
| size       	      | int(10) unsigned    | NO   |     | NULL    |                |
| version             | tinyint(3) unsigned | NO   |     | NULL    |                |
| merkle_root         | varchar(64)         | NO   |     | NULL    |                |
| time                | int(10) unsigned    | NO   |     | NULL    |                |
| previous_block_hash | varchar(64)         | NO   |     | NULL    |                |
| next_block_hash     | varchar(64)         | YES  |     | NULL    |                |
| tx_count            | int(10) unsigned    | YES  |     | 0       |                |
| tx                  | mediumtext          | YES  |     | NULL    |                |
+---------------------+---------------------+------+-----+---------+----------------+
*/

// Block 区块信息
type Block struct {
	ID                uint64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Height            uint32 `gorm:"column:height;UNIQUE"`
	BlockHash         string `gorm:"column:block_hash;UNIQUE"`
	MinerUID          string `gorm:"column:miner_uid"`
	MinerAddress      string `gorm:"column:miner_address"`
	Size              uint32 `gorm:"column:size"`
	Version           uint8  `gorm:"column:version"`
	MerkleRoot        string `gorm:"column:merkle_root"`
	Time              uint32 `gorm:"column:time"`
	PreviousBlockHash string `gorm:"column:previous_block_hash"`
	NextBlockHash     string `gorm:"column:next_block_hash"`
	TxCount           uint32 `gorm:"column:tx_count"`
	Tx                string `gorm:"tx"`
}

// TableName 表名
func (block Block) TableName() string {
	return "block_tbl"
}
