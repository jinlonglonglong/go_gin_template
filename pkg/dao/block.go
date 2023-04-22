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

// GetBlockByHeight 根据区块高度查询区块信息
func GetBlockByHeight(height uint32) (block models.Block) {
	db := data.MustGetDB("scan")
	db.First(&block, "height=?", height)

	return block
}

// GetBlockByHash 根据区块高度查询区块信息
func GetBlockByHash(hash string) (block models.Block) {
	db := data.MustGetDB("scan")
	db.First(&block, "block_hash=?", hash)

	return block
}

// GetLatestBlocks 查询最新的区块信息
func GetLatestBlocks() (blocks []models.Block) {
	db := data.MustGetDB("scan")
	err := db.Limit(util.DefaultLatestBlocks).Order("id desc").Find(&blocks).Error
	util.CheckError(err)

	return blocks
}

func GetBlocks(page uint32, limit uint32) (blocks []models.Block) {
	db := data.MustGetDB("scan")
	err := db.Limit(limit).Offset((page - 1) * limit).Order("id desc").Find(&blocks).Error
	util.CheckError(err)

	return blocks
}

// GetLatestBlock 从数据库获取最近一次同步的区块
func GetLatestBlock() (block models.Block) {
	db := data.MustGetDB("scan")
	// Attention: 初始化数据库时区块数据为空
	//err := db.Last(&block).Error // 按照主键的最后一条记录
	//util.CheckError(err)
	db.Last(&block)

	return block
}

// GetBlockCount 从数据库查询总区块数
func GetBlockCount() (count uint32) {
	db := data.MustGetDB("scan")
	err := db.Table("block_tbl").Count(&count).Error
	util.CheckError(err)

	return count
}

// SetBlock 写入区块信息到数据库
func SetBlock(block models.Block) {
	db := data.MustGetDB("scan")
	err := db.Create(&block).Error
	util.CheckError(err)
}

func UpdatePreviousBlock(block models.Block) {
	db := data.MustGetDB("scan")
	err := db.Table("block_tbl").Where("block_hash = ?", block.PreviousBlockHash).Update("next_block_hash", block.BlockHash).Error
	util.CheckError(err)
}

func SetBlockCountCache(count uint64) bool {
	redisClient := data.MustGetRedis("scan")
	err := redisClient.Set(util.BlockCountKey, count, util.DurationBlockCount*time.Second).Err()
	if err == redis.Nil {
		glog.Errorf("failed to set: %s to redis", util.BlockCountKey)
		return false
	}

	return true
}

func GetBlockCountCache() (count uint64, ret bool) {
	ret = false
	redisClient := data.MustGetRedis("scan")
	value, err := redisClient.Get(util.BlockCountKey).Result()

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

func SetBlockCacheByHeight(height uint32, block dtos.BlockDTO) bool {
	key := util.BuildBlockHeightKey(height)
	redisClient := data.MustGetRedis("scan")
	blockJson, _ := json.Marshal(block)
	err := redisClient.Set(key, blockJson, util.DurationBlock*time.Second).Err()
	if err == redis.Nil {
		glog.Errorf("failed to set: %s to redis", key)
		return false
	}

	return true
}

func GetBlockCacheByHeight(height uint32) (block dtos.BlockDTO, ret bool) {
	key := util.BuildBlockHeightKey(height)
	ret = false
	redisClient := data.MustGetRedis("scan")
	value, err := redisClient.Get(key).Result()

	if err == redis.Nil {
		return
	} else if err != nil {
		glog.Errorf("failed to get from redis: %s", err.Error())
		return
	}

	json.Unmarshal([]byte(value), &block)
	ret = true

	return
}

func SetBlockCacheByHash(hash string, block dtos.BlockDTO) bool {
	key := util.BuildBlockHashKey(hash)
	redisClient := data.MustGetRedis("scan")
	blockJson, _ := json.Marshal(block)
	err := redisClient.Set(key, blockJson, util.DurationBlock*time.Second).Err()
	if err == redis.Nil {
		glog.Errorf("failed to set: %s to redis", key)
		return false
	}

	return true
}

func GetBlockCacheByHash(hash string) (block dtos.BlockDTO, ret bool) {
	key := util.BuildBlockHashKey(hash)
	ret = false
	redisClient := data.MustGetRedis("scan")
	value, err := redisClient.Get(key).Result()

	if err == redis.Nil {
		return
	} else if err != nil {
		glog.Errorf("failed to get from redis: %s", err.Error())
		return
	}

	json.Unmarshal([]byte(value), &block)
	ret = true

	return
}

func SetLatestBlocksCache(blocks []dtos.BlockDTO) bool {
	redisClient := data.MustGetRedis("scan")
	blocksJson, _ := json.Marshal(blocks)
	err := redisClient.Set(util.LatestBlocksKey, blocksJson, util.DurationLatestBlocks*time.Second).Err()
	if err == redis.Nil {
		glog.Errorf("failed to set: %s to redis", util.LatestBlocksKey)
		return false
	}

	return true
}

func GetLatestBlocksCache() (blocks []dtos.BlockDTO, ret bool) {
	ret = false
	redisClient := data.MustGetRedis("scan")
	value, err := redisClient.Get(util.LatestBlocksKey).Result()

	if err == redis.Nil {
		return
	} else if err != nil {
		glog.Errorf("failed to get from redis: %s", err.Error())
		return
	}

	json.Unmarshal([]byte(value), &blocks)
	ret = true

	return
}
