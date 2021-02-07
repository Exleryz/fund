package backtest

import "github.com/shopspring/decimal"

type BackTest interface {
	// 盘前
	before()
	// 盘中
	ing()
	// 盘后
	after()
}

// 股票池 持有
type HoldStock struct {
	// 股票代码
	Code string
	// 持股数
	Count int
	// 成本价
	Cost decimal.Decimal
}

type TranHistory struct {
	// 交易类型 买入 1 卖出 -1
	Type int
	// 价格
	Price decimal.Decimal
	// 股数
	Count int
}
