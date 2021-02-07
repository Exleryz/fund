package jsl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// jisilu 集思录
func jisilu() {
	urlStr := fmt.Sprintf("https://www.jisilu.cn/data/etf/detail_hists/?___jsl=LST___t=%d", time.Now().UnixNano()/1000/1000)
	request, err := http.NewRequest("POST", urlStr, strings.NewReader("is_search=1&fund_id=159801&rp=50&page=1"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
	fmt.Println(string(all))
}
