package xq

type Resp struct {
	Data             RespData `json:"data"`
	ErrorCode        int      `json:"error_code"`
	ErrorDescription string   `json:"error_description"`
}

type RespData struct {
	Symbol string        `json:"symbol"`
	Column []string      `json:"column"`
	Items  []interface{} `json:"item"`
}
