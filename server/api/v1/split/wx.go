package split

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	wxRequest "github.com/flipped-aurora/gin-vue-admin/server/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WXApi struct {
}

var wxr = wxRequest.WXR

func (wx *WXApi) Jscode2session(c *gin.Context) {
	code := c.Query("code")

	res, err := wxr.Jscode2session(code)
	if err != nil {
		global.GVA_LOG.Error("获取code失败", zap.Error(err))
		response.FailWithDetailed(res, "获取code失败", c)
	} else {
		if res.Errcode != 0 {
			global.GVA_LOG.Error("获取openid失败", zap.Error(err))
			response.FailWithDetailed(res, "获取openid失败", c)
			return
		}
		response.OkWithDetailed(res, "获取成功", c)

	}
}
