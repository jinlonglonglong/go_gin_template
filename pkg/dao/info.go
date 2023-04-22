package dao

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/golang/glog"
	"scan/pkg/api/dtos"
	"scan/pkg/data"
	"scan/pkg/util"
	"time"
)

func SetInfoCache(info dtos.InfoDTO) bool {
	redisClient := data.MustGetRedis("scan")
	infoJson, _ := json.Marshal(info)
	err := redisClient.Set(util.InfoKey, infoJson, util.DurationInfo*time.Second).Err()
	if err == redis.Nil {
		glog.Errorf("failed to set: %s to redis", util.InfoKey)
		return false
	}

	return true
}

func GetInfoCache() (info dtos.InfoDTO, ret bool) {
	ret = false
	redisClient := data.MustGetRedis("scan")
	infoJson, err := redisClient.Get(util.InfoKey).Result()

	if err == redis.Nil {
		return
	} else if err != nil {
		glog.Errorf("failed to get from redis: %s", err.Error())
		return
	}

	json.Unmarshal([]byte(infoJson), &info)
	ret = true

	return
}
