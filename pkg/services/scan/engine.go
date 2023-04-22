package scan

import (
	"context"
	"net/http"
	"scan/pkg/api/dtos"
	"scan/pkg/dao"
	"scan/pkg/data"
	"scan/pkg/helpers"
	"scan/pkg/models"
	"scan/pkg/util"
	"scan/pkg/ws"
	"time"

	"github.com/go-errors/errors"
	"github.com/golang/glog"
)

type Engine struct {
	ctx        context.Context
	infoDTO    dtos.InfoDTO
	httpClient *http.Client
}

func NewEngine(ctx context.Context) (*Engine, error) {
	return &Engine{
		ctx: ctx,
		infoDTO: dtos.InfoDTO{
			//TotalIssuance:    0,
			//TotalCirculation: 0,
			TotalAccounts: 0,
			TotalTokens:   0,
			TotalBlocks:   0,
			TotalTxs:      0,
		},
		httpClient: data.MustGetHttp("node"),
	}, nil
}

func (engine *Engine) Run(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			goErr := errors.Wrap(err, 1)
			glog.Errorf("%s\n stack: %s", goErr.Error(), goErr.Stack())
		}
	}()

	syncTicker := time.NewTicker(time.Second * 3)
	defer func() {
		syncTicker.Stop()
	}()

	processCountTicker := time.NewTicker(time.Minute * 10)
	defer func() {
		processCountTicker.Stop()
	}()

	for {
		select {
		case <-syncTicker.C:
			engine.sync()
		case <-processCountTicker.C:
			engine.processCount()
		case <-ctx.Done():
			break
		}
	}
}

func (engine *Engine) Run2(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			goErr := errors.Wrap(err, 1)
			glog.Errorf("%s\n stack: %s", goErr.Error(), goErr.Stack())
		}
	}()

	processInfoTicker := time.NewTicker(time.Second * 30)
	defer func() {
		processInfoTicker.Stop()
	}()

	engine.processInfo()

	for {
		select {
		case <-processInfoTicker.C:
			engine.processInfo()
		case <-ctx.Done():
			break
		}
	}
}

func (engine *Engine) sync() {
	getInfoReq := &helpers.CommonReq{JsonRpc: "1.0", Id: "curltest", Method: "getinfo"}
	getInfoRsp, err := helpers.SendGetInfo(getInfoReq)
	if err != nil {
		glog.Errorf("call getinfo rpc failed: %s", err.Error())
		return
	}

	infoDTOInternal := getInfoRsp.Result
	targetSyncBlockHeight := infoDTOInternal.SynBlockHeight
	currentSyncBlockHeight := helpers.GetLatestBlockHeight()

	// nothing to do here
	if currentSyncBlockHeight == targetSyncBlockHeight {
		return
	}

	for currentSyncBlockHeight++; currentSyncBlockHeight <= targetSyncBlockHeight; currentSyncBlockHeight++ {
		// 1. fetch data from node
		getBlockReq := &helpers.CommonReq{JsonRpc: "1.0", Id: "curltest", Method: "getblock", Params: []interface{}{currentSyncBlockHeight}}
		getBlockRsp, err := helpers.SendGetBlock(getBlockReq)
		if err != nil {
			glog.Errorf("call getblock rpc failed: %s", err.Error())
			return
		}

		// 2. write block into db
		blockDTO := getBlockRsp.Result
		block, _ := helpers.BlockDTO2Model(blockDTO)
		dao.SetBlock(block)
		dao.UpdatePreviousBlock(block)

		if targetSyncBlockHeight-currentSyncBlockHeight < util.DefaultDisableWebSocketBlockHeight {
			ws.BroadcastChannel.Broadcast(ws.Message{
				Type: util.MessageTypeBlock,
				Data: blockDTO,
			})
		}

		// 3. write tx into db
		for _, txid := range blockDTO.Tx {
			getTxReq := &helpers.CommonReq{JsonRpc: "1.0", Id: "curltest", Method: "gettxdetail", Params: []interface{}{txid}}
			getTxRsp, err := helpers.SendGetTx(getTxReq)
			if err != nil {
				glog.Errorf("call gettxdetail rpc failed: %s", err.Error())
				return
			}

			txDTO := getTxRsp.Result
			tx, _ := helpers.TxDTO2Model(txDTO)
			dao.WriteTx(tx)

			accounts := helpers.FindRelatedAccounts(txDTO)
			for addr, _ := range accounts {
				accountTx := models.AccountTx{
					//ID:   0,
					Addr: addr,
					TxID: txid,
				}
				dao.SetAccountTx(accountTx)

				getAccountReq := &helpers.CommonReq{JsonRpc: "1.0", Id: "curltest", Method: "getaccountinfo", Params: []interface{}{addr}}
				getAccountRsp, err := helpers.SendGetAccount(getAccountReq)
				if err != nil {
					return
				}

				accountDTO := getAccountRsp.Result
				account, _ := helpers.AccountDTO2Model(accountDTO)
				if oldAccount, ret := dao.AccountExist(accountDTO.Address); ret {
					account.ID = oldAccount.ID
					account.CreateTime = oldAccount.CreateTime
				} else {
					account.CreateTime = tx.ConfirmedTime
				}
				dao.SetAccount(account) // update or create
			}

			if targetSyncBlockHeight-currentSyncBlockHeight < util.DefaultDisableWebSocketBlockHeight {
				ws.BroadcastChannel.Broadcast(ws.Message{
					Type: util.MessageTypeTx,
					Data: txDTO,
				})
			}
		}

		glog.Infof("currentSyncBlockHeight: %d, targetSyncBlockHeight: %d", currentSyncBlockHeight, targetSyncBlockHeight)
	}

	engine.processInfo()

	ws.BroadcastChannel.Broadcast(ws.Message{
		Type: util.MessageTypeInfo,
		Data: engine.infoDTO,
	})
}

func (engine *Engine) processCount() {
	timestamp := helpers.GetLatestTxCountDate()
	// 时间戳对应当天零点，实际上覆盖的时间戳范围为 [timestamp, timestamp + secondsOfDay)
	timestamp += util.SecondsOfDay

	for ; timestamp < util.GetZeroTimestampOfToday(); timestamp += util.SecondsOfDay {
		dao.SetTxCount(timestamp)
	}

	timestamp2 := helpers.GetLatestAccountCountDate()
	// 时间戳对应当天零点，实际上覆盖的时间戳范围为 [timestamp, timestamp + secondsOfDay)
	timestamp2 += util.SecondsOfDay

	for ; timestamp2 < util.GetZeroTimestampOfToday(); timestamp2 += util.SecondsOfDay {
		dao.SetAccountCount(timestamp2)
	}
}

func (engine *Engine) processInfo() {
	engine.infoDTO.TotalAccounts = dao.GetAccountCount()
	engine.infoDTO.TotalTokens = dao.GetAssetCount()
	engine.infoDTO.TotalBlocks = dao.GetBlockCount()
	engine.infoDTO.TotalTxs = dao.GetTxCountByTxType("")

	dao.SetInfoCache(engine.infoDTO)
}
