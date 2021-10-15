package sina

import (
	"fmt"
	"fund/util"
	"io"
	"net/http"
	"os"
	"strings"
)

type SinaData struct {
}

func (sina *SinaData) GetData(stockCode []string) {

	stockCodeList := ""
	for _, v := range stockCode {
		// daily
		// min
		// weekly
		// monthly
		imageUrl := fmt.Sprintf("https://image.sinajs.cn/newchart/daily/n/%s.gif", v)
		imgResp, err := http.Get(imageUrl)
		if err != nil {
			util.PrintErr(err.Error())
		}
		defer imgResp.Body.Close()

		gifFile, err := os.OpenFile(fmt.Sprintf("%s.gif", v), os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			util.PrintErr(err)
		}
		defer gifFile.Close()

		gifAll, _ := io.ReadAll(imgResp.Body)
		gifFile.Write(gifAll)

		stockCodeList = fmt.Sprintf("%s,%s", stockCodeList, v)
	}
	stockCodeList = stockCodeList[1:]

	url := "https://hq.sinajs.cn/list=" + stockCodeList

	response, err := http.Get(url)
	if err != nil {
		util.PrintErr(err.Error())
	}
	defer response.Body.Close()

	resp, _ := io.ReadAll(response.Body)
	stockList := strings.Split(util.DecoderConvertBytes(resp), "\n")
	fmt.Println(stockList)

	for _, v := range stockList {
		if len(v) == 0 {
			continue
		}

		// 去掉最后的 ";
		stockInfo := v[strings.Index(v, "\"")+1 : len(v)-2]

		split := strings.Split(stockInfo, ",")
		if len(split) == 0 {
			continue
		}
		// 股票名字
		// 今日开盘价
		// 昨日收盘价
		// 当前价格
		// 今日最高价
		// 今日最低价
		// 竞买价，即“买一”报价
		// 竞卖价，即“卖一”报价
		// 成交的股票数，由于股票交易以一百股为基本单位，所以在使用时，通常把该值除以一百；
		// 成交金额，单位为“元”，为了一目了然，通常以“万元”为成交金额的单位，所以通常把该值除以一万；
		// “买一”申请4695股，即47手；
		// “买一”报价；
		// “买二”
		// “买二”
		// “买三”
		// “买三”
		// “买四”
		// “买四”
		// “买五”
		// “买五”
		// 卖一”申报3100股，即31手
		// “卖一”报价
		// (22, 23), (24, 25), (26,27), (28, 29)分别为“卖二”至“卖四的情况”
		// 日期
		// 时间

		//stockName := split[0]
		//openPrice := split[1]
		//yesterdayClosePrice := split[2]
		//currentPrice := split[3]

		util.PrintlnInfo(fmt.Sprintf(
			"股票名称: %s, "+
				"今日开盘价: %s, "+
				"昨日收盘价: %s, "+
				"当前价格: %s, "+
				"今日最高价: %s, "+
				"今日最低价: %s, "+
				"竞买价: %s, "+
				"竞卖价: %s, "+
				"成交的股票数: %s, "+
				"成交金额: %s, "+
				"买一: %s, "+
				"买一: %s, "+
				"买二: %s, "+
				"买二: %s, "+
				"买三: %s, "+
				"买三: %s, "+
				"买四: %s, "+
				"买四: %s, "+
				"买五: %s, "+
				"买五: %s, "+
				"卖一: %s, "+
				"卖一: %s, "+
				"卖二: %s, "+
				"卖二: %s, "+
				"卖三: %s, "+
				"卖三: %s, "+
				"卖四: %s, "+
				"卖四: %s, "+
				"卖五: %s, "+
				"卖五: %s, "+
				"日期: %s, "+
				"时间: %s, "+
				"",
			split[0],
			split[1],
			split[2],
			split[3],
			split[4],
			split[5],
			split[6],
			split[7],
			split[8],
			split[9],
			split[10],
			split[11],
			split[12],
			split[13],
			split[14],
			split[15],
			split[16],
			split[17],
			split[18],
			split[19],
			split[20],
			split[21],
			split[22],
			split[23],
			split[24],
			split[25],
			split[26],
			split[27],
			split[28],
			split[29],
			split[30],
			split[31],
		))
		// 32 可能会有 33
	}

}
