package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"scan/pkg/api/dtos"
	config "scan/pkg/config"
	"scan/pkg/data"
)

/*
{
	"jsonrpc": "1.0",
	"id": "curltest",
	"method": "getinfo",
	"params": []
}
*/
type CommonReq struct {
	JsonRpc string        `json:"jsonrpc"`
	Id      string        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

/*
{
	"result": {
		"version": "v2.1.2.2-a9f094f-dirty-release-linux (2019-12-02 15:32:44 +0800)",
		"protocol_version": 10001,
		"net_type": "REGTEST_NET",
		"proxy": "",
		"ext_ip": "",
		"conf_dir": "/opt/ibs/./IBS.conf",
		"data_dir": "/opt/ibs/./regtest",
		"block_interval": 10,
		"genblock": 1,
		"total_coins" : 20789999900000000,
		"total_accounts" : 14,
		"total_contracts" : 1,
		"wallet_balance": 207900000.00000000,
		"relay_fee_perkb": 0.00001000,
		"tipblock_fuel_rate": 1,
		"tipblock_fuel": 0,
		"tipblock_time": 1575781549,
		"tipblock_hash": "6d57f35d4bf930cb3d810e25ffc0ebe30a02c09bf042a064f8e4c2ac34af44d9",
		"tipblock_height": 51399,
		"synblock_height": 51399,
		"connections": 0,
		"errors": ""
	},
	"error": null,
	"id": "curltest"
}
*/

type GetInfoRsp struct {
	Id     string               `json:"id"`
	Error  interface{}          `json:"error"`
	Result dtos.InfoDTOInternal `json:"result"`
}

func SendGetInfo(request *CommonReq) (*GetInfoRsp, error) {
	client := data.MustGetHttp("node")
	rpcConf := config.GetRPCConfig()

	// create request
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", rpcConf.URL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(rpcConf.UserName, rpcConf.Password)
	req.Header.Set("Content-Type", "application/json;encoding=utf-8;")

	// execute request
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	// check response status
	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d error: %s", rsp.StatusCode, rsp.Status)
	}

	// build return
	response := new(GetInfoRsp)
	if err := json.NewDecoder(rsp.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}

/*
{
	"result": {
		"block_hash": "b7cd972c9bdfe5fc0f96ddb926d3c4f5a8337508c42d93d5aeae1fb84b9cf8fc",
		"miner_uid": "0-9",
		"miner_address": "wWLCNZhng3hrUDfG8FuBgqYc59nfbBbtxF",
		"confirmations": 10,
		"size": 169,
		"height": 52500,
		"version": 1,
		"merkle_root": "d5a275c544cfa7b255b53db840faeafc6ccf3032cf63f9c60a0afdd074e39fa2",
		"tx_count": 1,
		"tx": ["d5a275c544cfa7b255b53db840faeafc6ccf3032cf63f9c60a0afdd074e39fa2"],
		"time": 1575792559,
		"previous_block_hash": "aa8e5706636ee32d08977bc5d53b83b3a9e6f640d5a27188c4b2cdf2b99fe767",
		"next_block_hash": "f46f0ddffbfa5ec1aa33d6df5391d545edcc5f2e493ded5ee5dc4a2338882cf1"
	},
	"error": null,
	"id": "curltest"
}
*/

type GetBlockRsp struct {
	Id     string        `json:"id"`
	Error  interface{}   `json:"error"`
	Result dtos.BlockDTO `json:"result"`
}

func SendGetBlock(request *CommonReq) (*GetBlockRsp, error) {
	client := data.MustGetHttp("node")
	rpcConf := config.GetRPCConfig()

	// create request
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", rpcConf.URL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(rpcConf.UserName, rpcConf.Password)
	req.Header.Set("Content-Type", "application/json;encoding=utf-8;")

	// execute request
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	// check response status
	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d error: %s", rsp.StatusCode, rsp.Status)
	}

	// build return
	response := new(GetBlockRsp)
	if err := json.NewDecoder(rsp.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}

/*
{
	"result": {
		"txid": "b7761b01a623674b287b64318eb30467416de0955709d3f246b7ca3ee0951c1c",
		"tx_type": "BLOCK_REWARD_TX",
		"version": 1,
		"tx_uid": "0-4",
		"to_addr": "wQVy5YkzwPWe1FUeGfGWRRBZrp6UsvtCnN",
		"valid_height": 121704,
		"coin_symbol": "IBS"
		"coin_amount": 0,
		"confirmations": 180,
		"confirmed_height": 121704,
		"confirmed_time": 1576484731,
		"block_hash": "55db5a26495a4ad479db4186be4640cd11d9b7c4bfc94f8f138fdb40401f8513",
		"receipts": [{
			"from_addr": "",
			"to_addr": "wQVy5YkzwPWe1FUeGfGWRRBZrp6UsvtCnN",
			"coin_symbol": "IBS",
			"coin_amount": 0,
			"receipt_code": 101,
			"memo": "block reward to miner"
		}],
		"rawtx": "010186b56802000400"
	},
	"error": null,
	"id": "curltest"
}
*/

type GetTxRsp struct {
	Id     string      `json:"id"`
	Error  interface{} `json:"error"`
	Result dtos.TxDTO  `json:"result"`
}

func SendGetTx(request *CommonReq) (*GetTxRsp, error) {
	client := data.MustGetHttp("node")
	rpcConf := config.GetRPCConfig()

	// create request
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", rpcConf.URL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(rpcConf.UserName, rpcConf.Password)
	req.Header.Set("Content-Type", "application/json;encoding=utf-8;")

	// execute request
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	// check response status
	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d error: %s", rsp.StatusCode, rsp.Status)
	}

	// build return
	response := new(GetTxRsp)
	if err := json.NewDecoder(rsp.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}

/*
{
	"result": {
		"address": "wYKpRM6Vwi2g9DpmaPGkrUefDjiB1JodbM",
		"keyid": "8b44ece8e24d5126da44cb787ec54d9962860f63",
		"nickid": "",
		"regid": "0-1",
		"regid_mature": true,
		"owner_pubkey": "03986936d27211dc30b947676ab6117008ea554d62336b7ef8277b2e973300e21e",
		"miner_pubkey": "",
		"tokens": {
			"IBS": {
				"free_amount": 20789944896800000,
				"staked_amount": 0,
				"frozen_amount": 0,
				"voted_amount": 210000000000000
			},
			"CNY": {
				"free_amount": 9999999999900000,
				"staked_amount": 0,
				"frozen_amount": 0,
				"voted_amount": 0
			}
		},
		"received_votes": 0,
		"vote_list": [{
			"candidate_uid": {
				"id_type": "PubKey",
				"id": "03beef5987cb50dceab0242a660cdab2e0e8911d42e2b2fe77eb35b4520d9092c3"
			},
			"voted_bcoins": 210000000000000
		}],
		"position": "inblock"
	},
	"error": null,
	"id": "curltest"
}
*/
type GetAccountInfoRsp struct {
	Id     string          `json:"id"`
	Error  interface{}     `json:"error"`
	Result dtos.AccountDTO `json:"result"`
}

func SendGetAccount(request *CommonReq) (*GetAccountInfoRsp, error) {
	client := data.MustGetHttp("node")
	rpcConf := config.GetRPCConfig()

	// create request
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", rpcConf.URL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(rpcConf.UserName, rpcConf.Password)
	req.Header.Set("Content-Type", "application/json;encoding=utf-8;")

	// execute request
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	// check response status
	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d error: %s", rsp.StatusCode, rsp.Status)
	}

	// build return
	response := new(GetAccountInfoRsp)
	if err := json.NewDecoder(rsp.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}
