package index

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type HomeApi struct {
}

func (indexApi *HomeApi) Index(c *gin.Context) {
	response.OkWithData(gin.H{"reauthor": "123"}, c)
}
