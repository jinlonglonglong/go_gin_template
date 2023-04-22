package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scan/pkg/api/dtos"
	"scan/pkg/dao"
	"scan/pkg/helpers"
	"scan/pkg/util"
)

// AccountController 账户信息
type AccountController struct {
	Router *gin.RouterGroup
}

// Setup 初始化路由
func (controller AccountController) Setup() {
	handle := controller.Router.Group("account")
	{
		handle.POST("address", controller.getAccountByAddress)
		handle.POST("list", controller.getAccounts)
	}
}

func (controller *AccountController) getAccountByAddress(c *gin.Context) {
	var req dtos.GetAccountByAddressDTO
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParams, "error_msg": "invalid params", "data": nil})
		return
	}

	if accountDTO, ret := dao.GetAccountCache(req.Address); ret == true {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": accountDTO})
		return
	}

	rpcReq := &helpers.CommonReq{JsonRpc: "1.0", Id: "curltest", Method: "getaccountinfo", Params: []interface{}{req.Address}}
	rsp, err := helpers.SendGetAccount(rpcReq)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInternalError, "error_msg": "internal error", "data": nil})
		return
	}

	accountDTO := rsp.Result
	dao.SetAccountCache(req.Address, accountDTO)
	c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": accountDTO})
}

func (controller *AccountController) getAccounts(c *gin.Context) {
	var req dtos.GetAccountsDTO
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParams, "error_msg": "invalid params", "data": nil})
		return
	}

	page := req.Page
	count := req.Count
	if (count == 0 || count > util.DefaultAccountPageSize) || (page == 0 || page > util.DefaultMaxAccountPage) {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParamRange, "error_msg": "invalid param range", "data": nil})
		return
	}

	accounts := dao.GetAccounts(page, count)
	if len(accounts) == 0 {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeNotExist, "error_msg": "not exist", "data": nil})
		return
	}

	accountSketchDTOs := make([]dtos.AccountSketchDTO, 0, len(accounts))
	for _, account := range accounts {
		accountSketchDTO := dtos.AccountSketchDTO{
			Address: account.Address,
			Balance: account.Balance,
			Vote:    account.Vote,
		}

		accountSketchDTOs = append(accountSketchDTOs, accountSketchDTO)
	}

	records := dao.GetAccountCount()
	accountListDTO := dtos.AccountListDTO{
		Data:  accountSketchDTOs,
		Count: records,
	}

	c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": accountListDTO})
}
