package service

import (
	"bishe/dao/kafka"

	"github.com/IBM/sarama"
)

// 连接kafka
func ServiceInitKafka() (err error) {
	return kafka.DaoInitKafka()
}

// 关闭
func Closekafka() {
	kafka.Closekafka()
}

type ConsumerGroupHandler struct {
	Ready       chan bool
	ProcessFunc func([]byte) error
}

func (h *ConsumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	close(h.Ready)
	return nil
}

func (h *ConsumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				return nil
			}
			h.ProcessFunc(message.Value)
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
