package model

import "time"

//搜索结构体
type Search struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//用户id
	Uid int `json:"uid" gorm:"column:uid" mapstructure:"uid"`
	//圈子id
	Cid int `json:"cid" gorm:"column:cid" mapstructure:"cid"`
	//搜索标题
	Title string `json:"title" gorm:"column:title" mapstructure:"title"`
	//搜索次数
	SearchNum int `json:"searchNum" gorm:"column:search_num" mapstructure:"searchNum"`

	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	IsDeleted   int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

// 指定Search对应的表名
func (Search) TableName() string {
	return "search"
}
