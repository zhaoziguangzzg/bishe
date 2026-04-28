package model

import "time"

/*
-- knowledge_sharing.admin_user definition

CREATE TABLE `admin_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户密码',
  `email` varchar(100) NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone` bigint unsigned NOT NULL DEFAULT '0' COMMENT '电话',
  `avatar` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '图片名',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '1删除，0未删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='管理用户表';
*/

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
	//角色
	RoleId int `json:"roleId" gorm:"column:role_id" mapstructure:"roleId"`
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
