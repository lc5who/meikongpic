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

type AuthorApi struct {
}

var authorService = service.ServiceGroupApp.AuthorServiceGroup.AuthorService


// CreateAuthor 创建Author
// @Tags Author
// @Summary 创建Author
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body author.Author true "创建Author"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /author/createAuthor [post]
func (authorApi *AuthorApi) CreateAuthor(c *gin.Context) {
	var author author.Author
	_ = c.ShouldBindJSON(&author)
	if err := authorService.CreateAuthor(author); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAuthor 删除Author
// @Tags Author
// @Summary 删除Author
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body author.Author true "删除Author"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /author/deleteAuthor [delete]
func (authorApi *AuthorApi) DeleteAuthor(c *gin.Context) {
	var author author.Author
	_ = c.ShouldBindJSON(&author)
	if err := authorService.DeleteAuthor(author); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAuthorByIds 批量删除Author
// @Tags Author
// @Summary 批量删除Author
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Author"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /author/deleteAuthorByIds [delete]
func (authorApi *AuthorApi) DeleteAuthorByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := authorService.DeleteAuthorByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAuthor 更新Author
// @Tags Author
// @Summary 更新Author
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body author.Author true "更新Author"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /author/updateAuthor [put]
func (authorApi *AuthorApi) UpdateAuthor(c *gin.Context) {
	var author author.Author
	_ = c.ShouldBindJSON(&author)
	if err := authorService.UpdateAuthor(author); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAuthor 用id查询Author
// @Tags Author
// @Summary 用id查询Author
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query author.Author true "用id查询Author"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /author/findAuthor [get]
func (authorApi *AuthorApi) FindAuthor(c *gin.Context) {
	var author author.Author
	_ = c.ShouldBindQuery(&author)
	if reauthor, err := authorService.GetAuthor(author.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reauthor": reauthor}, c)
	}
}

// GetAuthorList 分页获取Author列表
// @Tags Author
// @Summary 分页获取Author列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query authorReq.AuthorSearch true "分页获取Author列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /author/getAuthorList [get]
func (authorApi *AuthorApi) GetAuthorList(c *gin.Context) {
	var pageInfo authorReq.AuthorSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := authorService.GetAuthorInfoList(pageInfo); err != nil {
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
