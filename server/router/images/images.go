package images

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ImagesRouter struct {
}

// InitImagesRouter 初始化 Images 路由信息
func (s *ImagesRouter) InitImagesRouter(Router *gin.RouterGroup) {
	imagesRouter := Router.Group("images").Use(middleware.OperationRecord())
	imagesRouterWithoutRecord := Router.Group("images")
	var imagesApi = v1.ApiGroupApp.ImagesApiGroup.ImagesApi
	{
		imagesRouter.POST("createImages", imagesApi.CreateImages)   // 新建Images
		imagesRouter.DELETE("deleteImages", imagesApi.DeleteImages) // 删除Images
		imagesRouter.DELETE("deleteImagesByIds", imagesApi.DeleteImagesByIds) // 批量删除Images
		imagesRouter.PUT("updateImages", imagesApi.UpdateImages)    // 更新Images
	}
	{
		imagesRouterWithoutRecord.GET("findImages", imagesApi.FindImages)        // 根据ID获取Images
		imagesRouterWithoutRecord.GET("getImagesList", imagesApi.GetImagesList)  // 获取Images列表
	}
}
