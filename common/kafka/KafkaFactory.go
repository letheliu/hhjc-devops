package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	logTraceService "github.com/letheliu/hhjc-devops/business/service/logTrace"
	config2 "github.com/letheliu/hhjc-devops/config"
)

const (
	TopicLogTrace string = "logTrace"
)

var (
	kafkaConsumer KafkaConsumer
	kafkaProducer KafkaProducer
)

func Init() {
	if config2.G_AppConfig.KafkaSwitch != config2.Kafka_switch_on {
		return
	}
	// init consumer
	kafkaConsumer.Init()

	// start linsten
	kafkaConsumer.StartKafkaListen(kafkaListen)

	// init producer
	kafkaProducer.Init()

}

// kafka listen
func kafkaListen(message *sarama.ConsumerMessage) {
	fmt.Println(message.Topic, "收到卡夫卡消息", string(message.Value))
	if message.Topic == TopicLogTrace {
		saveLogTrace(message)
	}
}

// save log
func saveLogTrace(message *sarama.ConsumerMessage) {

	var (
		logTraceService logTraceService.LogTraceService
	)

	logTraceService.SaveLogTraces(string(message.Value))

}

// send message
func SendMessage(topic string, data string) {
	kafkaProducer.SendMessage(topic, []byte(data))
}
