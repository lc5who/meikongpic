package author

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AuthorRouter struct {
}

// InitAuthorRouter 初始化 Author 路由信息
func (s *AuthorRouter) InitAuthorRouter(Router *gin.RouterGroup) {
	authorRouter := Router.Group("author").Use(middleware.OperationRecord())
	authorRouterWithoutRecord := Router.Group("author")
	var authorApi = v1.ApiGroupApp.AuthorApiGroup.AuthorApi
	{
		authorRouter.POST("createAuthor", authorApi.CreateAuthor)   // 新建Author
		authorRouter.DELETE("deleteAuthor", authorApi.DeleteAuthor) // 删除Author
		authorRouter.DELETE("deleteAuthorByIds", authorApi.DeleteAuthorByIds) // 批量删除Author
		authorRouter.PUT("updateAuthor", authorApi.UpdateAuthor)    // 更新Author
	}
	{
		authorRouterWithoutRecord.GET("findAuthor", authorApi.FindAuthor)        // 根据ID获取Author
		authorRouterWithoutRecord.GET("getAuthorList", authorApi.GetAuthorList)  // 获取Author列表
	}
}
