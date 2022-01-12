package service

import (
	"encoding/json"
	"fmt"
	"fund/conf"
	"fund/dao"
	"fund/dao/xq"
	"fund/out"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
	"time"
)

type StockService struct {
	Period string
	// 数据源
	stockDao dao.StockDao
}

func NewStockService(dataSource conf.DataSource, period string) *StockService {
	return &StockService{stockDao: dao.NewStockDao(dataSource, period)}
}

func (stockService *StockService) GetPrice(stockCode string) decimal.Decimal {
	return stockService.stockDao.GetStockPrice(stockCode)
}

// GetCSVFile 雪球
func GetCSVFile(stockCode, period string) {
	// https://xueqiu.com/S/%s
	// 需要先请求页面获取cookie
	begin := time.Now().Add(24*time.Hour).UnixNano() / 1000 / 1000
	//begin :=
	//period := "day"
	//period:= "week"
	timeType := "before"
	count := 1000000

	urlString := fmt.Sprintf("https://stock.xueqiu.com/v5/stock/chart/kline.json?symbol=%s&begin=%d&period=%s&type=%s&count=%d&indicator=%s",
		stockCode,
		begin,
		period,
		timeType,
		-count,
		"kline,pe,pb,ps,pcf,market_capital,agt,ggt,balance",
	)
	fmt.Println(urlString)

	request, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Add("referer", fmt.Sprintf("https://xq.com/S/%s", stockCode))
	request.Header.Add("user-agent", conf.UA)

	// 获取 & 添加cookie
	cookies := xq.GetCookies(stockCode)
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

	respDto := &xq.Resp{}
	json.Unmarshal(all, respDto)

	out.SaveCSV(respDto.Data, period)
}
