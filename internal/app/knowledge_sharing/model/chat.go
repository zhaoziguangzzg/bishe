package model

import (
	"time"
)

// Chat 定义私信结构体
type Chat struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//发送uid
	SendUid int `json:"sendUid" gorm:"column:send_uid" mapstructure:"sendUid"`
	//接收uid
	ReceiveUid int `json:"receiveUid" gorm:"column:receive_uid" mapstructure:"receiveUid"`
	//私信内容
	Content string `json:"content" gorm:"column:content" mapstructure:"content"`

	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	IsDeleted   int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	CHAT_MAX_CONTENT int = 100 //私信最长100字
)

// 指定Chat对应的表名
func (Chat) TableName() string {
	return "chat"
}
