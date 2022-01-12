package entity

import "github.com/shopspring/decimal"

type HoldConfig struct {
	StockList []StockList `json:"StockList" yaml:"StockList"`
}

type StockList struct {
	Stock StockConfig `json:"stock" yaml:"stock"`
}

type StockConfig struct {
	Name string `json:"name" yaml:"name"`
	Code string `json:"code" yaml:"code"`
	// 类型 1: 股票 2: 基金
	Type int `json:"type" yaml:"type"`
	// 持仓数量
	Count decimal.Decimal `json:"count" yaml:"count"`
	// 渠道 zfb
	Source string `json:"source" yaml:"source"`
}
