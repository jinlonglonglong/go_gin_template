package helpers

import (
	"scan/pkg/api/dtos"
)

func FindRelatedAccounts(tx dtos.TxDTO) (accounts map[string]struct{}) {
	accounts = make(map[string]struct{})

	if len(tx.FromAddr) != 0 {
		accounts[tx.FromAddr] = struct{}{}
	}

	if len(tx.ToAddr) != 0 {
		accounts[tx.ToAddr] = struct{}{}
	}

	for _, receipt := range tx.Receipts {
		if len(receipt.FromAddr) != 0 {
			accounts[receipt.FromAddr] = struct{}{}
		}

		if len(receipt.ToAddr) != 0 {
			accounts[receipt.ToAddr] = struct{}{}
		}
	}

	for _, candidateVotes := range tx.CandidateVotes {
		accounts[candidateVotes.CandidateAddress] = struct{}{}
	}

	return accounts
}
