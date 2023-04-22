package dao

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/golang/glog"
	"scan/pkg/api/dtos"
	"scan/pkg/data"
	"scan/pkg/models"
	"scan/pkg/util"
	"time"
)

// GetLatestTxCount 从数据库获取最后一条交易数量的记录
func GetLatestTxCount() (txCount models.TxCount) {
	db := data.MustGetDB("scan")
	db.Last(&txCount)

	return txCount
}

// GetLatestTxCounts 查询最新的交易数量信息
func GetLatestTxCounts(limit uint32) (txCounts []models.TxCount) {
	db := data.MustGetDB("scan")
	err := db.Limit(limit).Order("id desc").Find(&txCounts).Error
	util.CheckError(err)

	return txCounts
}

func SetTxCount(timestamp int64) {
	sql := fmt.Sprintf("insert into tx_count_tbl(tx_count, date) "+
		"select COUNT(id), '%d' from tx_tbl t "+
		"where t.confirmed_time >= %d and t.confirmed_time < %d",
		timestamp,
		timestamp,
		timestamp+util.SecondsOfDay)
	db := data.MustGetDB("scan")
	err := db.Exec(sql).Error
	util.CheckError(err)
}

func SetTxCountsCache(limit uint32, txCounts []dtos.TxCountDTO) bool {
	key := util.BuildTxCountKey(limit)
	redisClient := data.MustGetRedis("scan")
	txCountsJson, _ := json.Marshal(txCounts)
	err := redisClient.Set(key, txCountsJson, util.DurationLatestTxCounts*time.Second).Err()
	if err == redis.Nil {
		glog.Errorf("failed to set: %s to redis", key)
		return false
	}

	return true
}

func GetTxCountsCache(limit uint32) (txCounts []dtos.TxCountDTO, ret bool) {
	key := util.BuildTxCountKey(limit)
	ret = false
	redisClient := data.MustGetRedis("scan")
	value, err := redisClient.Get(key).Result()

	if err == redis.Nil {
		return
	} else if err != nil {
		glog.Errorf("failed to get from redis: %s", err.Error())
		return
	}

	json.Unmarshal([]byte(value), &txCounts)
	ret = true

	return
}
