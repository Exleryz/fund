package out

import (
	"bufio"
	"fmt"
	"fund/util"
	"fund/xq"
	"os"
	"time"
)

func SaveCSV(data xq.RespData, period string) {
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
