package dtos

type GetAccountByAddressDTO struct {
	Address string `json:"address"`
}

type GetAccountsDTO struct {
	Page  uint32 `json:"page"`
	Count uint32 `json:"count"`
}

type TokensDTO map[string]struct {
	FreeAmount   uint64 `json:"free_amount"`
	StakedAmount uint64 `json:"staked_amount"`
	FrozenAmount uint64 `json:"frozen_amount"`
	VotedAmount  uint64 `json:"voted_amount"`
}

type CandidateUIDDTO struct {
	IDType string `json:"id_type"`
	ID     string `json:"id"`
}

type VoteListDTO []struct {
	CandidateUID     CandidateUIDDTO `json:"candidate_uid"`
	CandidateAddress string          `json:"candidate_address"`
	VotedBcoins      uint64          `json:"voted_bcoins"`
}

type AccountDTO struct {
	Address       string      `json:"address"`
	RegID         string      `json:"regid"`
	Tokens        TokensDTO   `json:"tokens"`
	ReceivedVotes uint64      `json:"received_votes"`
	VoteList      VoteListDTO `json:"vote_list"`
}

type AccountSketchDTO struct {
	Address string `json:"address"`
	Balance uint64 `json:"balance"`
	Vote    uint64 `json:"vote"`
}

type AccountListDTO struct {
	Data  []AccountSketchDTO `json:"data,omitempty"`
	Count uint32             `json:"count,omitempty"`
}
