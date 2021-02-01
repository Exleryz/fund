package main

import (
	"flag"
	"fmt"
	"fund/xq"
)

var stockCode string
var Period string

func init() {
	flag.StringVar(&stockCode, "sc", "SH000001", "A:SH/SZ,H:09988,M:TSLA")
	/**
	"1m"
	"5m"
	"15m"
	"60m"
	"120m"
	"day"
	"week"
	"quarter"
	"year"
	*/
	flag.StringVar(&Period, "p", "day", "A:SH/SZ,H:09988,M:TSLA")
	flag.Parse()
}

func main() {
	fmt.Println("starting get: ", stockCode, Period)
	//stockCode := "TSLA"

	xq.Xueqiu(stockCode, Period)
	//https://stock.xueqiu.com/v5/stock/chart/minute.json?symbol=SH513050&period=1d
	//https://stock.xueqiu.com/v5/stock/chart/kline.json?symbol=SH513050&begin=1611851971895&period=day&type=before&count=-284&indicator=kline,pe,pb,ps,pcf,market_capital,agt,ggt,balance
}
