package models

/*
+----------+---------------------+------+-----+---------+----------------+
| Field    | Type                | Null | Key | Default | Extra          |
+----------+---------------------+------+-----+---------+----------------+
| id       | int(10) unsigned    | NO   | PRI | NULL    | auto_increment |
| tx_count | int(10) unsigned    | YES  |     | NULL    |                |
| date     | bigint(20)          | YES  | UNI | NULL    |                |
+----------+---------------------+------+-----+---------+----------------+
*/

// TxCount 交易数量
type TxCount struct {
	ID      uint64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	TxCount uint32 `gorm:"column:tx_count"`
	Date    int64  `gorm:"column:date;UNIQUE"`
}

// TableName 表名
func (txCount TxCount) TableName() string {
	return "tx_count_tbl"
}
