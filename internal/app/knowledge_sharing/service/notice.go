package service

import (
	"bishe/internal/app/knowledge_sharing/dao/kafka"
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
	"time"
)

// 添加通知
func UserAddNotice(noticeUid int, content string, typei int, createTime time.Time) (err error) {
	return mysql.UserAddNotice(noticeUid, content, typei, createTime)
}

// 获取通知列表
func GetNoticeList(uid int, page int, pageSize int) (notices []model.Notice, err error) {
	return mysql.GetNoticeList(uid, page, pageSize)
}

// 根据类型获取通知列表
func GetNoticeListByType(uid int, typei int, page int, pageSize int) (notices []model.Notice, err error) {
	return mysql.GetNoticeListByType(uid, typei, page, pageSize)
}

// 发送notice数据到kafka
func ProduceKafkaNoticeMessage(noticeMsg *model.NoticeMsg) (partition int32, offset int64, err error) {
	return kafka.ProduceKafkaNoticeMessage(noticeMsg)
}
