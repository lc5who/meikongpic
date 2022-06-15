package author

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/author"
	authorReq "github.com/flipped-aurora/gin-vue-admin/server/model/author/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AuthorsService struct {
}

// CreateAuthors 创建Authors记录
// Author [piexlmax](https://github.com/piexlmax)
func (authorsService *AuthorsService) CreateAuthors(authors author.Authors) (err error) {
	err = global.GVA_DB.Create(&authors).Error
	return err
}

// DeleteAuthors 删除Authors记录
// Author [piexlmax](https://github.com/piexlmax)
func (authorsService *AuthorsService) DeleteAuthors(authors author.Authors) (err error) {
	err = global.GVA_DB.Delete(&authors).Error
	return err
}

// DeleteAuthorsByIds 批量删除Authors记录
// Author [piexlmax](https://github.com/piexlmax)
func (authorsService *AuthorsService) DeleteAuthorsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]author.Authors{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAuthors 更新Authors记录
// Author [piexlmax](https://github.com/piexlmax)
func (authorsService *AuthorsService) UpdateAuthors(authors author.Authors) (err error) {
	err = global.GVA_DB.Save(&authors).Error
	return err
}

// GetAuthors 根据id获取Authors记录
// Author [piexlmax](https://github.com/piexlmax)
func (authorsService *AuthorsService) GetAuthors(id uint) (authors author.Authors, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&authors).Error
	return
}

// GetAuthorsInfoList 分页获取Authors记录
// Author [piexlmax](https://github.com/piexlmax)
func (authorsService *AuthorsService) GetAuthorsInfoList(info authorReq.AuthorsSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&author.Authors{})
	var authorss []author.Authors
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&authorss).Error
	return authorss, total, err
}

func (authorsService *AuthorsService) GetTopAuthorList(count int) (authors []author.Authors, err error) {
	db := global.GVA_DB.Model(&author.Authors{})
	err = db.Order("id desc").Limit(count).Find(&authors).Error
	return authors, err
}
