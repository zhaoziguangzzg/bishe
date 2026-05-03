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
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	brokers := []string{"localhost:9092"}
	topic := model.KAFKA_TOPIC_NOTICE
	groupID := "group_notice"
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		service.Logger.Error("NewConsumerGroup err", zap.Error(err))
		return
	}
	defer func() {
		err := consumerGroup.Close()
		if err != nil {
			service.Logger.Error("consumerGroup.Close err", zap.Error(err))
		}
	}()
	ctx := context.Background()

	go func() {
		for {
			err := consumerGroup.Consume(ctx, []string{topic}, &consumerGroupHandler{})
			if err != nil {
				service.Logger.Error("ConsumerGroup.Consume err", zap.Error(err))
			}
		}
	}()

	go func() {
	notice:
		for {
			select {
			case sig := <-sigChan:
				service.Logger.Info("KafkaNotice get sig", zap.Any("sig", sig))
				break notice
			case msg := <-noticeChan:
				var noticeMsg model.NoticeMsg
				err := json.Unmarshal(msg, &noticeMsg)
				if err != nil {
					service.Logger.Error("Unmarshal noticeMsg err", zap.Error(err))
					continue
				}

				//判断type，组合通知内容
				userName := noticeMsg.UserName
				var noticeContent string
				switch noticeMsg.Type {
				case model.NOTICE_TYPE_FOLLOW:
					noticeContent = "又有用户" + userName + "关注啦"
				case model.NOTICE_TYPE_LIKE:
					noticeContent = "又有用户" + userName + "点赞啦"
				case model.NOTICE_TYPE_COMMENT:
					noticeContent = "又有用户" + userName + "评论啦"
				case model.NOTICE_TYPE_DISPATCH:
					noticeContent = "又有用户" + userName + "关注发文啦"
				case model.NOTICE_TYPE_JOIN:
					noticeContent = "又有用户" + userName + "加入圈子啦"
				default:
					noticeContent = "又有通知了"
				}
				fmt.Println(noticeContent)

				service.Logger.Info("notice msg", zap.Any("noticeMsg", noticeMsg))
			default:
			}
		}
	}()
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
