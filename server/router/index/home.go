package index

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type HomeRouter struct {
}

// InitAuthorRouter 初始化 Author 路由信息
func (s *HomeRouter) InitHomeRouter(Router *gin.RouterGroup) {

	indexRouterWithoutRecord := Router.Group("home")
	var homeApi = v1.ApiGroupApp.IndexApiGroup.HomeApi
	{
		indexRouterWithoutRecord.GET("index", homeApi.Index) // 根据ID获取Author
	}

}
