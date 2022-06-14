package index

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

type IndexApi struct {
}

var authorService = service.ServiceGroupApp.AuthorServiceGroup.AuthorService

func (indexApi *IndexApi) Index(c *gin.Context) {
	response.OkWithData(gin.H{"reauthor": "123"}, c)
	//var author author.Author
	//_ = c.ShouldBindQuery(&author)
	//if reauthor, err := authorService.GetAuthor(author.ID); err != nil {
	//	global.GVA_LOG.Error("查询失败!", zap.Error(err))
	//	response.FailWithMessage("查询失败", c)
	//} else {
	//	response.OkWithData(gin.H{"reauthor": reauthor}, c)
	//}
}
