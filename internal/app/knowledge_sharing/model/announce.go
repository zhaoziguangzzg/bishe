package model

import "time"

// Announce 定义公告结构体
type Announce struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//公告标题
	Title string `json:"title" gorm:"column:title" mapstructure:"title"`
	//公告内容
	Content string `json:"content" gorm:"column:content" mapstructure:"content"`
	//公告开始时间
	StartTime *time.Time `json:"startTime" gorm:"column:start_time" mapstructure:"startTime"`
	//公告结束时间
	EndTime *time.Time `json:"endTime" gorm:"column:end_time" mapstructure:"endTime"`

	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	IsDeleted   int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	ANNOUNCE_MAX_CONTENT int = 200 //公告内容最长200字
	ANNOUNCE_MAX_TITLE   int = 50  //公告标题最长50字
)

// 指定Announce对应的表名
func (Announce) TableName() string {
	return "announce"
}
