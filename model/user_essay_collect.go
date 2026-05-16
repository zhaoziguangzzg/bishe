package model

import "time"

//用户对文章收藏的结构体
type UserEssayCollect struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//用户id
	UserId int `json:"userId" gorm:"column:user_id" mapstructure:"userId"`
	//文章id
	EssayId int `json:"essayId" gorm:"column:essay_id" mapstructure:"essayId"`
	//收藏夹id
	FavoriteId int `json:"favoriteId" gorm:"column:favorite_id" mapstructure:"favoriteId"`

	CreateAt      *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr   string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt      *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr   string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	CollectStatus int        `json:"collectStatus" gorm:"column:collect_status" mapstructure:"collectStatus"`
}

const (
	COLLECT_STATUS_NORMAL int = 0 //已收藏
	COLLECT_STATUS_REVIEW int = 1 //未收藏
)

// 指定UserEssayCollect对应的表名
func (UserEssayCollect) TableName() string {
	return "user_essay_collect"
}
