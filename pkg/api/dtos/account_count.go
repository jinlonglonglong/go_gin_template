package dtos

type GetLatestAccountCountsDTO struct {
	Count uint32 `json:"count"`
}

type AccountCountDTO struct {
	AccountCount uint32 `json:"account_count"`
	Date         string `json:"date"`
}
