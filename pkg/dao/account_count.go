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

func GetLatestAccountCount() (accountCount models.AccountCount) {
	db := data.MustGetDB("scan")
	db.Last(&accountCount)

	return accountCount
}

func GetLatestAccountCounts(limit uint32) (accountCounts []models.AccountCount) {
	db := data.MustGetDB("scan")
	err := db.Limit(limit).Order("id desc").Find(&accountCounts).Error
	util.CheckError(err)

	return accountCounts
}

func SetAccountCount(timestamp int64) {
	sql := fmt.Sprintf("insert into account_count_tbl(account_count, date) "+
		"select COUNT(id), '%d' from account_tbl t "+
		"where t.create_time >= %d and t.create_time < %d",
		timestamp,
		timestamp,
		timestamp+util.SecondsOfDay)
	db := data.MustGetDB("scan")
	err := db.Exec(sql).Error
	util.CheckError(err)
}

func SetAccountCountsCache(limit uint32, accountCounts []dtos.AccountCountDTO) bool {
	key := util.BuildAccountCountKey(limit)
	redisClient := data.MustGetRedis("scan")
	jsonStr, _ := json.Marshal(accountCounts)
	err := redisClient.Set(key, jsonStr, util.DurationLatestAccountCounts*time.Second).Err()
	if err == redis.Nil {
		glog.Errorf("failed to set: %s to redis", key)
		return false
	}

	return true
}

func GetAccountCountsCache(limit uint32) (accountCounts []dtos.AccountCountDTO, ret bool) {
	key := util.BuildAccountCountKey(limit)
	ret = false
	redisClient := data.MustGetRedis("scan")
	value, err := redisClient.Get(key).Result()

	if err == redis.Nil {
		return
	} else if err != nil {
		glog.Errorf("failed to get from redis: %s", err.Error())
		return
	}

	json.Unmarshal([]byte(value), &accountCounts)
	ret = true

	return
}
