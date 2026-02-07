package model

import "time"

// information 定义消息结构体
type Information struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//发送uid
	SendId int `json:"sendId" gorm:"column:send_id" mapstructure:"sendId"`
	//接收name
	ReceiveId int `json:"receiveId" gorm:"column:receive_id" mapstructure:"receiveId"` //TODO 修改数据库
	//消息内容
	Content string `json:"content" gorm:"column:content" mapstructure:"content"`

	CreateAt          *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr       string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt          *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr       string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	InformationStatus int        `json:"informationStatus" gorm:"column:information_status" mapstructure:"informationStatus"`
	IsDeleted         int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	INFORMATION_STATUS_NORMAL int = 0 //正常
	INFORMATION_STATUS_REVIEW int = 1 //审核

	INFORMATION_NOT_DELETED int = 0 //未被删除
	INFORMATION_IS_DELETED  int = 1 //被删除

	INFORMATION_MAX_RECEIVE_NAME int = 20  //消息接收者名最长20字
	INFORMATION_MAX_CONTENT      int = 100 //消息最长100字
)

// 指定Information对应的表名
func (Information) TableName() string {
	return "information"
}
