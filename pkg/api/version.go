package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scan/pkg/util"
)

// VersionController 版本信息
type VersionController struct {
	Router *gin.RouterGroup
}

// Setup 初始化路由
func (controller VersionController) Setup() {
	handle := controller.Router.Group("version")
	{
		handle.GET("", controller.getVersion)
		handle.POST("", controller.getVersion)
	}
}

func (controller *VersionController) getVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Version": util.Version, "Compile": util.Compile})
}
