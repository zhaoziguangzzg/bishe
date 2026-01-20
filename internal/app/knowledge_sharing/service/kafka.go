package service

import "bishe/internal/app/knowledge_sharing/dao/kafka"

// 连接kafka
func ServiceInitKafka() (err error) {
	return kafka.DaoInitKafka()
}

// 关闭
func Closekafka() {
	kafka.Closekafka()
}
