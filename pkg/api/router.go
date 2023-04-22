package api

import (
	"github.com/gin-gonic/gin"
)

// Setup sets up all controllers.
func Setup(router *gin.Engine) {

	api := router.Group("api")

	account := AccountController{Router: api}
	account.Setup()

	block := BlockController{Router: api}
	block.Setup()

	info := InfoController{Router: api}
	info.Setup()

	tx := TxController{Router: api}
	tx.Setup()

	txCount := TxCountController{Router: api}
	txCount.Setup()

	accountCount := AccountCountController{Router: api}
	accountCount.Setup()

	txType := TxTypeController{Router: api}
	txType.Setup()

	market := MarketController{Router: api}
	market.Setup()

	version := VersionController{Router: api}
	version.Setup()
}
