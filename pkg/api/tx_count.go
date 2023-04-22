package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scan/pkg/api/dtos"
	"scan/pkg/dao"
	"scan/pkg/util"
)

// TxCountController 交易数量
type TxCountController struct {
	Router *gin.RouterGroup
}

// Setup 初始化路由
func (controller TxCountController) Setup() {
	handle := controller.Router.Group("tx_count")
	{
		handle.POST("latest", controller.getLatestTxCounts)
	}
}

func (controller *TxCountController) getLatestTxCounts(c *gin.Context) {
	var req dtos.GetLatestTxCountsDTO
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParams, "error_msg": "invalid params", "data": nil})
		return
	}

	if req.Count == 0 || req.Count > util.DefaultMaxLatestTxCounts {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParamRange, "error_msg": "invalid param range", "data": nil})
		return
	}

	limit := req.Count

	if txCountDTOs, ret := dao.GetTxCountsCache(limit); ret == true {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": txCountDTOs})
		return
	}

	txCounts := dao.GetLatestTxCounts(limit)
	if len(txCounts) == 0 {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeNotExist, "error_msg": "not exist", "data": nil})
		return
	}

	txCountDTOs := make([]dtos.TxCountDTO, 0, len(txCounts))
	for _, txCount := range txCounts {
		txCountDTO := dtos.TxCountDTO{
			TxCount: txCount.TxCount,
			Date:    util.TransferTimestampToDate(txCount.Date),
		}

		txCountDTOs = append(txCountDTOs, txCountDTO)
	}

	dao.SetTxCountsCache(limit, txCountDTOs)
	c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": txCountDTOs})
}
