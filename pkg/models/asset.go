package models

/*
+---------------+---------------------+------+-----+---------+----------------+
| Field         | Type                | Null | Key | Default | Extra          |
+---------------+---------------------+------+-----+---------+----------------+
| id            | int(10) unsigned    | NO   | PRI | NULL    | auto_increment |
| asset_symbol  | varchar(20)         | NO   | UNI | NULL    |                |
| asset_name    | varchar(100)        | YES  |     | NULL    |                |
| owner_addr    | varbinary(40)       | YES  |     | NULL    |                |
| total_supply | bigint(20)          | YES  |     | NULL    |                |
| mintable      | tinyint(1) unsigned | YES  |     | NULL    |                |
+---------------+---------------------+------+-----+---------+----------------+
*/

type Asset struct {
	ID          uint64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	AssetSymbol string `gorm:"column:asset_symbol"`
	AssetName   string `gorm:"column:asset_name"`
	OwnerAddr   string `gorm:"column:owner_addr"`
	TotalSupply uint64 `gorm:"column:total_supply"`
	Mintable    bool   `gorm:"column:mintable"`
}

// TableName 表名
func (asset Asset) TableName() string {
	return "asset_tbl"
}
