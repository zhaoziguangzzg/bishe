package model

import "time"

// Essay 定义文章结构体
type Essay struct {
	Id          int        `json:"id" gorm:"column:id" mapstructure:"id"`
	Title       string     `json:"title" gorm:"column:title" mapstructure:"title"`
	CircleId    int        `json:"circleId" gorm:"column:circleId" mapstructure:"circleId"`
	Content     string     `json:"content" gorm:"column:content" mapstructure:"content"`
	AuthorId    int        `json:"authorId" gorm:"column:author_id" mapstructure:"authorId"`
	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	EssayStatus int        `json:"essayStatus" gorm:"column:circle_status" mapstructure:"circleStatus"`
}

const (
	EssayDelete int = 0 //删除
	EssayNormal int = 1 //正常
	EssayReview int = 2 //审核
)

// 指定Essay对应的表名
func (Essay) TableName() string {
	return "essay"
}
