package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/images"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ImagesSearch struct{
    images.Images
    request.PageInfo
}
