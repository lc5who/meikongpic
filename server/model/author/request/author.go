package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/author"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AuthorSearch struct{
    author.Author
    request.PageInfo
}
