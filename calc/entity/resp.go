package entity

import "github.com/shopspring/decimal"

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
