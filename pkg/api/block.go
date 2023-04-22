package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"scan/pkg/api/dtos"
	"scan/pkg/dao"
	"scan/pkg/models"
	"scan/pkg/util"
)

// BlockController 区块信息
type BlockController struct {
	Router *gin.RouterGroup
}

// Setup 初始化路由
func (controller BlockController) Setup() {
	handle := controller.Router.Group("block")
	{
		handle.POST("count", controller.getBlockCount)
		handle.POST("hash", controller.getBlockByHash)
		handle.POST("height", controller.getBlockByHeight)
		handle.POST("latest", controller.getLatestBlocks)
		handle.POST("list", controller.getBlocks)
	}
}

func (controller *BlockController) getBlockCount(c *gin.Context) {
	var count uint64
	if count, ret := dao.GetBlockCountCache(); ret == true {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": count})
		return
	}

	count = uint64(dao.GetBlockCount())

	dao.SetBlockCountCache(count)
	c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": count})
}

func (controller *BlockController) getBlockByHeight(c *gin.Context) {
	var req dtos.GetBlockByHeightDTO
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParams, "error_msg": "invalid params", "data": nil})
		return
	}

	if blockDTO, ret := dao.GetBlockCacheByHeight(req.Height); ret == true {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": blockDTO})
		return
	}

	block := dao.GetBlockByHeight(req.Height)
	if (models.Block{}) == block {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeNotExist, "error_msg": "not exist", "data": nil})
		return
	}

	var tx []string
	json.Unmarshal([]byte(block.Tx), &tx)

	blockDTO := dtos.BlockDTO{
		BlockHash:    block.BlockHash,
		MinerUID:     block.MinerUID,
		MinerAddress: block.MinerAddress,
		// TODO:
		//Confirmations:     nil,
		Size:              block.Size,
		Height:            block.Height,
		Version:           block.Version,
		MerkleRoot:        block.MerkleRoot,
		TxCount:           block.TxCount,
		Tx:                tx,
		Time:              block.Time,
		PreviousBlockHash: block.PreviousBlockHash,
		NextBlockHash:     block.NextBlockHash,
	}

	dao.SetBlockCacheByHeight(req.Height, blockDTO)
	c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": blockDTO})
}

func (controller *BlockController) getBlockByHash(c *gin.Context) {
	var req dtos.GetBlockByHashDTO
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParams, "error_msg": "invalid params", "data": nil})
		return
	}

	if blockDTO, ret := dao.GetBlockCacheByHash(req.Hash); ret == true {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": blockDTO})
		return
	}

	block := dao.GetBlockByHash(req.Hash)
	if (models.Block{}) == block {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeNotExist, "error_msg": "not exist", "data": nil})
		return
	}

	var tx []string
	json.Unmarshal([]byte(block.Tx), &tx)

	blockDTO := dtos.BlockDTO{
		BlockHash:    block.BlockHash,
		MinerUID:     block.MinerUID,
		MinerAddress: block.MinerAddress,
		// TODO:
		//Confirmations:     nil,
		Size:              block.Size,
		Height:            block.Height,
		Version:           block.Version,
		MerkleRoot:        block.MerkleRoot,
		TxCount:           block.TxCount,
		Tx:                tx,
		Time:              block.Time,
		PreviousBlockHash: block.PreviousBlockHash,
		NextBlockHash:     block.NextBlockHash,
	}

	dao.SetBlockCacheByHash(req.Hash, blockDTO)
	c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": blockDTO})
}

func (controller *BlockController) getLatestBlocks(c *gin.Context) {
	if blockDTOs, ret := dao.GetLatestBlocksCache(); ret == true {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": blockDTOs})
		return
	}

	blocks := dao.GetLatestBlocks()
	if len(blocks) == 0 {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeNotExist, "error_msg": "not exist", "data": nil})
		return
	}

	blockDTOs := make([]dtos.BlockDTO, 0, util.DefaultLatestBlocks)
	for _, block := range blocks {
		var tx []string
		json.Unmarshal([]byte(block.Tx), &tx)

		blockDTO := dtos.BlockDTO{
			BlockHash:    block.BlockHash,
			MinerUID:     block.MinerUID,
			MinerAddress: block.MinerAddress,
			// TODO:
			//Confirmations:     nil,
			Size:              block.Size,
			Height:            block.Height,
			Version:           block.Version,
			MerkleRoot:        block.MerkleRoot,
			TxCount:           block.TxCount,
			Tx:                tx,
			Time:              block.Time,
			PreviousBlockHash: block.PreviousBlockHash,
			NextBlockHash:     block.NextBlockHash,
		}

		blockDTOs = append(blockDTOs, blockDTO)
	}

	dao.SetLatestBlocksCache(blockDTOs)
	c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": blockDTOs})
}

func (controller *BlockController) getBlocks(c *gin.Context) {
	var req dtos.GetBlocksDTO
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParams, "error_msg": "invalid params", "data": nil})
		return
	}

	page := req.Page
	count := req.Count
	if (count == 0 || count > util.DefaultBlockPageSize) || (page == 0 || page > util.DefaultMaxBlockPage) {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParamRange, "error_msg": "invalid param range", "data": nil})
		return
	}

	blocks := dao.GetBlocks(page, count)
	if len(blocks) == 0 {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeNotExist, "error_msg": "not exist", "data": nil})
		return
	}

	blockDTOs := make([]dtos.BlockDTO, 0, len(blocks))
	for _, block := range blocks {
		var tx []string
		json.Unmarshal([]byte(block.Tx), &tx)

		blockDTO := dtos.BlockDTO{
			BlockHash:    block.BlockHash,
			MinerUID:     block.MinerUID,
			MinerAddress: block.MinerAddress,
			// TODO:
			//Confirmations:     nil,
			Size:              block.Size,
			Height:            block.Height,
			Version:           block.Version,
			MerkleRoot:        block.MerkleRoot,
			TxCount:           block.TxCount,
			Tx:                tx,
			Time:              block.Time,
			PreviousBlockHash: block.PreviousBlockHash,
			NextBlockHash:     block.NextBlockHash,
		}

		blockDTOs = append(blockDTOs, blockDTO)
	}

	records := dao.GetBlockCount()
	blockListDTO := dtos.BlockListDTO{
		Data:  blockDTOs,
		Count: records,
	}

	c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": blockListDTO})
}
