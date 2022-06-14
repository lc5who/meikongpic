package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/author"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/index"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	AuthorApiGroup  author.ApiGroup
	IndexApiGroup   index.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
