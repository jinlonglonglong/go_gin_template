package dao

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/golang/glog"
	"scan/pkg/api/dtos"
	"scan/pkg/data"
	"scan/pkg/models"
	"scan/pkg/util"
	"strconv"
	"time"
)

// GetTxByTxID 根据交易哈希查询交易信息
func GetTxByTxID(txid string) (tx models.Tx) {
	db := data.MustGetDB("scan")
	db.First(&tx, "txid=?", txid)

	return tx
}

func GetTxsByTxIDs(txids []string) (txs []models.Tx) {
	db := data.MustGetDB("scan")
	err := db.Order("id desc").Where("txid in (?)", txids).Find(&txs).Error
	util.CheckError(err)

	return txs
}

// GetLatestTxs 查询最新的交易信息
func GetLatestTxs() (txs []models.Tx) {
	db := data.MustGetDB("scan")
	err := db.Limit(util.DefaultLatestTxs).Order("id desc").Find(&txs).Error
	util.CheckError(err)

	return txs
}

func GetTxs(page uint32, limit uint32, txType string) (txs []models.Tx) {
	db := data.MustGetDB("scan")
	var err error
	if len(txType) == 0 {
		err = db.Limit(limit).Offset((page - 1) * limit).Order("id desc").Find(&txs).Error
	} else {
		err = db.Limit(limit).Offset((page-1)*limit).Order("id desc").Where("tx_type = ?", txType).Find(&txs).Error
	}

	util.CheckError(err)

	return txs
}

func GetTxsByBlockHash(page uint32, limit uint32, hash string) (txs []models.Tx) {
	db := data.MustGetDB("scan")
	err := db.Limit(limit).Offset((page-1)*limit).Order("id desc").Where("block_hash = ?", hash).Find(&txs).Error
	util.CheckError(err)

	return txs
}

func GetTxCountByBlockHash(hash string) (count uint32) {
	db := data.MustGetDB("scan")
	err := db.Table("tx_tbl").Where("block_hash = ?", hash).Count(&count).Error
	util.CheckError(err)

	return count
}

// GetTxCountByTxType 从数据库查询总交易数
func GetTxCountByTxType(txType string) (count uint32) {
	db := data.MustGetDB("scan")
	var err error
	if len(txType) == 0 {
		err = db.Table("tx_tbl").Count(&count).Error
	} else {
		err = db.Table("tx_tbl").Where("tx_type = ?", txType).Count(&count).Error
	}

	util.CheckError(err)

	return count
}

// WriteTx 写入交易信息到数据库
func WriteTx(tx models.Tx) {
	db := data.MustGetDB("scan")
	err := db.Create(&tx).Error
	util.CheckError(err)
}

func SetTxCountCache(count uint64) bool {
	redisClient := data.MustGetRedis("scan")
	err := redisClient.Set(util.TxCountKey, count, util.DurationTxCount*time.Second).Err()
	if err == redis.Nil {
		glog.Errorf("failed to set: %s to redis", util.TxCountKey)
		return false
	}

	return true
}

func GetTxCountCache() (count uint64, ret bool) {
	ret = false
	redisClient := data.MustGetRedis("scan")
	value, err := redisClient.Get(util.TxCountKey).Result()

	if err == redis.Nil {
		return
	} else if err != nil {
		glog.Errorf("failed to get from redis: %s", err.Error())
		return
	}

	count, _ = strconv.ParseUint(value, 10, 32)
	ret = true

	return
}

func SetTxCache(txid string, tx dtos.TxDTO) bool {
	key := util.BuildTxIDKey(txid)
	redisClient := data.MustGetRedis("scan")
	txJson, _ := json.Marshal(tx)
	err := redisClient.Set(key, txJson, util.DurationTx*time.Second).Err()
	if err == redis.Nil {
		glog.Errorf("failed to set: %s to redis", key)
		return false
	}

	return true
}

func GetTxCache(txid string) (tx dtos.TxDTO, ret bool) {
	key := util.BuildTxIDKey(txid)
	ret = false
	redisClient := data.MustGetRedis("scan")
	value, err := redisClient.Get(key).Result()

	if err == redis.Nil {
		return
	} else if err != nil {
		glog.Errorf("failed to get from redis: %s", err.Error())
		return
	}

	json.Unmarshal([]byte(value), &tx)
	ret = true

	return
}

func SetLatestTxsCache(txs []dtos.TxDTO) bool {
	redisClient := data.MustGetRedis("scan")
	txJson, _ := json.Marshal(txs)
	err := redisClient.Set(util.LatestTxsKey, txJson, util.DurationLatestTxs*time.Second).Err()
	if err == redis.Nil {
		glog.Errorf("failed to set: %s to redis", util.LatestTxsKey)
		return false
	}

	return true
}

func GetLatestTxsCache() (txs []dtos.TxDTO, ret bool) {
	ret = false
	redisClient := data.MustGetRedis("scan")
	value, err := redisClient.Get(util.LatestTxsKey).Result()

	if err == redis.Nil {
		return
	} else if err != nil {
		glog.Errorf("failed to get from redis: %s", err.Error())
		return
	}

	json.Unmarshal([]byte(value), &txs)
	ret = true

	return
}
