package author

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AuthorsRouter struct {
}

// InitAuthorsRouter 初始化 Authors 路由信息
func (s *AuthorsRouter) InitAuthorsRouter(Router *gin.RouterGroup) {
	authorsRouter := Router.Group("authors").Use(middleware.OperationRecord())
	authorsRouterWithoutRecord := Router.Group("authors")
	var authorsApi = v1.ApiGroupApp.AuthorApiGroup.AuthorsApi
	{
		authorsRouter.POST("createAuthors", authorsApi.CreateAuthors)   // 新建Authors
		authorsRouter.DELETE("deleteAuthors", authorsApi.DeleteAuthors) // 删除Authors
		authorsRouter.DELETE("deleteAuthorsByIds", authorsApi.DeleteAuthorsByIds) // 批量删除Authors
		authorsRouter.PUT("updateAuthors", authorsApi.UpdateAuthors)    // 更新Authors
	}
	{
		authorsRouterWithoutRecord.GET("findAuthors", authorsApi.FindAuthors)        // 根据ID获取Authors
		authorsRouterWithoutRecord.GET("getAuthorsList", authorsApi.GetAuthorsList)  // 获取Authors列表
	}
}
