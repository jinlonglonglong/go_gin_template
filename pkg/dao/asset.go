package dao

import (
	"scan/pkg/data"
	"scan/pkg/models"
	"scan/pkg/util"
)

func SetAsset(asset models.Asset) {
	db := data.MustGetDB("scan")
	err := db.Save(&asset).Error
	util.CheckError(err)
}

func GetAssetCount() (count uint32) {
	db := data.MustGetDB("scan")
	err := db.Table("asset_tbl").Count(&count).Error
	util.CheckError(err)

	return count
}

func AssetExist(assetSymbol string) (asset models.Asset, ret bool) {
	ret = false

	db := data.MustGetDB("scan")
	db.First(&asset, "asset_symbol=?", assetSymbol)

	if (models.Asset{}) == asset {
		return
	}

	ret = true

	return
}
