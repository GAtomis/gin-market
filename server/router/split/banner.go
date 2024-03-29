package split

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BannerRouter struct {
}

// InitBannerRouter 初始化 Banner 路由信息
func (s *BannerRouter) InitBannerRouter(Router *gin.RouterGroup, PubRouter *gin.RouterGroup) {
	bannerRouter := Router.Group("banner").Use(middleware.OperationRecord())
	bannerPubRouter := PubRouter.Group("banner")
	bannerRouterWithoutRecord := Router.Group("banner")
	var bannerApi = v1.ApiGroupApp.SplitApiGroup.BannerApi
	{
		bannerRouter.POST("createBanner", bannerApi.CreateBanner)             // 新建Banner
		bannerRouter.DELETE("deleteBanner", bannerApi.DeleteBanner)           // 删除Banner
		bannerRouter.DELETE("deleteBannerByIds", bannerApi.DeleteBannerByIds) // 批量删除Banner
		bannerRouter.PUT("updateBanner", bannerApi.UpdateBanner)              // 更新Banner
	}
	{
		bannerRouterWithoutRecord.GET("findBanner", bannerApi.FindBanner) // 根据ID获取Banner
		//bannerRouterWithoutRecord.GET("getBannerList", bannerApi.GetBannerList) // 获取Banner列表
		bannerPubRouter.GET("getBannerList", bannerApi.GetBannerList)
	}
}
