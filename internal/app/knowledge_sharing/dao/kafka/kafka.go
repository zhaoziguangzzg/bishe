package kafka

import "github.com/IBM/sarama"

var KafkaClient sarama.SyncProducer

// 连接kafka
func DaoInitKafka() (err error) {
	//TODO get broker from config
	brokerlist := []string{"localhost:9092"}
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true // 成功交付的消息将在success channel返回

	KafkaClient, err = sarama.NewSyncProducer(brokerlist, config)
	if err != nil {
		return
	}
	return
}

// 关闭
func Closekafka() {

	KafkaClient.Close()

}
