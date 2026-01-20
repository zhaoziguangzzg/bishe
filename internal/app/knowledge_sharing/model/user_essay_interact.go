package model

//用户点赞文章的结构体
type UserEssayInteract struct {
	Id            int    `json:"id" gorm:"column:id" mapstructure:"id"`
	UserId        int    `json:"userId" gorm:"column:user_id" mapstructure:"userId"`
	CircleId      int    `json:"circleId" gorm:"column:circle_id" mapstructure:"circleId"`
	EssayId       int    `json:"essayId" gorm:"column:essay_id" mapstructure:"essayId"`
	LikeStatus    int    `json:"likeStatus" gorm:"column:like_status" mapstructure:"likeStatus"`
	CollectStatus int    `json:"collectStatus" gorm:"column:collect_status" mapstructure:"collectStatus"`
	Favorite      string `json:"favorite" gorm:"column:favorite" mapstructure:"favorite"`
}

const (
	UserEssayInteractNotJoin int = 0 //退出
	UserEssayInteractJoin    int = 1 //正常
)

// 指定UserEssayInteract对应的表名
func (UserEssayInteract) TableName() string {
	return "user_essay_interact"
}
