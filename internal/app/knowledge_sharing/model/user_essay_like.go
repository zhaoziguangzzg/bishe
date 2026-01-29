package model

import "time"

//用户对文章点赞的结构体
type UserEssayLike struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//用户id
	UserId int `json:"userId" gorm:"column:user_id" mapstructure:"userId"`
	//文章id
	EssayId int `json:"essayId" gorm:"column:essay_id" mapstructure:"essayId"`

	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	LikeStatus  int        `json:"likeStatus" gorm:"column:like_status" mapstructure:"likeStatus"`
	IsDeleted   int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	LIKE_STATUS_NORMAL int = 0 //正常
	LIKE_STATUS_REVIEW int = 1 //审核

	LIKE_NOT_DELETED int = 0 //未被删除
	LIKE_IS_DELETED  int = 1 //被删除
)

// 指定UserEssayLike对应的表名
func (UserEssayLike) TableName() string {
	return "user_essay_Like"
}
