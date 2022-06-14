package author

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/author"
	authorReq "github.com/flipped-aurora/gin-vue-admin/server/model/author/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AuthorService struct {
}

// CreateAuthor 创建Author记录
// Author [piexlmax](https://github.com/piexlmax)
func (authorService *AuthorService) CreateAuthor(author author.Author) (err error) {
	err = global.GVA_DB.Create(&author).Error
	return err
}

// DeleteAuthor 删除Author记录
// Author [piexlmax](https://github.com/piexlmax)
func (authorService *AuthorService) DeleteAuthor(author author.Author) (err error) {
	err = global.GVA_DB.Delete(&author).Error
	return err
}

// DeleteAuthorByIds 批量删除Author记录
// Author [piexlmax](https://github.com/piexlmax)
func (authorService *AuthorService) DeleteAuthorByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]author.Author{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAuthor 更新Author记录
// Author [piexlmax](https://github.com/piexlmax)
func (authorService *AuthorService) UpdateAuthor(author author.Author) (err error) {
	err = global.GVA_DB.Save(&author).Error
	return err
}

// GetAuthor 根据id获取Author记录
// Author [piexlmax](https://github.com/piexlmax)
func (authorService *AuthorService) GetAuthor(id uint) (author author.Author, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&author).Error
	return
}

// GetAuthorInfoList 分页获取Author记录
// Author [piexlmax](https://github.com/piexlmax)
func (authorService *AuthorService) GetAuthorInfoList(info authorReq.AuthorSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&author.Author{})
	var authors []author.Author
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Nickname != "" {
		db = db.Where("nickname = ?", info.Nickname)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&authors).Error
	return authors, total, err
}

func (authorService *AuthorService) GetTopAuthorList(count int) (authors []author.Author, err error) {
	db := global.GVA_DB.Model(&author.Author{})
	err = db.Order("id desc").Limit(count).Find(&authors).Error
	return authors, err
}
