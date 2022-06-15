// 自动生成模板Authors
package author

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Authors 结构体
// 如果含有time.Time 请自行import time包
type Authors struct {
      global.GVA_MODEL
      Nickname  string `json:"nickname" form:"nickname" gorm:"column:nickname;comment:昵称;size:50;"`
      Info  string `json:"info" form:"info" gorm:"column:info;comment:;size:255;"`
      Avatar  string `json:"avatar" form:"avatar" gorm:"column:avatar;comment:;size:255;"`
      Credit  *float64 `json:"credit" form:"credit" gorm:"column:credit;comment:积分;size:10,2;"`
      Popular  *int `json:"popular" form:"popular" gorm:"column:popular;comment:查看次数;size:11;"`
}


// TableName Authors 表名
func (Authors) TableName() string {
  return "authors"
}

