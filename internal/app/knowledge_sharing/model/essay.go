package model

import "time"

// Essay 定义文章结构体
type Essay struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//文章标题
	Title string `json:"title" gorm:"column:title" mapstructure:"title"`
	//所在圈子id
	CircleId int `json:"circleId" gorm:"column:circle_id" mapstructure:"circleId"`
	//文章内容
	Content string `json:"content" gorm:"column:content" mapstructure:"content"`
	//作者id
	AuthorId int `json:"authorId" gorm:"column:author_id" mapstructure:"authorId"`
	//点赞数
	LikeNum int `json:"likeNum" gorm:"column:like_num" mapstructure:"likeNum"`
	//收藏数
	CollectNum int `json:"collectNum" gorm:"column:collect_num" mapstructure:"collectNum"`
	//评论数
	CommentNum int `json:"commentNum" gorm:"column:comment_num" mapstructure:"commentNum"`
	//周刊
	IsWeekly int `json:"isWeekly" gorm:"column:is_weekly" mapstructure:"isWeekly"`
	//精华
	IsEssence int `json:"isEssence" gorm:"column:is_essence" mapstructure:"isEssence"`

	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`

	//文章评价
	Quality     int `json:"quality" gorm:"quality" mapstructure:"quality"`
	EssayStatus int `json:"essayStatus" gorm:"column:essay_status" mapstructure:"essayStatus"`
	IsDeleted   int `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	ESSAY_STATUS_NORMAL int = 0 //正常
	ESSAY_STATUS_REVIEW int = 1 //审核

	ESSAY_NOT_DELETED int = 0 //未被删除
	ESSAY_IS_DELETED  int = 1 //被删除

	ESSAY_IS_WEEKLY  int = 1 //是周刊
	ESSAY_NOT_WEEKLY int = 0 //不是周刊

	ESSAY_IS_ESSENCE  int = 1 //是精华
	ESSAY_NOT_ESSENCE int = 0 //不是精华

	ESSAY_MAX_TITLE   int = 100 //文章标题最长100字
	ESSAY_MAX_CONTENT int = 200 //文章内容最长200字

)

// 指定Essay对应的表名
func (Essay) TableName() string {
	return "essay"
}
