package tool

import (
	"encoding/json"
	"io"
)

type JsonParse struct {
}

// 解析io流，并转化为json格式
func Decode(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(v)
}
