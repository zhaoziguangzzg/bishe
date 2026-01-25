package model

import "time"

//用户加入圈子的结构体
type UserCircleJoin struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//加入圈子用户id
	UserId int `json:"userId" gorm:"column:user_id" mapstructure:"userId"`
	//用户加入的圈子id
	CircleId int `json:"circleId" gorm:"column:circle_id" mapstructure:"circleId"`

	JoinTime      *time.Time `json:"joinTime" gorm:"column:join_time" mapstructure:""`
	JoinTimeStr   string     `json:"-" gorm:"-" mapstructure:"joinTime"`
	UpdateAt      *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr   string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	NotJoinStatus int        `json:"notJoinStatus" gorm:"column:not_join_status" mapstructure:"notJoinStatus"`
}

const (
	USER_CIRCLE_NOT_NO_JOIN int = 0 //加入
	USER_CIRCLE_NOT_JOIN    int = 1 //未加入
)

// 指定UserCircleJoin对应的表名
func (UserCircleJoin) TableName() string {
	return "user_circle_join"
}
