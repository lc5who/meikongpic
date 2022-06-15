package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/author"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/images"
	"github.com/flipped-aurora/gin-vue-admin/server/router/index"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Author  author.RouterGroup
	Index   index.RouterGroup
	Images  images.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
