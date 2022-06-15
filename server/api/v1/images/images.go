package images

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/images"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    imagesReq "github.com/flipped-aurora/gin-vue-admin/server/model/images/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ImagesApi struct {
}

var imagesService = service.ServiceGroupApp.ImagesServiceGroup.ImagesService


// CreateImages 创建Images
// @Tags Images
// @Summary 创建Images
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body images.Images true "创建Images"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /images/createImages [post]
func (imagesApi *ImagesApi) CreateImages(c *gin.Context) {
	var images images.Images
	_ = c.ShouldBindJSON(&images)
	if err := imagesService.CreateImages(images); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteImages 删除Images
// @Tags Images
// @Summary 删除Images
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body images.Images true "删除Images"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /images/deleteImages [delete]
func (imagesApi *ImagesApi) DeleteImages(c *gin.Context) {
	var images images.Images
	_ = c.ShouldBindJSON(&images)
	if err := imagesService.DeleteImages(images); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteImagesByIds 批量删除Images
// @Tags Images
// @Summary 批量删除Images
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Images"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /images/deleteImagesByIds [delete]
func (imagesApi *ImagesApi) DeleteImagesByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := imagesService.DeleteImagesByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateImages 更新Images
// @Tags Images
// @Summary 更新Images
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body images.Images true "更新Images"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /images/updateImages [put]
func (imagesApi *ImagesApi) UpdateImages(c *gin.Context) {
	var images images.Images
	_ = c.ShouldBindJSON(&images)
	if err := imagesService.UpdateImages(images); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindImages 用id查询Images
// @Tags Images
// @Summary 用id查询Images
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query images.Images true "用id查询Images"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /images/findImages [get]
func (imagesApi *ImagesApi) FindImages(c *gin.Context) {
	var images images.Images
	_ = c.ShouldBindQuery(&images)
	if reimages, err := imagesService.GetImages(images.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reimages": reimages}, c)
	}
}

// GetImagesList 分页获取Images列表
// @Tags Images
// @Summary 分页获取Images列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query imagesReq.ImagesSearch true "分页获取Images列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /images/getImagesList [get]
func (imagesApi *ImagesApi) GetImagesList(c *gin.Context) {
	var pageInfo imagesReq.ImagesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := imagesService.GetImagesInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
