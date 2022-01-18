package main

import (
	"flag"
	"fmt"
	"fund/conf"
	"fund/service"
	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var holdConf *conf.HoldConfig
var confFilePath *string

func init() {
	confFilePath = flag.String("c", "calc/conf/position.yml", "config file path, default: conf/position.yml")
	flag.Parse()

	holdConf = &conf.HoldConfig{}
	readYml(*confFilePath, holdConf)
	log.Println(holdConf)
}

// readYml 读取yml文件
func readYml(filePath string, config interface{}) {
	data, _ := ioutil.ReadFile(filePath)

	err := yaml.Unmarshal(data, config)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Println("获取", filePath, "数据成功")
	}
}

func main() {
	// 对数据进行整理求和
	mapping := make(map[string]decimal.Decimal, len(holdConf.StockList))
	for _, v := range holdConf.StockList {
		calcDecimal := calc(v.Stock)
		fmt.Println(v.Stock.Name, calcDecimal)

		if value, ok := mapping[v.Stock.Code]; ok {
			// 有值 需要添加
			mapping[v.Stock.Code] = value.Add(calcDecimal)
		} else {
			mapping[v.Stock.Code] = calcDecimal
		}
	}

	fmt.Println()

	sum := decimal.Decimal{}
	// 输出
	for _, v := range holdConf.StockList {
		if value, ok := mapping[v.Stock.Code]; ok {
			sum = sum.Add(value)
			fmt.Println(v.Stock.Name, value)
			delete(mapping, v.Stock.Code)
		}
	}

	fmt.Println(sum, "仍需努力")
}

func calc(stock conf.StockConfig) decimal.Decimal {
	switch stock.Type {
	case conf.Stock:
		// 	//period := "day"
		//	//period:= "week"
		stockService := service.NewStockService(conf.XueQiu, "day")
		return stock.Count.Mul(stockService.GetPrice(stock.Code))
	case conf.Fund:
		fundService := service.NewFundService(conf.XueQiu)

		// 获取对应 基金的 净值
		if stock.Source == conf.TT {
			// 天天
			return stock.Count.Mul(fundService.GetPrice(stock.Code)).Truncate(2)
		} else {
			// 支付宝
			return stock.Count.Mul(fundService.GetPrice(stock.Code)).Round(2)
		}

	}
	return decimal.Decimal{}
}
