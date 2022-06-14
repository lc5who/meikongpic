// 自动生成模板Author
package author

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Author 结构体
// 如果含有time.Time 请自行import time包
type Author struct {
      global.GVA_MODEL
      Nickname  string `json:"nickname" form:"nickname" gorm:"column:nickname;comment:;size:50;"`
      Avatar  string `json:"avatar" form:"avatar" gorm:"column:avatar;comment:;size:255;"`
      Credit  *float64 `json:"credit" form:"credit" gorm:"column:credit;comment:;size:10,2;"`
}


// TableName Author 表名
func (Author) TableName() string {
  return "author"
}

