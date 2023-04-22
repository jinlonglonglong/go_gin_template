package helpers

import (
	"scan/pkg/dao"
	"scan/pkg/models"
	"scan/pkg/util"
)

func GetLatestBlockHeight() uint32 {
	block := dao.GetLatestBlock()
	return block.Height
}

func GetLatestTxCountDate() int64 {
	txCount := dao.GetLatestTxCount()
	if (models.TxCount{}) == txCount {
		return util.GetTimestampOfAMonthAgo()
	}

	return txCount.Date
}

func GetLatestAccountCountDate() int64 {
	accountCount := dao.GetLatestAccountCount()
	if (models.AccountCount{}) == accountCount {
		return util.GetTimestampOfAMonthAgo()
	}

	return accountCount.Date
}
