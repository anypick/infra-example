package testx

import (
	"github.com/anypick/infra"
	"github.com/anypick/infra-logrus"
	"github.com/anypick/infra-mysql"
	"github.com/anypick/infra-rabbit"
	"github.com/anypick/infra/base/props"
	"github.com/anypick/infra/base/props/container"
)

// 测试之前，初始化操作
func init() {
	infra.Register(&container.YamlStarter{})
	infra.Register(&baselog.LogrusStarter{})
	infra.Register(&basesql.MySqlStarter{})
	infra.Register(&baserabbitmq.RabbitMQStarter{})

	infra.Register(&infra.BaseInitializerStarter{})
	yamlConf := props.NewYamlSource("../../../resources/application.yml")
	infra.New(*yamlConf).Start()
}
