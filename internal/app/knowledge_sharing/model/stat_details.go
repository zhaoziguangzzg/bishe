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
	//状态
	StatStatus int `json:"statStatus" gorm:"column:stat_status" mapstructure:"statStatus"`

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

	STAT_DETAILS_STATUS_INCR int = 0 //增加数据
	STAT_DETAILS_STATUS_DECR int = 1 //减少数据

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

// 按日期统计详情结构体
type StatDetailsDateCount struct {
	Type  int    `json:"type"`
	Date  string `json:"date"`
	Total int    `json:"total"`
}
