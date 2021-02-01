package conf

import (
	"fund/util"
	"github.com/go-ini/ini"
	"net/http"
)

var Config *ini.File
var UA string
var CookiesCache = make(map[string][]*http.Cookie)

func init() {
	var err error
	Config, err = ini.Load("conf.ini")
	if err != nil {
		util.PrintErr(err.Error())
		return
	}
	LoadUA()
}

func LoadUA() {
	UA = Config.Section("").Key("UA").MustString("Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36")
}
