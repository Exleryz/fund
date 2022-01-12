package xq

import (
	"encoding/json"
	"fmt"
	"fund/conf"
	"fund/dao/base"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
	"time"
)

type XueQiuFundDao struct {
	base.BaseDao
}

func (fundDao *XueQiuFundDao) GetFundPrice(fundCode string) decimal.Decimal {
	// https://fund.xueqiu.com/dj/open/fund/growth/008975?day=30
	// day = 30
	// day = 180
	urlString := fmt.Sprintf("https://fund.xueqiu.com/dj/open/fund/growth/%s?day=30", fundCode)
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

	respDto := &FundResp{}
	json.Unmarshal(all, respDto)

	return respDto.Data.FundNavGrowth[len(respDto.Data.FundNavGrowth)-1].Nav
}

type XueQiuStockDao struct {
	base.BaseDao
	Period string
}

func (stockDao *XueQiuStockDao) GetStockPrice(stockCode string) decimal.Decimal {
	// https://xueqiu.com/S/%s
	// 需要先请求页面获取cookie
	begin := time.Now().Add(24*time.Hour).UnixNano() / 1000 / 1000
	//begin :=
	//period := "day"
	//period:= "week"
	timeType := "before"
	count := 2

	urlString := fmt.Sprintf("https://stock.xueqiu.com/v5/stock/chart/kline.json?symbol=%s&begin=%d&period=%s&type=%s&count=%d&indicator=%s",
		stockCode,
		begin,
		stockDao.Period,
		timeType,
		-count,
		"kline,pe,pb,ps,pcf,market_capital,agt,ggt,balance",
	)
	fmt.Println(urlString)

	request, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		fmt.Println(err.Error())
		return decimal.Decimal{}
	}
	request.Header.Add("referer", fmt.Sprintf("https://xq.com/S/%s", stockCode))
	request.Header.Add("user-agent", conf.UA)

	// 获取 & 添加cookie
	cookies := GetCookies(stockCode)
	for _, v := range cookies {
		request.AddCookie(v)
	}

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

	respDto := &Resp{}
	json.Unmarshal(all, respDto)

	// 构建对象
	d, _ := respDto.Data.Items[len(respDto.Data.Items)-1].([]interface{})
	a, _ := decimal.NewFromString(fmt.Sprintf("%v", d[5]))
	return a
}
