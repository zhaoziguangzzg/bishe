package model

import "time"

/*
create table purchase(
  id bigint unsigned not null auto_increment comment 'id',
  user_id bigint unsigned not null default 0 comment '用户id',
  course_id bigint unsigned not null default 0 comment '课程id',
  create_at datetime not null default current_timestamp comment '创建时间',
  update_at DATETIME not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
  `purchase_status` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '1已购买，0未购买',
  primary key (`id`),
  key `purchase_uid_cid` (`user_id`,`course_id`)
  ) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 comment='购买课程表';
*/

//Purchase 定义购买课程结构体
type Purchase struct {
	//课时id
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//课程id
	CourseId int `json:"courseId" gorm:"column:course_id" mapstructure:"courseId"`
	//用户id
	UserId int `json:"userId" gorm:"column:user_id" mapstructure:"userId"`

	CreateAt       *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr    string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt       *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr    string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	PurchaseStatus int        `json:"purchaseStatus" gorm:"column:purchase_status" mapstructure:"purchaseStatus"`
}

const (
	PURCHASE_STATUS_BUY     int = 1 //已购买
	PURCHASE_STATUS_OVER    int = 2 //过期
	PURCHASE_STATUS_NOT_BUY int = 0 //未购买
)

// 指定Purchase对应的表名
func (Purchase) TableName() string {
	return "purchase"
}
