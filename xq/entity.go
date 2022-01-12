package xq

import "github.com/shopspring/decimal"

type Resp struct {
	Data             RespData `json:"data"`
	ErrorCode        int      `json:"error_code"`
	ErrorDescription string   `json:"error_description"`
}

type RespData struct {
	Symbol string        `json:"symbol"`
	Column []string      `json:"column"`
	Items  []interface{} `json:"item"`
}

type FundResp struct {
	Data struct {
		FundNavGrowth []struct {
			// 日期
			Date string `json:"date"`
			// 净值
			Nav        decimal.Decimal `json:"nav"`
			Percentage string          `json:"percentage"`
			ThanValue  string          `json:"than_value"`
			// 净值走势
			Value string `json:"value"`
		} `json:"fund_nav_growth"`
		GrowthLines []string `json:"growth_lines"`
	} `json:"data" yaml:"data"`
}
