package example

import (
	"github.com/anypick/infra"
	"github.com/anypick/infra-gin"
	"github.com/anypick/infra-logrus"
	"github.com/anypick/infra/base/props/container"
)

func init() {
	infra.Register(&container.YamlStarter{})
	infra.Register(&baselog.LogrusStarter{})
	infra.Register(&basegin.GinStarter{})

	infra.Register(&infra.BaseInitializerStarter{})
}