package test

import (
	"encoding/json"
	"example/src/rabbitmq"
	_ "example/testx"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"testing"
)

func TestConsumerSiteQueue(t *testing.T) {
	rabbit := rabbitmq.GetSiteRabbit()
	data := make(map[string]interface{})
	data["name"] = "张三"
	data["age"] = 20
	dataByte, _ := json.Marshal(data)
	err := rabbit.Send(amqp.Publishing{Body: dataByte, ContentType: "application/json"})
	if err != nil {
		logrus.Error(err)
	}
}
