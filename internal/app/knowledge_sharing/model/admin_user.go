package model

import "time"

// AdminUser 定义管理员用户结构体
type AdminUser struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//用户名
	Name string `json:"name" gorm:"column:name" mapstructure:"name"`
	//密码
	Password string `json:"password" gorm:"column:password" mapstructure:"password"`
	//邮箱
	Email string `json:"email" gorm:"column:email" mapstructure:"email"`
	//电话
	Phone int `json:"phone" gorm:"column:phone" mapstructure:"phone"`
	//头像
	Avatar string `json:"avatar" gorm:"column:avatar" mapstructure:"avatar"`

	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	IsDeleted   int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

// 指定AdminUser对应的表名
func (AdminUser) TableName() string {
	return "admin_user"
}
