package utils

import (
	"io"
	"net/http"
	"net/url"
)

func GetParams(params map[string]string) string {
	queryUrl := url.Values{}
	for key, value := range params {
		queryUrl.Set(key, value)
	}
	p := queryUrl.Encode()
	return p

}

func RequestUrlByGet(p string, domain string) ([]byte, error) {

	cli := http.Client{}
	req, err := http.NewRequest(http.MethodGet, domain+"?"+p, nil)
	if err != nil {
		return nil, err
	}
	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	resb, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resb, nil

}
