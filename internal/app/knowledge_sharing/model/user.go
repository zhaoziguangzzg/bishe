package model

import "time"

// User 定义用户结构体
type User struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//用户名
	Name string `json:"name" gorm:"column:name" mapstructure:"name"`
	//密码
	Password string `json:"password" gorm:"column:password" mapstructure:"password"`
	//头像
	Avatar string `json:"avatar" gorm:"column:avatar" mapstructure:"avatar"`
	//邮箱
	Email string `json:"email" gorm:"column:email" mapstructure:"email"`
	//年龄
	Age int `json:"age" gorm:"column:age" mapstructure:"age"`
	//电话
	Phone int `json:"phone" gorm:"column:phone" mapstructure:"phone"`
	//创建时间
	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	//更新时间
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	//用户状态
	UserStatus int `json:"userStatus" gorm:"column:user_status" mapstructure:"userStatus"`
	//是否被删除
	IsDeleted int `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	USER_STATUS_NORMAL int = 0 //正常
	USER_STATUS_REVIEW int = 1 //审核

	USER_NOT_DELETED int = 0 //未被删除
	USER_IS_DELETED  int = 1 //被删除

	USER_MAX_AGE int = 150 //用户年龄最大150
)

// 指定User对应的表名
func (User) TableName() string {
	return "users"
}
