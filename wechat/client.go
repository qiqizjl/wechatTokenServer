package wechat

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func httpGetJSON(url string, v interface{}) error {
	c := &fasthttp.Client{}
	_, body, err := c.Get(nil, url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}
