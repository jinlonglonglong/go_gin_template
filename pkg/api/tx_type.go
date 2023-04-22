package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scan/pkg/api/dtos"
	"scan/pkg/dao"
	"scan/pkg/util"
)

// TxTypeController 交易类型
type TxTypeController struct {
	Router *gin.RouterGroup
}

// Setup 初始化路由
func (controller TxTypeController) Setup() {
	handle := controller.Router.Group("tx_type")
	{
		handle.POST("list", controller.getTxTypes)
	}
}

func (controller *TxTypeController) getTxTypes(c *gin.Context) {
	if txTypeDTOs, ret := dao.GetTxTypesCache(); ret == true {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": txTypeDTOs})
		return
	}

	txTypes := dao.GetTxTypes()
	if len(txTypes) == 0 {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeNotExist, "error_msg": "not exist", "data": nil})
		return
	}

	txTypeDTOs := make([]dtos.TxTypeDTO, 0)
	for _, txType := range txTypes {
		txTypeDTO := dtos.TxTypeDTO{
			Enum:          txType.Enum,
			TxType:        txType.TxType,
			CNDescription: txType.CNDescription,
			ENDescription: txType.ENDescription,
		}

		txTypeDTOs = append(txTypeDTOs, txTypeDTO)
	}

	dao.SetTxTypesCache(txTypeDTOs)
	c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": txTypeDTOs})
}
