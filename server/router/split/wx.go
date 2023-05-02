package split

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type WXRouter struct {
}

// InitWXRouter 初始化 WX 路由信息
func (s *BannerRouter) InitWXRouter(_ *gin.RouterGroup, PubRouter *gin.RouterGroup) {
	//bannerRouter := Router.Group("banner").Use(middleware.OperationRecord())
	WXPubRouter := PubRouter.Group("wx")
	var WXApi = v1.ApiGroupApp.SplitApiGroup.WXApi
	
	{
		//bannerRouterWithoutRecord.GET("findBanner", bannerApi.FindBanner) // 根据ID获取Banner
		//bannerRouterWithoutRecord.GET("getBannerList", bannerApi.GetBannerList) // 获取Banner列表
		WXPubRouter.GET("Jscode2session", WXApi.Jscode2session)
	}
}
