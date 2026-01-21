package model

import "time"

// information 定义消息结构体
type Information struct {
	Id             int        `json:"id" gorm:"column:id" mapstructure:"id"`
	SendId         int        `json:"sendId" gorm:"column:send_id" mapstructure:"sendId"`
	ReceiveAccount int        `json:"receiveAccount" gorm:"column:receive_account" mapstructure:"receiveAccount"`
	Content        string     `json:"content" gorm:"column:content" mapstructure:"content"`
	CreateAt       *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr    string     `json:"-" gorm:"-" mapstructure:"createAt"`
}

// 指定Information对应的表名
func (Information) TableName() string {
	return "information"
}
