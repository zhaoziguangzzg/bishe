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
	//跳转地址
	Url string `json:"url" gorm:"column:url" mapstructture:"url"`

	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	IsDeleted   int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

type NoticeMsg struct {
	Type int `json:"type,omitempty"`
	//产生时间
	Time int64 `json:"time,omitempty"`
	//作者id
	AuthorUid int `json:"authorUid,omitempty"`
	//文章id
	EssayId int `json:"essayId,omitempty"`
	//圈子id
	CircleId int `json:"circleId,omitempty"`
	//A关注了B中的A
	FanUid int `json:"fanUid,omitempty"`
	//A关注了B中的A
	FollowUid int `json:"followUid,omitempty"`
	//点赞的用户id
	LikeUid int `json:"likeUid,omitempty"`
	//评论的用户id
	CommentUid int `json:"commentUid,omitempty"`
	//加入圈子的用户id
	JoinUid int `json:"joinUid,omitempty"`
	//举报的用户id
	AccusationUid int `json:"accusationUid,omitempty"`
	//举报的id
	AccusationId int `json:"accusationId,omitempty"`
	//反馈的用户id
	FeedbackUid int `json:"feedbackUid,omitempty"`
	//反馈的id
	FeedbackId int `json:"feedbackId,omitempty"`
}

const (
	NOTICE_MAX_CONTENT int = 100 //通知最长100字

	NOTICE_TYPE_OTHER        int = 0 //其他
	NOTICE_TYPE_FOLLOW       int = 1 //关注
	NOTICE_TYPE_LIKE         int = 2 //点赞
	NOTICE_TYPE_COMMENT      int = 3 //评论
	NOTICE_TYPE_ACCUSATION   int = 4 //举报
	NOTICE_TYPE_FEEDBACK     int = 5 //反馈
	NOTICE_TYPE_ESSAY_ADD    int = 6 //发布文章通知关注者
	NOTICE_TYPE_ESSENCE      int = 7 //加精
	NOTICE_TYPE_JOIN         int = 8 //加入圈子
	NOTICE_TYPE_ACCUSATIONED int = 9 //被举报

	//通知topic
	KAFKA_TOPIC_NOTICE string = "topic_user_notice"
)

type NoticeTypeName struct {
	Type int    `json:"type"`
	Name string `json:"name"`
}

// 指定Notice对应的表名
func (Notice) TableName() string {
	return "notice"
}
