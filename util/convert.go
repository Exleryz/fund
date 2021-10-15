package util

import "github.com/axgle/mahonia"

func DecoderConvertBytes(info []byte) string {
	return DecoderConvert(string(info))
}

func DecoderConvert(info string) string {
	return mahonia.NewDecoder("gbk").ConvertString(info)
}
