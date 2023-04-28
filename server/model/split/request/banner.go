package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/split"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BannerSearch struct{
    split.Banner
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
