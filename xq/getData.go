package xq

import (
	"bufio"
	"encoding/json"
	"fmt"
	"fund/conf"
	"fund/util"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// getCookies 获取cookie
func getCookies(stockCode string) []*http.Cookie {
	// todo 需要判断缓存失效
	if v, ok := conf.CookiesCache["xueqiu"]; ok {
		return v
	}

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

// xueqiu 雪球
func Xueqiu(stockCode, period string) {
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
	request.Header.Add("referer", "https://xq.com/S/SH513050")
	request.Header.Add("user-agent", conf.UA)

	// 获取 & 添加cookie
	cookies := getCookies(stockCode)
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

	saveCSV(respDto.Data, period)
}

func saveCSV(data RespData, period string) {
	// 查看文件路径是否存在
	if !util.IsExist("./tmp") {
		err := os.MkdirAll("./tmp/", 0777)
		if err != nil {
			util.PrintErr(err)
		}
	}
	filePath := fmt.Sprintf("./tmp/%s-%s.csv", data.Symbol, period)

	// 新建文件
	file, err := os.Create(filePath)
	if err != nil {
		util.PrintErr(err.Error())
		panic(err)
	}

	writer := bufio.NewWriter(file)
	// column 写入列头
	for i, v := range data.Column {
		writer.WriteString(v)
		if i != len(data.Column)-1 {
			writer.WriteString(",")
		} else {
			writer.WriteString("\r\n")
		}
	}

	// 写入数据
	for l, v := range data.Items {
		line := ""
		d, ok := v.([]interface{})
		if !ok {
			util.PrintErr(v)
		}

		// 构建 记录
		for index, j := range d {
			if j != nil {
				switch index {
				case 0:
					timestamp := j.(float64)
					// 转换为秒
					timestamp = timestamp / 1000
					tm := time.Unix(int64(timestamp), 0)
					line += tm.Format("2006-01-02 15:04:05")
				default:
					// j.(type) == float64
					switch t := j.(type) {
					//case int:
					//	fmt.Println("int", t)
					//case int64:
					//	fmt.Println("int64", t)
					case float64:
						if t > 10000 {
							line = fmt.Sprintf("%s%.2f", line, t)
						} else {
							//line += strconv.FormatFloat(t, 'E', -1, 64)
							// 取小数点后4位
							line = fmt.Sprintf("%s%g", line, t)
						}
					default:
						util.PrintErr("default", index, j)
					}
				}
			}

			if index != len(d)-1 {
				line += ","
			}
		}
		// 换行
		writer.WriteString(line + "\r\n")
		if l%100 == 0 {
			writer.Flush()
		}
	}
	writer.Flush()
	defer file.Close()
}
