package request

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/model/split/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type wxR struct {
}

var WXR = new(wxR)

func (w *wxR) Jscode2session(code string) (*response.WxRes, error) {
	m := make(map[string]string)
	m["appid"] = "wxe80ce993af853788"
	m["secret"] = "81061b03747fc107f05dbf684417a4f4"
	m["js_code"] = code
	m["grant_type"] = "authorization_code"
	queryStr := utils.GetParams(m)
	resb, err := utils.RequestUrlByGet(queryStr, "https://api.weixin.qq.com/sns/jscode2session")

	var wxr response.WxRes
	err = json.Unmarshal(resb, &wxr)
	if err != nil {
		return nil, err
	}
	return &wxr, nil

}
