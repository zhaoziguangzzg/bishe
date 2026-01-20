package model

import "time"

//用户加入圈子的结构体
type UserCircleJoin struct {
	Id          int        `json:"id" gorm:"column:id" mapstructure:"id"`
	UserId      int        `json:"userId" gorm:"column:user_id" mapstructure:"userId"`
	CircleId    int        `json:"circleId" gorm:"column:circle_id" mapstructure:"circleId"`
	JoinTime    *time.Time `json:"joinTime" gorm:"column:join_time" mapstructure:""`
	JoinTimeStr string     `json:"-" gorm:"-" mapstructure:"joinTime"`
	Status      int        `json:"status" gorm:"column:status" mapstructure:"status"`
}

const (
	UserCircleJoinQuit   int = 0 //退出
	UserCircleJoinNormal int = 1 //正常
)

// 指定UserCircleJoin对应的表名
func (UserCircleJoin) TableName() string {
	return "user_circle_join"
}
