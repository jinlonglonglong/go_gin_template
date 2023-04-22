package dao

import (
	"scan/pkg/data"
	"scan/pkg/models"
	"scan/pkg/util"
)

func GetAccountTxs(addr string, page uint32, limit uint32) (accountTxs []models.AccountTx) {
	db := data.MustGetDB("scan")
	err := db.Limit(limit).Offset((page-1)*limit).Order("id desc").Where("addr = ?", addr).Find(&accountTxs).Error
	util.CheckError(err)

	return accountTxs
}

func SetAccountTx(accountTx models.AccountTx) {
	db := data.MustGetDB("scan")
	err := db.Create(&accountTx).Error
	util.CheckError(err)
}

func GetAccountTxCount(addr string) (count uint32) {
	db := data.MustGetDB("scan")
	err := db.Table("account_tx_tbl").Where("addr = ?", addr).Count(&count).Error
	util.CheckError(err)

	return count
}
