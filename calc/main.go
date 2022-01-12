package main

import (
	"encoding/json"
	"fmt"
	"fund/calc/entity"
	"fund/conf"
	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/http"
)

var holdConf *entity.HoldConfig

func init() {
	holdConf = &entity.HoldConfig{}
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

func calc(stock entity.StockConfig) decimal.Decimal {
	switch stock.Type {
	case 2:
		// 获取对应 基金的 净值
		if stock.Source == "tt" {
			// 天天
			return stock.Count.Mul(getFundNetWorth(stock.Code)).Truncate(2)
		} else {
			// 支付宝
			return stock.Count.Mul(getFundNetWorth(stock.Code)).Round(2)
		}

	}
	return decimal.Decimal{}
}

func getFundNetWorth(code string) decimal.Decimal {
	// https://fund.xueqiu.com/dj/open/fund/growth/008975?day=30
	urlString := fmt.Sprintf("https://fund.xueqiu.com/dj/open/fund/growth/%s?day=30", code)
	request, _ := http.NewRequest("GET", urlString, nil)
	request.Header.Add("user-agent", conf.UA)
	/*request.Header.Add("referer", fmt.Sprintf("https://xq.com/S/F%s", code))

	// 获取 & 添加cookie
	cookies := xq.GetCookies("F" + code)
	for _, v := range cookies {
		request.AddCookie(v)
	}*/

	client := &http.Client{}
	do, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer do.Body.Close()
	all, err := ioutil.ReadAll(do.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	respDto := &entity.FundResp{}
	json.Unmarshal(all, respDto)

	return respDto.Data.FundNavGrowth[len(respDto.Data.FundNavGrowth)-1].Nav
}
