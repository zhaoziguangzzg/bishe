package model

import (
	"time"
)

// Notice 定义通知结构体
type Notice struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//通知uid
	NoticeUid int `json:"noticeUid" gorm:"column:notice_uid" mapstructure:"noticeUid"`
	//通知内容
	Content string `json:"content" gorm:"column:content" mapstructure:"content"`
	//类型
	Type int `json:"type" gorm:"column:type" mapstructture:"type"`

	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	IsDeleted   int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

type NoticeMsg struct {
	Type int `json:"type,omitempty"`
	Uid  int `json:"uid,omitempty"`
	//产生时间
	Time int64 `json:"time,omitempty"`
	//谁关注了uid
	UserName string `json:"userName,omitempty"`
}

const (
	NOTICE_MAX_CONTENT int = 100 //通知最长100字

	NOTICE_TYPE_OTHER      int = 0 //其他
	NOTICE_TYPE_FOLLOW     int = 1 //关注
	NOTICE_TYPE_LIKE       int = 2 //点赞
	NOTICE_TYPE_COMMENT    int = 3 //评论
	NOTICE_TYPE_ACCUSATION int = 4 //举报
	NOTICE_TYPE_FEEDBACK   int = 5 //反馈
	NOTICE_TYPE_DISPATCH   int = 6 //关注发文
	NOTICE_TYPE_ESSENCE    int = 7 //加精
	//通知topic
	KAFKA_TOPIC_NOTICE string = "topic_user_notice"
)

// 指定Notice对应的表名
func (Notice) TableName() string {
	return "notice"
}
