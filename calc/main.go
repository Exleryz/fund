package main

import (
	"fmt"
	"fund/conf"
	"fund/service"
	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var holdConf *conf.HoldConfig

func init() {
	holdConf = &conf.HoldConfig{}
	readYml("conf/a.yml", holdConf)
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
	for i, v := range holdConf.StockList {
		fmt.Println(i, calc(v.Stock))
	}
}

func calc(stock conf.StockConfig) decimal.Decimal {
	switch stock.Type {
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
