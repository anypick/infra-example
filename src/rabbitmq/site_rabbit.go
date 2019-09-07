package rabbitmq

import (
	"fmt"
	"github.com/anypick/infra"
	"github.com/anypick/infra-rabbit/rabbitmqutil"
)

var siteRabbit *SiteRabbit

func init() {
	infra.RegisterApi(&SiteRabbit{})
}

func GetSiteRabbit() *SiteRabbit {
	return siteRabbit
}

type SiteRabbit struct {
	rabbitmqutil.RabbitOperator
}

func (r *SiteRabbit) Init() {
	siteRabbit = r
	siteRabbit.Exchange = "hmall.site.direct.exchange"
	siteRabbit.Queue = "hmall.site.direct.queue"
	siteRabbit.RoutingKey = "hmall.site.direct.queue"
	siteRabbit.MessageListener(ConsumerSiteQueue, true)
}

func ConsumerSiteQueue(data []byte) error {
	fmt.Println(string(data))
	return nil
}
