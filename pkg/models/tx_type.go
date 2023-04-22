package models

/*
+----------------+---------------------+------+-----+---------+----------------+
| Field          | Type                | Null | Key | Default | Extra          |
+----------------+---------------------+------+-----+---------+----------------+
| id             | int(10) unsigned    | NO   | PRI | NULL    | auto_increment |
| enum           | tinyint(3) unsigned | YES  |     | NULL    |                |
| tx_type        | varchar(50)         | YES  |     | NULL    |                |
| cn_description | varchar(50)         | YES  |     | NULL    |                |
| en_description | varchar(50)         | YES  |     | NULL    |                |
+----------------+---------------------+------+-----+---------+----------------+
*/

// TxType 交易类型
type TxType struct {
	ID            uint64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Enum          uint32 `gorm:"column:enum"`
	TxType        string `gorm:"column:tx_type"`
	CNDescription string `gorm:"column:cn_description"`
	ENDescription string `gorm:"column:en_description"`
}

// TableName 表名
func (txType TxType) TableName() string {
	return "tx_type_tbl"
}
