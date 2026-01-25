package model

import "time"

// Circle 定义圈子结构体
type Circle struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//圈子标题
	Title string `json:"title" gorm:"column:title" mapstructure:"title"`
	//圈子价格
	Price int `json:"price" gorm:"column:price" mapstructure:"price"`
	//圈主id
	CircleOwnerId int `json:"circleOwnerId" gorm:"column:circle_own_id" mapstructure:"circleOwnerId"`
	//圈子简介
	Introduction string `json:"introduction" gorm:"column:introduction" mapstructure:"introduction"`
	//加入圈子人数
	JoinNum int `json:"joinNum" gorm:"column:join_num" mapstructure:"joinNum"`

	CreateAt     *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr  string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt     *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr  string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	CircleStatus int        `json:"circleStatus" gorm:"column:circle_status" mapstructure:"circleStatus"`
	IsDeleted    int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	CIRCLE_STATUS_NORMAL int = 0 //正常
	CIRCLE_STATUS_REVIEW int = 1 //审核

	CIRCLE_NOT_DELETED int = 0 //未被删除
	CIRCLE_IS_DELETED  int = 1 //被删除

	CIRCLE_MAX_TITLE        int = 100   //圈子标题最长100字
	CIRCLE_MAX_INTRODUCTION int = 100   //圈子简介最长100字
	CIRCLE_MAX_PRICE        int = 10000 //圈子价格最大1w
)

// 指定Circle对应的表名
func (Circle) TableName() string {
	return "circle"
}
