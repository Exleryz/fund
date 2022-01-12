package xq

import (
	"encoding/json"
	"fmt"
	"fund/conf"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
)

type XueQiuFundDao struct {
	DataSource conf.DataSource
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

func (fundDao *XueQiuFundDao) GetDataSource() conf.DataSource {
	return fundDao.DataSource
}
