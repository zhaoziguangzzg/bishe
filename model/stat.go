package model

import (
	"time"
)

// Stat 定义统计结构体
type Stat struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//统计uid
	StatUid int `json:"statUid" gorm:"column:stat_uid" mapstructure:"statUid"`
	//数据量
	Sum int `json:"sum" gorm:"column:sum" mapstructure:"sum"`
	//类型
	Type int `json:"type" gorm:"column:type" mapstructture:"type"`

	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	IsDeleted   int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	STAT_TYPE_COLLECT    int = 0 //收藏
	STAT_TYPE_FOLLOW     int = 1 //关注
	STAT_TYPE_LIKE       int = 2 //点赞
	STAT_TYPE_COMMENT    int = 3 //评论
	STAT_TYPE_CIRCLE_NUM int = 4 //圈子人数
	STAT_TYPE_ESSAY      int = 5 //文章

	STAT_TYPE_FAN   int = 6 //被关注
	STAT_TYPE_LIKED int = 7 //被点赞

)

// 指定Stat对应的表名
func (Stat) TableName() string {
	return "stat"
}
