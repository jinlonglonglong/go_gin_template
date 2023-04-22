package dtos

type MarketDTO struct {
	CNYPrice        float32 `json:"cny_price"`
	CNYMarketValue  uint64  `json:"cny_market_value"`
	CNYTurnover     uint64  `json:"cny_turnover"`
	USDPrice        float32 `json:"usd_price"`
	USDYMarketValue uint64  `json:"usd_market_value"`
	USDTurnover     uint64  `json:"usd_turnover"`
	Changed         float32 `json:"changed"` // 24h
}
