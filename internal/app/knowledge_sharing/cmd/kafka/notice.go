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
	"time"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

func main() {
	err := service.ServiceInit()
	if err != nil {
		panic(err)
	}
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	kafkaHandler := &consumerGroupHandler{
		ready: make(chan bool),
	}

	go func() {
		for {
			//先执行Consume，再执行setup
			if err := consumerGroup.Consume(ctx, []string{topic}, kafkaHandler); err != nil {
				service.Logger.Error("ConsumerGroup.Consume err", zap.Error(err))
			}
			if ctx.Err() != nil {
				return
			}
			kafkaHandler.ready = make(chan bool)
		}
	}()

	<-kafkaHandler.ready

	sig := <-sigChan
	service.Logger.Info("KafkaNotice get sig", zap.Any("sig", sig))
	cancel()
}

func processOneNoticeStr(msg []byte) (err error) {
	var noticeMsg model.NoticeMsg
	err = json.Unmarshal(msg, &noticeMsg)
	if err != nil {
		service.Logger.Error("Unmarshal noticeMsg err", zap.Error(err))
		return
	}

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
	case model.NOTICE_TYPE_ACCUSATION:
		noticeContent = "用户" + userName + "举报没有违规"
	case model.NOTICE_TYPE_ACCUSATIONED:
		noticeContent = "用户" + userName + "举报有违规"
	case model.NOTICE_TYPE_FEEDBACK:
		noticeContent = "用户" + userName + "反馈了"
	default:
		noticeContent = "又有通知了"
	}

	fmt.Println(noticeContent)
	nowTime := time.Now()

	err = service.UserAddNotice(noticeMsg.Uid, noticeContent, noticeMsg.Type, nowTime)
	if err != nil {
		service.Logger.Error("UserAddNotice err", zap.Error(err))
		return
	}

	service.Logger.Info("notice msg", zap.Any("noticeMsg", noticeMsg))
	return
}

type consumerGroupHandler struct {
	ready chan bool
}

func (h *consumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	close(h.ready)
	return nil
}

func (h *consumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				return nil
			}
			processOneNoticeStr(message.Value)
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
