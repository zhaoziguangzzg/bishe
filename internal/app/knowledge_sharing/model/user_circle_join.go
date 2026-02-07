package model

import "time"

//用户加入圈子的结构体
type UserCircleJoin struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//加入圈子用户id
	UserId int `json:"userId" gorm:"column:user_id" mapstructure:"userId"`
	//用户加入的圈子id
	CircleId int `json:"circleId" gorm:"column:circle_id" mapstructure:"circleId"`

	JoinTime    *time.Time `json:"joinTime" gorm:"column:join_time" mapstructure:""`
	JoinTimeStr string     `json:"-" gorm:"-" mapstructure:"joinTime"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	JoinStatus  int        `json:"joinStatus" gorm:"column:join_status" mapstructure:"joinStatus"`
}

const (
	USER_CIRCLE_JOIN_JOIN_STATUS_NO_JOIN int = 0 //未加入
	USER_CIRCLE_JOIN_JOIN_STATUS_JOIN    int = 1 //加入
)

// 指定UserCircleJoin对应的表名
func (UserCircleJoin) TableName() string {
	return "user_circle_join"
}
