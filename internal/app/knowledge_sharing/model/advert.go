package model

import "time"

// Advert 定义文章结构体
type Advert struct {
	Id          int        `json:"id" gorm:"column:id" mapstructure:"id"`
	Position    string     `json:"position" gorm:"column:position" mapstructure:"position"`
	Content     string     `json:"content" gorm:"column:content" mapstructure:"content"`
	CreatorId   int        `json:"creatorId" gorm:"column:creator_id" mapstructure:"creatorId"`
	StartTime   *time.Time `json:"startTime" gorm:"column:start_time" mapstructure:"startTime"`
	EndTime     *time.Time `json:"endTime" gorm:"column:end_time" mapstructure:"endTime"`
	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	Status      int        `json:"essayStatus" gorm:"column:circle_status" mapstructure:"circleStatus"`
}

const (
	AdvertDelete int = 0 //删除
	AdvertNormal int = 1 //正常
)

// 指定Advert对应的表名
func (Advert) TableName() string {
	return "advert"
}
