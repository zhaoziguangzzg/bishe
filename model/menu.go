package model

import "time"

/*
create table menu(
  id bigint unsigned not null auto_increment comment 'id',
  path bigint unsigned not null default 0 comment '权限路径',
  menu_name varchar(100) not null default '' comment '权限名',
  create_at datetime not null default current_timestamp comment '创建时间',
  update_at DATETIME not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
  `is_deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '1删除，0未删除',
  primary key (`id`),
  unique key `uni_menu_name`(`menu_name`)
  ) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 comment='权限表';
*/

// Menu 定义管理员权限结构体
type Menu struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//权限路径
	Path string `json:"path" gorm:"column:path" mapstructure:"path"`
	//权限名
	MenuName    string     `json:"menuName" gorm:"column:menu_name" mapstructure:"menuName"`
	Weight      int        `json:"weight" gorm:"column:weight" mapstructure:"weight"`
	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	IsDeleted   int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

// 指定Menu对应的表名
func (Menu) TableName() string {
	return "menu"
}
