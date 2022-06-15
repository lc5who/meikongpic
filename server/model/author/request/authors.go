package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/author"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AuthorsSearch struct{
    author.Authors
    request.PageInfo
}
