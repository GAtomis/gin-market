package split

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/split"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    splitReq "github.com/flipped-aurora/gin-vue-admin/server/model/split/request"
)

type BannerService struct {
}

// CreateBanner 创建Banner记录
// Author [piexlmax](https://github.com/piexlmax)
func (bannerService *BannerService) CreateBanner(banner *split.Banner) (err error) {
	err = global.GVA_DB.Create(banner).Error
	return err
}

// DeleteBanner 删除Banner记录
// Author [piexlmax](https://github.com/piexlmax)
func (bannerService *BannerService)DeleteBanner(banner split.Banner) (err error) {
	err = global.GVA_DB.Delete(&banner).Error
	return err
}

// DeleteBannerByIds 批量删除Banner记录
// Author [piexlmax](https://github.com/piexlmax)
func (bannerService *BannerService)DeleteBannerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]split.Banner{},"id in ?",ids.Ids).Error
	return err
}

// UpdateBanner 更新Banner记录
// Author [piexlmax](https://github.com/piexlmax)
func (bannerService *BannerService)UpdateBanner(banner split.Banner) (err error) {
	err = global.GVA_DB.Save(&banner).Error
	return err
}

// GetBanner 根据id获取Banner记录
// Author [piexlmax](https://github.com/piexlmax)
func (bannerService *BannerService)GetBanner(id uint) (banner split.Banner, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&banner).Error
	return
}

// GetBannerInfoList 分页获取Banner记录
// Author [piexlmax](https://github.com/piexlmax)
func (bannerService *BannerService)GetBannerInfoList(info splitReq.BannerSearch) (list []split.Banner, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&split.Banner{})
    var banners []split.Banner
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&banners).Error
	return  banners, total, err
}
