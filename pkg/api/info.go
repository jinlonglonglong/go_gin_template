package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scan/pkg/dao"
	"scan/pkg/util"
)

// TxController 交易信息
type InfoController struct {
	Router *gin.RouterGroup
}

// Setup 初始化路由
func (controller InfoController) Setup() {
	handle := controller.Router.Group("info")
	{
		handle.POST("", controller.getInfo)
	}
}

func (controller *InfoController) getInfo(c *gin.Context) {
	if infoDTO, ret := dao.GetInfoCache(); ret == true {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": infoDTO})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInternalError, "error_msg": "internal error, try again", "data": nil})
}
