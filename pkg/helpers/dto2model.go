package helpers

import (
	"encoding/json"
	"scan/pkg/api/dtos"
	"scan/pkg/config"
	"scan/pkg/models"
	"scan/pkg/util"
)

func BlockDTO2Model(blockDTO dtos.BlockDTO) (block models.Block, ret bool) {
	ret = false
	tx, err := json.Marshal(blockDTO.Tx)
	if err != nil {
		return block, ret
	}

	block = models.Block{
		//ID:              0,
		Height:            blockDTO.Height,
		BlockHash:         blockDTO.BlockHash,
		MinerUID:          blockDTO.MinerUID,
		MinerAddress:      blockDTO.MinerAddress,
		Size:              blockDTO.Size,
		Version:           blockDTO.Version,
		MerkleRoot:        blockDTO.MerkleRoot,
		Time:              blockDTO.Time,
		PreviousBlockHash: blockDTO.PreviousBlockHash,
		NextBlockHash:     blockDTO.NextBlockHash,
		TxCount:           blockDTO.TxCount,
		Tx:                string(tx),
	}

	ret = true
	return block, ret
}

func TxDTO2Model(txDTO dtos.TxDTO) (tx models.Tx, ret bool) {
	ret = false
	receipts, err := json.Marshal(txDTO.Receipts)
	if err != nil {
		return tx, ret
	}
	candidateDTO, err := json.Marshal(txDTO.CandidateVotes)
	if err != nil {
		return tx, ret
	}
	transfers, err := json.Marshal(txDTO.Transfers)
	if err != nil {
		return
	}
	signatures, err := json.Marshal(txDTO.Signatures)
	if err != nil {
		return tx, ret
	}

	tx = models.Tx{
		//ID:              0,
		TxID:            txDTO.TxID,
		TxType:          txDTO.TxType,
		Version:         txDTO.Version,
		ValidHeight:     txDTO.ValidHeight,
		ConfirmedHeight: txDTO.ConfirmedHeight,
		ConfirmedTime:   txDTO.ConfirmedTime,
		BlockHash:       txDTO.BlockHash,
		Receipts:        string(receipts),
		RawTx:           txDTO.RawTx,
		TxUID:           txDTO.TxUID,
		FromAddr:        txDTO.FromAddr,
		ToUID:           txDTO.ToUID,
		ToAddr:          txDTO.ToAddr,
		FeeSymbol:       txDTO.FeeSymbol,
		Fees:            txDTO.Fees,
		Pubkey:          txDTO.PubKey,
		MinerPubkey:     txDTO.MinerPubKey,
		Signature:       txDTO.Signature,
		CandidateVotes:  string(candidateDTO),
		Transfers:       string(transfers),
		Memo:            util.DecodeHexString(txDTO.Memo),
		RequiredSigs:    txDTO.RequiredSigs,
		Signatures:      string(signatures),
		StakeType:       txDTO.StakeType,
		CoinSymbol:      txDTO.CoinSymbol,
		CoinAmount:      txDTO.CoinAmount,
		VMType:          txDTO.VMType,
		Upgradable:      txDTO.Upgradable,
		Code:            txDTO.Code,
		Abi:             txDTO.Abi,
		Arguments:       txDTO.Arguments,
		AssetName:       txDTO.AssetName,
		AssetSymbol:     txDTO.AssetSymbol,
		OwnerUID:        txDTO.OwnerUID,
		OwnerAddr:       txDTO.OwnerAddr,
		TotalSupply:     txDTO.TotalSupply,
		Mintable:        txDTO.Mintable,
		UpdateType:      txDTO.UpdateType,
		UpdateValue:     txDTO.UpdateValue,
		Domain:          util.DecodeHexString(txDTO.Domain),
		Key:             util.DecodeHexString(txDTO.Key),
		Value:           util.DecodeHexString(txDTO.Value),
	}

	ret = true

	return tx, ret
}

func AccountDTO2Model(accountDTO dtos.AccountDTO) (account models.Account, ret bool) {
	ret = false

	tokens := accountDTO.Tokens
	coinSymbol := config.GetRPCConfig().TokenSymbol
	if token, ok := tokens[coinSymbol]; ok {
		account = models.Account{
			//ID:      0,
			Address: accountDTO.Address,
			Balance: token.FreeAmount,
			Vote:    accountDTO.ReceivedVotes,
		}
	} else {
		account = models.Account{
			//ID:      0,
			Address: accountDTO.Address,
			Balance: 0,
			Vote:    accountDTO.ReceivedVotes,
		}
	}

	ret = true

	return account, ret
}
