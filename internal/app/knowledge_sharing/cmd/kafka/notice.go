package main

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

var noticeChan = make(chan []byte, 10)

//消费kafka，发notice

func main() {
	err := service.ServiceInit()
	if err != nil {
		panic(err)
	}
	//main结束之前将日志写到文件
	defer service.SyncLogger()
	defer service.Closekafka()

	Notice()
}

func Notice() {
	sigChan := make(chan os.Signal, 1)
	//windows收不到命令，只能收到 Ctrl+C / Ctrl+Break,不是sleep，是收不到信号
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	brokers := []string{"localhost:9092"}
	topic := model.KAFKA_TOPIC_NOTICE
	groupID := "group_notice"
	config := sarama.NewConfig()
	// config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		service.Logger.Error("NewConsumerGroup err", zap.Error(err))
	}
	defer func() {
		err := consumerGroup.Close()
		if err != nil {
			service.Logger.Error("consumerGroup.Close err", zap.Error(err))
		}
	}()
	ctx := context.Background()

notice:
	for msg := range noticeChan {
		select {
		case sig := <-sigChan:
			//当收到信号时，记录日志，结束循环
			service.Logger.Info("KafkaNotice get sig", zap.Any("sig", sig))
			break notice
		default:
			//没有信号，继续for循环
		}

		var noticeMsg model.NoticeMsg
		err := json.Unmarshal(msg, &noticeMsg)
		if err != nil {
			service.Logger.Error("Unmarshal noticeMsg err", zap.Error(err))
			continue
		}

		userName := noticeMsg.UserName
		content := "又有新用户" + userName + "关注啦"
		fmt.Println(content)
		// //添加通知
		// err = service.UserAddNotice(noticeMsg.Uid, content, noticeMsg.Type, time.Unix(noticeMsg.Time, 0))
		// if err != nil {
		// 	service.Logger.Error("UserAddNotice err", zap.Error(err))
		// 	return
		// }

		service.Logger.Info("v", zap.Any("v value", noticeMsg))
		err = consumerGroup.Consume(ctx, []string{topic}, &consumerGroupHandler{})
		if err != nil {
			service.Logger.Error("ConsumerGroup.Consume err", zap.Error(err))
		}

	}

	close(noticeChan)

}

type consumerGroupHandler struct{}

// Setup 在每个会话开始前调用
func (h consumerGroupHandler) Setup(session sarama.ConsumerGroupSession) (err error) {
	return nil
}

// Cleanup 在每个会话结束后调用
func (h consumerGroupHandler) Cleanup(claim sarama.ConsumerGroupSession) (err error) {
	return nil
}

// ConsumeClaim 处理分配给该消费者的分区中的消息
func (h consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) (err error) {

	for message := range claim.Messages() {

		noticeChan <- message.Value
		session.MarkMessage(message, "")
	}

	return nil
}
