package models

// AccountCount 账户数量
type AccountCount struct {
	ID           uint64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	AccountCount uint32 `gorm:"column:account_count"`
	Date         int64  `gorm:"column:date;UNIQUE"`
}

// TableName 表名
func (accountCount AccountCount) TableName() string {
	return "account_count_tbl"
}
