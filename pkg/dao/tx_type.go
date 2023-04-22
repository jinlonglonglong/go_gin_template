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

func GetTxTypes() (txTypes []models.TxType) {
	db := data.MustGetDB("scan")
	err := db.Order("id asc").Find(&txTypes).Error
	util.CheckError(err)

	return txTypes
}

func SetTxTypesCache(txTypes []dtos.TxTypeDTO) bool {
	redisClient := data.MustGetRedis("scan")
	txTypesJson, _ := json.Marshal(txTypes)
	err := redisClient.Set(util.TxTypesKey, txTypesJson, util.DurationTxTypes*time.Second).Err()
	if err == redis.Nil {
		glog.Errorf("failed to set: %s to redis", util.TxTypesKey)
		return false
	}

	return true
}

func GetTxTypesCache() (txTypes []dtos.TxTypeDTO, ret bool) {
	ret = false
	redisClient := data.MustGetRedis("scan")
	value, err := redisClient.Get(util.TxTypesKey).Result()

	if err == redis.Nil {
		return
	} else if err != nil {
		glog.Errorf("failed to get from redis: %s", err.Error())
		return
	}

	json.Unmarshal([]byte(value), &txTypes)
	ret = true

	return
}
