package model

// User 定义用户结构体
type User struct {
	Id       int    `json:"id" gorm:"column:id" mapstructure:"id"`
	Account  int    `json:"account" gorm:"column:account" mapstructure:"account"`
	Password string `json:"password" gorm:"column:password" mapstructure:"password"`
	Email    string `json:"email" gorm:"column:email" mapstructure:"email"`
	Age      int    `json:"age" gorm:"column:age" mapstructure:"age"`
	Phone    int    `json:"phone" gorm:"column:phone" mapstructure:"phone"`
}

// 指定User对应的表名
func (User) TableName() string {
	return "users"
}
