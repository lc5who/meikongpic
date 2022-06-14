package index

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type IndexRouter struct {
}

// InitAuthorRouter 初始化 Author 路由信息
func (s *IndexRouter) InitIndexRouter(Router *gin.RouterGroup) {

	indexRouterWithoutRecord := Router.Group("index")
	var indexApi = v1.ApiGroupApp.IndexApiGroup.IndexApi
	{
		indexRouterWithoutRecord.GET("index", indexApi.Index) // 根据ID获取Author
	}
}
