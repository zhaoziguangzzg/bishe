package model

import "time"

/*
create table role(
  id bigint unsigned not null auto_increment comment 'id',
  mids varchar(1000) not null default '' comment '权限列表',
  role_name varchar(100) not null default '' comment '角色名',
  create_at datetime not null default current_timestamp comment '创建时间',
  update_at DATETIME not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
  `is_deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '1删除，0未删除',
  primary key (`id`),
  unique key `uni_role_name`(`role_name`)
  ) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 comment='角色表';
*/

// Role 定义管理员角色结构体
type Role struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//角色名
	RoleName string `json:"roleName" gorm:"column:role_name" mapstructure:"roleName"`
	//权限列表
	Mids  string `json:"mids" gorm:"column:mids" mapstructure:"mids"`
	IsSys int    `json:"isSys" gorm:"column:is_sys" mapstructure:"isSys"`

	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	IsDeleted   int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	ADMIN_ROLE_IS_SYS = 1   //是否系统管理员
	ROLE_NAME_LEN_MAX = 100 //角色名最长为100个字
)

// 指定Role对应的表名
func (Role) TableName() string {
	return "role"
}

func (r *Role) IsSysRole() bool {
	return r.IsSys == ADMIN_ROLE_IS_SYS
}
