package jsl

type Resp struct {
}

type RespCell struct {
	// 日期
	HistDt string `json:"hist_dt"`
	// 净值
	EstimateValue string `json:"estimate_value"`
	//
	EstimateDiff string `json:"estimate_diff"`
	// 收盘价
	TradePrice string `json:"trade_price"`
	FundNav    string `json:"fund_nav"`
	// 溢价率
	DiscountRt string `json:"discount_rt"`
	// 指数涨幅
	IdxIncrRt string `json:"idx_incr_rt"`
	// 指数PB
	IdxPb string `json:"idx_pb"`
	// 指数PE
	IdxPe string `json:"idx_pe"`
	// 场内份额
	Amount int `json:"amount"`
	// 场内新增
	AmountIncr int `json:"amount_incr"`
	// 份额涨幅
	IncreaseRt string `json:"increase_rt"`
}
