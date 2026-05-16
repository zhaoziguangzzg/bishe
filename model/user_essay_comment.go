package model

import "time"

//用户评论文章的结构体
type UserEssayComment struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//评论用户id
	UserId int `json:"userId" gorm:"column:user_id" mapstructure:"userId"`
	//文章id
	EssayId int `json:"essayId" gorm:"column:essay_id" mapstructure:"essayId"`
	//评论内容
	Content string `json:"content" gorm:"column:content" mapstructure:"content"`

	CreateAt      *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr   string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt      *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr   string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	CommentStatus int        `json:"CommentStatus" gorm:"column:comment_status" mapstructure:"CommentStatus"`
	IsDeleted     int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	COMMENT_STATUS_NORMAL int = 0 //正常
	COMMENT_STATUS_REVIEW int = 1 //审核

	COMMENT_NOT_DELETED int = 0 //未被删除
	COMMENT_IS_DELETED  int = 1 //被删除

	COMMENT_MAX_CONTENT int = 100 //评论内容最长100字

)

// 指定UserEssayComment对应的表名
func (UserEssayComment) TableName() string {
	return "user_essay_comment"
}
