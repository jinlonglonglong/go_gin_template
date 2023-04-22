package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scan/pkg/api/dtos"
	"scan/pkg/util"
)

// MarketController 市场数据
type MarketController struct {
	Router *gin.RouterGroup
}

// Setup 初始化路由
func (controller MarketController) Setup() {
	handle := controller.Router.Group("market")
	{
		handle.POST("info", controller.getInfo)
	}
}

func (controller *MarketController) getInfo(c *gin.Context) {
	// TODO: according to actual market information
	marketDTO := dtos.MarketDTO{
		CNYPrice:        10,
		CNYMarketValue:  10000000000,
		CNYTurnover:     10000000,
		USDPrice:        66,
		USDYMarketValue: 66000000000,
		USDTurnover:     66000000,
		Changed:         2.58,
	}

	c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": marketDTO})
}
