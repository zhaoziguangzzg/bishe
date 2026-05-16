package model

import "time"

//Contact 定义用户联系人结构体
type Contact struct {
	Id            int        `json:"id" gorm:"column:id" mapstructure:"id"`
	SendId        int        `json:"sendId" gorm:"column:send_id" mapstructure:"sendId"`
	ReceiveId     int        `json:"receiveId" gorm:"column:receive_id" mapstructure:"receiveId"`
	CreateAt      *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr   string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt      *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr   string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	ContactStatus int        `json:"contactStatus" gorm:"column:contact_status" mapstructure:"contactStatus"`
	IsDeleted     int        `json:"isDeleted" gorm:"column:is_deleted" mapstructure:"isDeleted"`
}

const (
	CONTACT_STATUS_NORMAL int = 0 //正常
	CONTACT_STATUS_REVIEW int = 1 //审核

	CONTACT_NOT_DELETED int = 0 //未被删除
	CONTACT_IS_DELETED  int = 1 //被删除
)

//指定Contact的表名
func (Contact) TableName() string {
	return "contact"
}
