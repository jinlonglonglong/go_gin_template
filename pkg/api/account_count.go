package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scan/pkg/api/dtos"
	"scan/pkg/dao"
	"scan/pkg/util"
)

type AccountCountController struct {
	Router *gin.RouterGroup
}

// Setup 初始化路由
func (controller AccountCountController) Setup() {
	handle := controller.Router.Group("account_count")
	{
		handle.POST("latest", controller.getLatestAccountCounts)
	}
}

func (controller *AccountCountController) getLatestAccountCounts(c *gin.Context) {
	var req dtos.GetLatestAccountCountsDTO
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParams, "error_msg": "invalid params", "data": nil})
		return
	}

	if req.Count == 0 || req.Count > util.DefaultMaxLatestAccountCounts {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParamRange, "error_msg": "invalid param range", "data": nil})
		return
	}

	limit := req.Count

	if accountCountDTOs, ret := dao.GetAccountCountsCache(limit); ret == true {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": accountCountDTOs})
		return
	}

	accountCounts := dao.GetLatestAccountCounts(limit)
	if len(accountCounts) == 0 {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeNotExist, "error_msg": "not exist", "data": nil})
		return
	}

	accountCountDTOs := make([]dtos.AccountCountDTO, 0, len(accountCounts))
	for _, accountCount := range accountCounts {
		accountCountDTO := dtos.AccountCountDTO{
			AccountCount: accountCount.AccountCount,
			Date:         util.TransferTimestampToDate(accountCount.Date),
		}

		accountCountDTOs = append(accountCountDTOs, accountCountDTO)
	}

	dao.SetAccountCountsCache(limit, accountCountDTOs)
	c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": accountCountDTOs})
}
