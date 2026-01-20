package model

//用户评论文章的结构体
type UserEssayComment struct {
	Id       int    `json:"id" gorm:"column:id" mapstructure:"id"`
	UserId   int    `json:"userId" gorm:"column:user_id" mapstructure:"userId"`
	CircleId int    `json:"circleId" gorm:"column:circle_id" mapstructure:"circleId"`
	EssayId  int    `json:"essayId" gorm:"column:essay_id" mapstructure:"essayId"`
	Comment  string `json:"comment" gorm:"column:comment" mapstructure:"comment"`
}

// 指定UserEssayComment对应的表名
func (UserEssayComment) TableName() string {
	return "user_essay_comment"
}
