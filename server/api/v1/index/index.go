package index

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/author"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type IndexApi struct {
}

var authorService = service.ServiceGroupApp.AuthorServiceGroup.AuthorsService
var imagesService = service.ServiceGroupApp.ImagesServiceGroup.ImagesService

func (indexApi *IndexApi) Index(c *gin.Context) {

	var authors []author.Authors
	authors, err := authorService.GetTopAuthorList(10)
	imgs, err := imagesService.GetTopImageList(15)
	if err != nil {
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"authors": authors, "imagesList": imgs}, c)
	}
}

func (indexApi *IndexApi) Detail(c *gin.Context) {
	id, _ := c.GetQuery("id")
	Iid, _ := strconv.Atoi(id)
	if reimage, err := imagesService.GetImages(uint(Iid)); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		author, _ := authorService.GetAuthors(uint(*reimage.Author_id))
		response.OkWithData(gin.H{"image": reimage, "author": author}, c)
	}
}
func (indexApi *IndexApi) AuthorDetail(c *gin.Context) {
	id, _ := c.GetQuery("id")
	Iid, _ := strconv.Atoi(id)
	if reauthor, err := authorService.GetAuthors(uint(Iid)); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		reimages, _ := imagesService.GetTopImageList(0)
		response.OkWithData(gin.H{"image": reimages, "author": reauthor}, c)
	}
}

func (indexApi *IndexApi) List(c *gin.Context) {
	response.OkWithData(gin.H{"reauthor": "456"}, c)
	//var author author.Author
	//_ = c.ShouldBindQuery(&author)
	//if reauthor, err := authorService.GetAuthor(author.ID); err != nil {
	//	global.GVA_LOG.Error("查询失败!", zap.Error(err))
	//	response.FailWithMessage("查询失败", c)
	//} else {
	//	response.OkWithData(gin.H{"reauthor": reauthor}, c)
	//}
}

func (indexApi *IndexApi) AuthorSIndex(c *gin.Context) {
	result := authorService.GetAuthorSIndex(4)
	response.OkWithData(gin.H{"result": result}, c)
}
