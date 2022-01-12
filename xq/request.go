package xq

import (
	"fmt"
	"fund/conf"
	"fund/util"
	"net/http"
)

// GetCookies 获取cookie
func GetCookies(stockCode string) []*http.Cookie {
	// todo 需要判断缓存失效
	if v, ok := conf.CookiesCache["xueqiu"]; ok {
		return v
	}

	// F 基金 S 股票/场内基金
	urlStr := fmt.Sprintf("https://xueqiu.com/S/%s", stockCode)
	request, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		util.PrintErr(err.Error())
		panic(err)
	}
	request.Header.Add("user-agent", conf.UA)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		util.PrintErr(err.Error())
		panic(err)
	}
	defer resp.Body.Close()

	cookies := resp.Cookies()
	for i, v := range cookies {
		util.PrintlnInfo(i, v)
	}

	// 添加缓存
	conf.CookiesCache["xueqiu"] = resp.Cookies()

	return conf.CookiesCache["xueqiu"]
}
