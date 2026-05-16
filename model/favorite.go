package model

import "time"

// Favorite 定义收藏夹结构体
type Favorite struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//收藏夹标题
	Title string `json:"title" gorm:"column:title" mapstructure:"title"`
	//用户id
	UserId int `json:"userId" gorm:"column:user_id" mapstructure:"userId"`

	CreateAt       *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr    string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt       *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr    string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	FavoriteStatus int        `json:"favoriteStatus" gorm:"column:favorite_status" mapstructure:"favoriteStatus"`
	IsDeleted      int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	FAVORITE_STATUS_NORMAL int = 0 //正常
	FAVORITE_STATUS_REVIEW int = 1 //审核

	FAVORITE_NOT_DELETED int = 0 //未被删除
	FAVORITE_IS_DELETED  int = 1 //被删除

	FAVORITE_MAX_TITLE int = 50 //收藏夹标题最长50字
)

// 指定Favorite对应的表名
func (Favorite) TableName() string {
	return "favorite"
}
