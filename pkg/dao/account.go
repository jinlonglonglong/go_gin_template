package dao

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/golang/glog"
	"scan/pkg/api/dtos"
	"scan/pkg/data"
	"scan/pkg/models"
	"scan/pkg/util"
	"time"
)

func SetAccountCache(address string, account dtos.AccountDTO) bool {
	key := util.BuildAccountKey(address)
	redisClient := data.MustGetRedis("scan")
	accountJson, _ := json.Marshal(account)
	err := redisClient.Set(key, accountJson, util.DurationAccount*time.Second).Err()
	if err == redis.Nil {
		glog.Errorf("failed to set: %s to redis", key)
		return false
	}

	return true
}

func GetAccountCache(address string) (account dtos.AccountDTO, ret bool) {
	key := util.BuildAccountKey(address)
	ret = false
	redisClient := data.MustGetRedis("scan")
	value, err := redisClient.Get(key).Result()

	if err == redis.Nil {
		return
	} else if err != nil {
		glog.Errorf("failed to get from redis: %s", err.Error())
		return
	}

	json.Unmarshal([]byte(value), &account)
	ret = true

	return
}

func GetAccounts(page uint32, limit uint32) (accounts []models.Account) {
	db := data.MustGetDB("scan")
	err := db.Limit(limit).Offset((page - 1) * limit).Order("balance desc").Find(&accounts).Error
	util.CheckError(err)

	return accounts
}

func AccountExist(addr string) (account models.Account, ret bool) {
	ret = false

	db := data.MustGetDB("scan")
	db.First(&account, "addr=?", addr)

	if (models.Account{}) == account {
		return
	}

	ret = true

	return
}

func SetAccount(account models.Account) {
	db := data.MustGetDB("scan")
	err := db.Save(&account).Error
	util.CheckError(err)
}

func GetAccountCount() (count uint32) {
	db := data.MustGetDB("scan")
	err := db.Table("account_tbl").Count(&count).Error
	util.CheckError(err)

	return count
}
