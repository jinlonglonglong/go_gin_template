package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"scan/pkg/api/dtos"
	"scan/pkg/config"
	"scan/pkg/dao"
	"scan/pkg/models"
	"scan/pkg/util"
)

// TxController 交易信息
type TxController struct {
	Router *gin.RouterGroup
}

// Setup 初始化路由
func (controller TxController) Setup() {
	handle := controller.Router.Group("tx")
	{
		handle.POST("count", controller.getTxCount)
		handle.POST("latest", controller.getLatestTxs)
		handle.POST("list", controller.getTxs)
		handle.POST("list/address", controller.getTxsByAccountAddress)
		handle.POST("list/hash", controller.getTxsByBlockHash)
		handle.POST("txid", controller.getTxByTxID)
	}
}

func (controller *TxController) getTxCount(c *gin.Context) {
	var count uint64
	if count, ret := dao.GetTxCountCache(); ret == true {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": count})
		return
	}

	count = uint64(dao.GetTxCountByTxType(""))

	dao.SetTxCountCache(count)
	c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": count})
}

func (controller *TxController) getTxByTxID(c *gin.Context) {
	var req dtos.GetTxByTxIDDTO
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParams, "error_msg": "invalid params", "data": nil})
		return
	}

	if txDTO, ret := dao.GetTxCache(req.TxID); ret == true {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": txDTO})
		return
	}

	tx := dao.GetTxByTxID(req.TxID)
	if (models.Tx{}) == tx {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeNotExist, "error_msg": "not exist", "data": nil})
		return
	}

	var receipts dtos.ReceiptsDTO
	json.Unmarshal([]byte(tx.Receipts), &receipts)
	var signatures []string
	json.Unmarshal([]byte(tx.Signatures), &signatures)
	var candidateVotes dtos.CandidateVotesDTO
	json.Unmarshal([]byte(tx.CandidateVotes), &candidateVotes)
	var transfers dtos.TransfersDTO
	json.Unmarshal([]byte(tx.Transfers), &transfers)

	txDTO := dtos.TxDTO{
		TxID:        tx.TxID,
		TxType:      tx.TxType,
		Version:     tx.Version,
		ValidHeight: tx.ValidHeight,
		// TODO:
		//Confirmations:   nil,
		ConfirmedHeight: tx.ConfirmedHeight,
		ConfirmedTime:   tx.ConfirmedTime,
		BlockHash:       tx.BlockHash,
		Receipts:        receipts,
		RawTx:           tx.RawTx,
		TxUID:           tx.TxUID,
		FromAddr:        tx.FromAddr,
		ToUID:           tx.ToUID,
		ToAddr:          tx.ToAddr,
		FeeSymbol:       tx.FeeSymbol,
		Fees:            tx.Fees,
		PubKey:          tx.Pubkey,
		MinerPubKey:     tx.MinerPubkey,
		Signature:       tx.Signature,
		CandidateVotes:  candidateVotes,
		Transfers:       transfers,
		Memo:            tx.Memo,
		RequiredSigs:    tx.RequiredSigs,
		Signatures:      signatures,
		StakeType:       tx.StakeType,
		CoinSymbol:      tx.CoinSymbol,
		CoinAmount:      tx.CoinAmount,
		VMType:          tx.VMType,
		Upgradable:      tx.Upgradable,
		Code:            tx.Code,
		Abi:             tx.Abi,
		Arguments:       tx.Arguments,
		AssetName:       tx.AssetName,
		AssetSymbol:     tx.AssetSymbol,
		OwnerUID:        tx.OwnerUID,
		OwnerAddr:       tx.OwnerAddr,
		TotalSupply:     tx.TotalSupply,
		Mintable:        tx.Mintable,
		UpdateType:      tx.UpdateType,
		UpdateValue:     tx.UpdateValue,
		Domain:          tx.Domain,
		Key:             tx.Key,
		Value:           tx.Value,
	}

	dao.SetTxCache(req.TxID, txDTO)
	c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": txDTO})
}

func (controller *TxController) getLatestTxs(c *gin.Context) {
	if txDTOs, ret := dao.GetLatestTxsCache(); ret == true {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": txDTOs})
		return
	}

	txs := dao.GetLatestTxs()
	if len(txs) == 0 {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeNotExist, "error_msg": "not exist", "data": nil})
		return
	}

	txDTOs := make([]dtos.TxDTO, 0, util.DefaultLatestTxs)
	for _, tx := range txs {
		var receipts dtos.ReceiptsDTO
		json.Unmarshal([]byte(tx.Receipts), &receipts)
		var signatures []string
		json.Unmarshal([]byte(tx.Signatures), &signatures)
		var candidateVotes dtos.CandidateVotesDTO
		json.Unmarshal([]byte(tx.CandidateVotes), &candidateVotes)
		var transfers dtos.TransfersDTO
		json.Unmarshal([]byte(tx.Transfers), &transfers)

		txDTO := dtos.TxDTO{
			TxID:        tx.TxID,
			TxType:      tx.TxType,
			Version:     tx.Version,
			ValidHeight: tx.ValidHeight,
			// TODO:
			//Confirmations:   nil,
			ConfirmedHeight: tx.ConfirmedHeight,
			ConfirmedTime:   tx.ConfirmedTime,
			BlockHash:       tx.BlockHash,
			Receipts:        receipts,
			RawTx:           tx.RawTx,
			TxUID:           tx.TxUID,
			FromAddr:        tx.FromAddr,
			ToUID:           tx.ToUID,
			ToAddr:          tx.ToAddr,
			FeeSymbol:       tx.FeeSymbol,
			Fees:            tx.Fees,
			PubKey:          tx.Pubkey,
			MinerPubKey:     tx.MinerPubkey,
			Signature:       tx.Signature,
			CandidateVotes:  candidateVotes,
			Transfers:       transfers,
			Memo:            tx.Memo,
			RequiredSigs:    tx.RequiredSigs,
			Signatures:      signatures,
			StakeType:       tx.StakeType,
			CoinSymbol:      tx.CoinSymbol,
			CoinAmount:      tx.CoinAmount,
			VMType:          tx.VMType,
			Upgradable:      tx.Upgradable,
			Code:            tx.Code,
			Abi:             tx.Abi,
			Arguments:       tx.Arguments,
			AssetName:       tx.AssetName,
			AssetSymbol:     tx.AssetSymbol,
			OwnerUID:        tx.OwnerUID,
			OwnerAddr:       tx.OwnerAddr,
			TotalSupply:     tx.TotalSupply,
			Mintable:        tx.Mintable,
			UpdateType:      tx.UpdateType,
			UpdateValue:     tx.UpdateValue,
		}
		txDTOs = append(txDTOs, txDTO)
	}

	dao.SetLatestTxsCache(txDTOs)
	c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": txDTOs})
}

func (controller *TxController) getTxs(c *gin.Context) {
	var req dtos.GetTxsDTO
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParams, "error_msg": "invalid params", "data": nil})
		return
	}

	page := req.Page
	count := req.Count
	if (count == 0 || count > util.DefaultTxPageSize) || (page == 0 || page > util.DefaultMaxTxPage) {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParamRange, "error_msg": "invalid param range", "data": nil})
		return
	}

	txType := req.TxType
	txs := dao.GetTxs(page, count, txType)
	if len(txs) == 0 {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeNotExist, "error_msg": "not exist", "data": nil})
		return
	}

	txDTOs := make([]dtos.TxDTO, 0, len(txs))
	for _, tx := range txs {
		var receipts dtos.ReceiptsDTO
		json.Unmarshal([]byte(tx.Receipts), &receipts)
		var signatures []string
		json.Unmarshal([]byte(tx.Signatures), &signatures)
		var candidateVotes dtos.CandidateVotesDTO
		json.Unmarshal([]byte(tx.CandidateVotes), &candidateVotes)
		var transfers dtos.TransfersDTO
		json.Unmarshal([]byte(tx.Transfers), &transfers)

		txDTO := dtos.TxDTO{
			TxID:        tx.TxID,
			TxType:      tx.TxType,
			Version:     tx.Version,
			ValidHeight: tx.ValidHeight,
			// TODO:
			//Confirmations:   nil,
			ConfirmedHeight: tx.ConfirmedHeight,
			ConfirmedTime:   tx.ConfirmedTime,
			BlockHash:       tx.BlockHash,
			Receipts:        receipts,
			RawTx:           tx.RawTx,
			TxUID:           tx.TxUID,
			FromAddr:        tx.FromAddr,
			ToUID:           tx.ToUID,
			ToAddr:          tx.ToAddr,
			FeeSymbol:       tx.FeeSymbol,
			Fees:            tx.Fees,
			PubKey:          tx.Pubkey,
			MinerPubKey:     tx.MinerPubkey,
			Signature:       tx.Signature,
			CandidateVotes:  candidateVotes,
			Transfers:       transfers,
			Memo:            tx.Memo,
			RequiredSigs:    tx.RequiredSigs,
			Signatures:      signatures,
			StakeType:       tx.StakeType,
			CoinSymbol:      tx.CoinSymbol,
			CoinAmount:      tx.CoinAmount,
			VMType:          tx.VMType,
			Upgradable:      tx.Upgradable,
			Code:            tx.Code,
			Abi:             tx.Abi,
			Arguments:       tx.Arguments,
			AssetName:       tx.AssetName,
			AssetSymbol:     tx.AssetSymbol,
			OwnerUID:        tx.OwnerUID,
			OwnerAddr:       tx.OwnerAddr,
			TotalSupply:     tx.TotalSupply,
			Mintable:        tx.Mintable,
			UpdateType:      tx.UpdateType,
			UpdateValue:     tx.UpdateValue,
		}
		txDTOs = append(txDTOs, txDTO)
	}

	records := dao.GetTxCountByTxType(req.TxType)
	txListDTO := dtos.TxListDTO{
		Data:  txDTOs,
		Count: records,
	}

	c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": txListDTO})
}

func (controller *TxController) getTxsByBlockHash(c *gin.Context) {
	var req dtos.GetTxsByBlockHashDTO
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParams, "error_msg": "invalid params", "data": nil})
		return
	}

	hash := req.BlockHash
	if len(hash) != util.BlockHashLength {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParamRange, "error_msg": "invalid block hash length", "data": nil})
		return
	}

	page := req.Page
	count := req.Count
	if (count == 0 || count > util.DefaultTxPageSize) || (page == 0 || page > util.DefaultMaxTxPage) {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParamRange, "error_msg": "invalid param range", "data": nil})
		return
	}

	txs := dao.GetTxsByBlockHash(page, count, hash)
	if len(txs) == 0 {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeNotExist, "error_msg": "not exist", "data": nil})
		return
	}

	txDTOs := make([]dtos.TxDTO, 0, len(txs))
	for _, tx := range txs {
		var receipts dtos.ReceiptsDTO
		json.Unmarshal([]byte(tx.Receipts), &receipts)
		var signatures []string
		json.Unmarshal([]byte(tx.Signatures), &signatures)
		var candidateVotes dtos.CandidateVotesDTO
		json.Unmarshal([]byte(tx.CandidateVotes), &candidateVotes)
		var transfers dtos.TransfersDTO
		json.Unmarshal([]byte(tx.Transfers), &transfers)

		txDTO := dtos.TxDTO{
			TxID:        tx.TxID,
			TxType:      tx.TxType,
			Version:     tx.Version,
			ValidHeight: tx.ValidHeight,
			// TODO:
			//Confirmations:   nil,
			ConfirmedHeight: tx.ConfirmedHeight,
			ConfirmedTime:   tx.ConfirmedTime,
			BlockHash:       tx.BlockHash,
			Receipts:        receipts,
			RawTx:           tx.RawTx,
			TxUID:           tx.TxUID,
			FromAddr:        tx.FromAddr,
			ToUID:           tx.ToUID,
			ToAddr:          tx.ToAddr,
			FeeSymbol:       tx.FeeSymbol,
			Fees:            tx.Fees,
			PubKey:          tx.Pubkey,
			MinerPubKey:     tx.MinerPubkey,
			Signature:       tx.Signature,
			CandidateVotes:  candidateVotes,
			Transfers:       transfers,
			Memo:            tx.Memo,
			RequiredSigs:    tx.RequiredSigs,
			Signatures:      signatures,
			StakeType:       tx.StakeType,
			CoinSymbol:      tx.CoinSymbol,
			CoinAmount:      tx.CoinAmount,
			VMType:          tx.VMType,
			Upgradable:      tx.Upgradable,
			Code:            tx.Code,
			Abi:             tx.Abi,
			Arguments:       tx.Arguments,
			AssetName:       tx.AssetName,
			AssetSymbol:     tx.AssetSymbol,
			OwnerUID:        tx.OwnerUID,
			OwnerAddr:       tx.OwnerAddr,
			TotalSupply:     tx.TotalSupply,
			Mintable:        tx.Mintable,
			UpdateType:      tx.UpdateType,
			UpdateValue:     tx.UpdateValue,
		}
		txDTOs = append(txDTOs, txDTO)
	}

	records := dao.GetTxCountByBlockHash(hash)
	txListDTO := dtos.TxListDTO{
		Data:  txDTOs,
		Count: records,
	}

	c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": txListDTO})
}

func (controller *TxController) getTxsByAccountAddress(c *gin.Context) {
	var req dtos.GetTxsByAccountAddress
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParams, "error_msg": "invalid params", "data": nil})
		return
	}

	addr := req.Address
	if len(addr) != config.GetRPCConfig().AddressLength {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParamRange, "error_msg": "invalid account address length", "data": nil})
		return
	}

	page := req.Page
	count := req.Count
	if (count == 0 || count > util.DefaultTxPageSize) || (page == 0 || page > util.DefaultMaxTxPage) {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeInvalidParamRange, "error_msg": "invalid param range", "data": nil})
		return
	}

	accountTxs := dao.GetAccountTxs(addr, page, count)
	if len(accountTxs) == 0 {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeNotExist, "error_msg": "not exist", "data": nil})
		return
	}

	txids := make([]string, 0, len(accountTxs))
	for _, tx := range accountTxs {
		txids = append(txids, tx.TxID)
	}

	txs := dao.GetTxsByTxIDs(txids)
	if len(txs) == 0 {
		c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeNotExist, "error_msg": "not exist", "data": nil})
		return
	}

	txDTOs := make([]dtos.TxDTO, 0, len(txs))
	for _, tx := range txs {
		var receipts dtos.ReceiptsDTO
		json.Unmarshal([]byte(tx.Receipts), &receipts)
		var signatures []string
		json.Unmarshal([]byte(tx.Signatures), &signatures)
		var candidateVotes dtos.CandidateVotesDTO
		json.Unmarshal([]byte(tx.CandidateVotes), &candidateVotes)
		var transfers dtos.TransfersDTO
		json.Unmarshal([]byte(tx.Transfers), &transfers)

		txDTO := dtos.TxDTO{
			TxID:        tx.TxID,
			TxType:      tx.TxType,
			Version:     tx.Version,
			ValidHeight: tx.ValidHeight,
			// TODO:
			//Confirmations:   nil,
			ConfirmedHeight: tx.ConfirmedHeight,
			ConfirmedTime:   tx.ConfirmedTime,
			BlockHash:       tx.BlockHash,
			Receipts:        receipts,
			RawTx:           tx.RawTx,
			TxUID:           tx.TxUID,
			FromAddr:        tx.FromAddr,
			ToUID:           tx.ToUID,
			ToAddr:          tx.ToAddr,
			FeeSymbol:       tx.FeeSymbol,
			Fees:            tx.Fees,
			PubKey:          tx.Pubkey,
			MinerPubKey:     tx.MinerPubkey,
			Signature:       tx.Signature,
			CandidateVotes:  candidateVotes,
			Transfers:       transfers,
			Memo:            tx.Memo,
			RequiredSigs:    tx.RequiredSigs,
			Signatures:      signatures,
			StakeType:       tx.StakeType,
			CoinSymbol:      tx.CoinSymbol,
			CoinAmount:      tx.CoinAmount,
			VMType:          tx.VMType,
			Upgradable:      tx.Upgradable,
			Code:            tx.Code,
			Abi:             tx.Abi,
			Arguments:       tx.Arguments,
			AssetName:       tx.AssetName,
			AssetSymbol:     tx.AssetSymbol,
			OwnerUID:        tx.OwnerUID,
			OwnerAddr:       tx.OwnerAddr,
			TotalSupply:     tx.TotalSupply,
			Mintable:        tx.Mintable,
			UpdateType:      tx.UpdateType,
			UpdateValue:     tx.UpdateValue,
		}
		txDTOs = append(txDTOs, txDTO)
	}

	records := dao.GetAccountTxCount(addr)
	txListDTO := dtos.TxListDTO{
		Data:  txDTOs,
		Count: records,
	}

	c.JSON(http.StatusOK, gin.H{"error_code": util.ErrorCodeSuccess, "error_msg": "success", "data": txListDTO})
}
