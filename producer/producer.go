package producer

import (
	"log"
	"strings"

	"github.com/Shopify/sarama"
)

var producer sarama.AsyncProducer

// 初始化生产者
func InitProducer(hosts string) (err error) {
	config := sarama.NewConfig()
	client, err := sarama.NewClient(strings.Split(hosts, ","), config)
	if err != nil {
		log.Println("unable to create kafka client: ", err)
		return
	}
	producer, err = sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

// 发送消息
func Send(topic, data string) {
	producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(data)}
	log.Println("kafka", "Produced message: ["+data+"]")
}

func Close() {
	if producer != nil {
		producer.Close()
	}
}
