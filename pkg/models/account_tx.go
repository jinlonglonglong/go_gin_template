package models

/*
+-------+------------------+------+-----+---------+----------------+
| Field | Type             | Null | Key | Default | Extra          |
+-------+------------------+------+-----+---------+----------------+
| id    | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| addr  | varbinary(34)    | YES  | MUL | NULL    |                |
| txid  | varchar(64)      | YES  |     | NULL    |                |
+-------+------------------+------+-----+---------+----------------+
*/

// AccountTx 账户相关交易信息
type AccountTx struct {
	ID   uint64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Addr string `gorm:"column:addr"`
	TxID string `gorm:"column:txid"`
}

// TableName 表名
func (accountTx AccountTx) TableName() string {
	return "account_tx_tbl"
}
