// 自动生成模板Images
package images

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Images 结构体
// 如果含有time.Time 请自行import time包
type Images struct {
      global.GVA_MODEL
      Url  string `json:"url" form:"url" gorm:"column:url;comment:图片地址;size:255;"`
      Author_id  *int `json:"author_id" form:"author_id" gorm:"column:author_id;comment:作者ID;size:11;"`
      View_count  *int `json:"view_count" form:"view_count" gorm:"column:view_count;comment:查看次数;size:11;"`
      Download_count  *int `json:"download_count" form:"download_count" gorm:"column:download_count;comment:下载次数;size:11;"`
}


// TableName Images 表名
func (Images) TableName() string {
  return "images"
}

