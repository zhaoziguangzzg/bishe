package model

import "time"

// Circle 定义圈子结构体
type Circle struct {
	Id            int        `json:"id" gorm:"column:id" mapstructure:"id"`
	Title         string     `json:"title" gorm:"column:title" mapstructure:"title"`
	Price         int        `json:"price" gorm:"column:price" mapstructure:"price"`
	CircleOwnerId int        `json:"circleOwnerId" gorm:"column:circle_own_id" mapstructure:"circleOwnerId"`
	Introduction  string     `json:"introduction" gorm:"column:introduction" mapstructure:"introduction"`
	JoinNum       int        `json:"joinNum" gorm:"column:join_num" mapstructure:"joinNum"`
	CreateAt      *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr   string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt      *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr   string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	Status        int        `json:"status" gorm:"column:status" mapstructure:"status"`
}

const (
	CircleDelete int = 0 //删除
	CircleNormal int = 1 //正常
	CircleReview int = 2 //审核
)

// 指定Circle对应的表名
func (Circle) TableName() string {
	return "circle"
}
