package conf

import "github.com/shopspring/decimal"

// DataSource 数据来源
type DataSource int

const (
	// XueQiu 雪球
	XueQiu DataSource = iota
	// JiSiLu 集思录
	JiSiLu
	// SINA 新浪
	SINA
)

// AssetsType 资产类型
type AssetsType int

const (
	// Stock 股票
	Stock AssetsType = iota
	// Fund 基金
	Fund
)

// BuyChannel 购买渠道
type BuyChannel int

const (
	// ZFB 支付宝 0
	ZFB BuyChannel = iota
	// HB 华宝 1
	HB
	// ZSYH 招商银行 2
	ZSYH
	// DFCF 东方财富 3
	DFCF
	// RY 睿远 4
	RY
	// TT 天天 5
	TT
)

type StockConfig struct {
	Name string `json:"name" yaml:"name"`
	Code string `json:"code" yaml:"code"`
	// 类型 0: 股票 1: 基金
	Type AssetsType `json:"type" yaml:"type"`
	// 持仓数量
	Count decimal.Decimal `json:"count" yaml:"count"`
	// 渠道 zfb
	Source BuyChannel `json:"source" yaml:"source"`
}

type HoldConfig struct {
	StockList []StockList `json:"StockList" yaml:"StockList"`
}

type StockList struct {
	Stock StockConfig `json:"stock" yaml:"stock"`
}
