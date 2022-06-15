package images

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/images"
	imagesReq "github.com/flipped-aurora/gin-vue-admin/server/model/images/request"
)

type ImagesService struct {
}

// CreateImages 创建Images记录
// Author [piexlmax](https://github.com/piexlmax)
func (imagesService *ImagesService) CreateImages(images images.Images) (err error) {
	err = global.GVA_DB.Create(&images).Error
	return err
}

// DeleteImages 删除Images记录
// Author [piexlmax](https://github.com/piexlmax)
func (imagesService *ImagesService) DeleteImages(images images.Images) (err error) {
	err = global.GVA_DB.Delete(&images).Error
	return err
}

// DeleteImagesByIds 批量删除Images记录
// Author [piexlmax](https://github.com/piexlmax)
func (imagesService *ImagesService) DeleteImagesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]images.Images{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateImages 更新Images记录
// Author [piexlmax](https://github.com/piexlmax)
func (imagesService *ImagesService) UpdateImages(images images.Images) (err error) {
	err = global.GVA_DB.Save(&images).Error
	return err
}

// GetImages 根据id获取Images记录
// Author [piexlmax](https://github.com/piexlmax)
func (imagesService *ImagesService) GetImages(id uint) (images images.Images, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&images).Error
	return
}

// GetImagesInfoList 分页获取Images记录
// Author [piexlmax](https://github.com/piexlmax)
func (imagesService *ImagesService) GetImagesInfoList(info imagesReq.ImagesSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&images.Images{})
	var imagess []images.Images
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&imagess).Error
	return imagess, total, err
}

func (imagesService *ImagesService) GetTopImageList(count int) (imgs []images.Images, err error) {
	db := global.GVA_DB.Model(&images.Images{})
	err = db.Order("id desc").Limit(count).Find(&imgs).Error
	return imgs, err
}
