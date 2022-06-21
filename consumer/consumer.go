package consumer

import (
	"log"
	"strings"

	"github.com/Shopify/sarama"
)

var consumer sarama.Consumer

// 消费者回调函数
type ConsumerCallback func(data []byte)

// 初始化消费者
func InitConsumer(hosts string) (err error) {
	config := sarama.NewConfig()
	client, err := sarama.NewClient(strings.Split(hosts, ","), config)
	if err != nil {
		log.Println("unable to create kafka client: ", err)
		return
	}

	consumer, err = sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

// 消费者循环
func LoopConsumer(topic string, callback ConsumerCallback) (err error) {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Println(err)
		return
	}
	defer partitionConsumer.Close()

	for {
		msg := <-partitionConsumer.Messages()
		if callback != nil {
			callback(msg.Value)
		}
	}
}

func Close() {
	if consumer != nil {
		consumer.Close()
	}
}
