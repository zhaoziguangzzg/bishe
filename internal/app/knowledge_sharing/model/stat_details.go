package model

import (
	"time"
)

// StatDetails 定义统计详情结构体
type StatDetails struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//数据uid
	StatUid int `json:"statUid" gorm:"column:stat_uid" mapstructure:"statUid"`
	//类型
	Type int `json:"type" gorm:"column:type" mapstructture:"type"`

	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	IsDeleted   int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	STAT_DETAILS_TYPE_COLLECT    int = 0 //收藏
	STAT_DETAILS_TYPE_FOLLOW     int = 1 //关注
	STAT_DETAILS_TYPE_LIKE       int = 2 //点赞
	STAT_DETAILS_TYPE_COMMENT    int = 3 //评论
	STAT_DETAILS_TYPE_CIRCLE_NUM int = 4 //圈子人数
	STAT_DETAILS_TYPE_ESSAY      int = 5 //文章

	STAT_DETAILS_TYPE_FAN   int = 6 //被关注
	STAT_DETAILS_TYPE_LIKED int = 7 //被点赞

)

// 指定StatDetails对应的表名
func (StatDetails) TableName() string {
	return "stat_details"
}

// 详情结构体
type StatDetailsTypeCount struct {
	Type  int `json:"type"`
	Total int `json:"total"`
}
