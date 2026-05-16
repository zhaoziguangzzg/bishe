package main

import (
	"bishe/model"
	"bishe/service"
	"context"
	"os"
	"os/signal"
	"syscall"

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

	kafkaHandler := &service.ConsumerGroupHandler{
		Ready:       make(chan bool),
		ProcessFunc: service.ProcessKafkaNotice,
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
			kafkaHandler.Ready = make(chan bool)
		}
	}()

	<-kafkaHandler.Ready

	sig := <-sigChan
	service.Logger.Info("KafkaNotice get sig", zap.Any("sig", sig))
	cancel()
}
