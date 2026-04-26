package model

import "time"

/*
-- knowledge_sharing.users definition

CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户密码',
  `email` varchar(100) NOT NULL DEFAULT '' COMMENT '邮箱',
  `age` int unsigned NOT NULL DEFAULT '0' COMMENT '年龄',
  `phone` bigint unsigned NOT NULL DEFAULT '0' COMMENT '电话',
  `avatar` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '图片名',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `user_status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0正常，1审核',
  `is_deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '1删除，0未删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';
*/

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
