package model

import (
	"strconv"
	"time"
)

// ChatContact 定义联系人结构体
type ChatContact struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//发送uid
	SendUid int `json:"sendUid" gorm:"column:send_uid" mapstructure:"sendUid"`
	//接收uid
	ReceiveUid int `json:"receiveUid" gorm:"column:receive_uid" mapstructure:"receiveUid"`
	//组合uid，小uid_大uid
	UnionUid string `json:"unionUid" gorm:"column:union_uid" mapstructure:"unionUid"`
	//最新消息
	Content string `json:"content" gorm:"column:content" mapstructure:"content"`

	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	//最新时间
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	IsDeleted   int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	CHAT_CONTACT_MAX_CONTENT int = 100 //私信最长100字
)

// 指定ChatContact对应的表名
func (ChatContact) TableName() string {
	return "chat_contact"
}

func MakeChatContactUnionUid(sendUid int, receiveUid int) (unionUid string) {
	if sendUid < receiveUid {
		unionUid = strconv.Itoa(sendUid) + "_" + strconv.Itoa(receiveUid)
	} else {
		unionUid = strconv.Itoa(receiveUid) + "_" + strconv.Itoa(sendUid)
	}

	return
}
