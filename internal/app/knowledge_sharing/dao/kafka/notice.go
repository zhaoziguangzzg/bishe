package kafka

import (
	"bishe/internal/app/knowledge_sharing/model"
	"encoding/json"

	"github.com/IBM/sarama"
)

// 发送notice数据
func ProduceKafkaNoticeMessage(noticeMsg *model.NoticeMsg) (partition int32, offset int64, err error) {
	value, err := json.Marshal(noticeMsg)
	if err != nil {
		return
	}

	msg := &sarama.ProducerMessage{
		Topic: model.KAFKA_TOPIC_NOTICE,
		Value: sarama.StringEncoder(value),
	}

	partition, offset, err = KafkaClient.SendMessage(msg)
	return
}
