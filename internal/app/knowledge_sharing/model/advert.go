package model

import "time"

// Advert 定义广告结构体
type Advert struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//广告位置
	Position string `json:"position" gorm:"column:position" mapstructure:"position"`
	//广告内容
	Content string `json:"content" gorm:"column:content" mapstructure:"content"`
	//广告开始时间
	StartTime *time.Time `json:"startTime" gorm:"column:start_time" mapstructure:"startTime"`
	//广告结束时间
	EndTime *time.Time `json:"endTime" gorm:"column:end_time" mapstructure:"endTime"`

	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	IsDeleted   int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	ADVERT_MAX_CONTENT  int = 100 //广告内容最长100字
	ADVERT_MAX_POSITION int = 100 //广告位置最长100字
)

// 指定Advert对应的表名
func (Advert) TableName() string {
	return "advert"
}
