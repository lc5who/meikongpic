package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/author"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/images"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	AuthorServiceGroup  author.ServiceGroup
	ImagesServiceGroup  images.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
