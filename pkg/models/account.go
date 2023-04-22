package models

/*
+---------+------------------+------+-----+---------+----------------+
| Field   | Type             | Null | Key | Default | Extra          |
+---------+------------------+------+-----+---------+----------------+
| id      | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| addr    | varbinary(34)    | YES  | UNI | NULL    |                |
| balance | bigint(20)       | YES  |     | NULL    |                |
| vote    | bigint(20)       | YES  |     | NULL    |                |
+---------+------------------+------+-----+---------+----------------+
*/

// Account 账户信息
type Account struct {
	ID         uint64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Address    string `gorm:"column:addr"`
	Balance    uint64 `gorm:"column:balance"`
	Vote       uint64 `gorm:"column:vote"`
	CreateTime uint32 `gorm:"column:create_time"`
}

// TableName 表名
func (account Account) TableName() string {
	return "account_tbl"
}
