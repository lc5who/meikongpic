package author

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/author"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    authorReq "github.com/flipped-aurora/gin-vue-admin/server/model/author/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type AuthorsApi struct {
}

var authorsService = service.ServiceGroupApp.AuthorServiceGroup.AuthorsService


// CreateAuthors 创建Authors
// @Tags Authors
// @Summary 创建Authors
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body author.Authors true "创建Authors"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authors/createAuthors [post]
func (authorsApi *AuthorsApi) CreateAuthors(c *gin.Context) {
	var authors author.Authors
	_ = c.ShouldBindJSON(&authors)
	if err := authorsService.CreateAuthors(authors); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAuthors 删除Authors
// @Tags Authors
// @Summary 删除Authors
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body author.Authors true "删除Authors"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /authors/deleteAuthors [delete]
func (authorsApi *AuthorsApi) DeleteAuthors(c *gin.Context) {
	var authors author.Authors
	_ = c.ShouldBindJSON(&authors)
	if err := authorsService.DeleteAuthors(authors); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAuthorsByIds 批量删除Authors
// @Tags Authors
// @Summary 批量删除Authors
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Authors"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /authors/deleteAuthorsByIds [delete]
func (authorsApi *AuthorsApi) DeleteAuthorsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := authorsService.DeleteAuthorsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAuthors 更新Authors
// @Tags Authors
// @Summary 更新Authors
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body author.Authors true "更新Authors"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /authors/updateAuthors [put]
func (authorsApi *AuthorsApi) UpdateAuthors(c *gin.Context) {
	var authors author.Authors
	_ = c.ShouldBindJSON(&authors)
	if err := authorsService.UpdateAuthors(authors); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAuthors 用id查询Authors
// @Tags Authors
// @Summary 用id查询Authors
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query author.Authors true "用id查询Authors"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /authors/findAuthors [get]
func (authorsApi *AuthorsApi) FindAuthors(c *gin.Context) {
	var authors author.Authors
	_ = c.ShouldBindQuery(&authors)
	if reauthors, err := authorsService.GetAuthors(authors.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reauthors": reauthors}, c)
	}
}

// GetAuthorsList 分页获取Authors列表
// @Tags Authors
// @Summary 分页获取Authors列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query authorReq.AuthorsSearch true "分页获取Authors列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authors/getAuthorsList [get]
func (authorsApi *AuthorsApi) GetAuthorsList(c *gin.Context) {
	var pageInfo authorReq.AuthorsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := authorsService.GetAuthorsInfoList(pageInfo); err != nil {
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
