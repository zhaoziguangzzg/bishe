package model

import "time"

/*
create table course(
  id bigint unsigned not null auto_increment comment 'id',
  uid bigint unsigned not null default 0 comment '创建者id',
  cid bigint unsigned not null default 0 comment '所属圈子id',
  title varchar(100) not null default '' comment '课程标题',
  content varchar(100) not null default '' comment '课程内容简介',
  price int unsigned not null default 0 comment '课程价格',
  join_num bigint unsigned not null default 0 comment '加入人数',
  create_at datetime not null default current_timestamp comment '创建时间',
  update_at DATETIME not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
  `is_deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '1删除，0未删除',
  primary key (`id`),
  key `key_cid_title`(`cid`,`title`)
  ) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 comment='课程表';
*/

// Course 定义课程结构体
type Course struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//课程标题
	Title string `json:"title" gorm:"column:title" mapstructure:"title"`
	// //所在圈子id
	// Cid int `json:"cid" gorm:"column:cid" mapstructure:"cid"`
	//课程作者id
	Uid int `json:"uid" gorm:"column:uid" mapstructure:"uid"`
	//课程内容简介
	Content string `json:"content" gorm:"column:content" mapstructure:"content"`
	//价格
	Price int `json:"price" gorm:"column:price" mapstructure:"price"`
	//加入人数
	JoinNum int `json:"joinNum" gorm:"column:join_num" mapstructure:"joinNum"`

	CreateAt     *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr  string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt     *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr  string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	CourseStatus int        `json:"courseStatus" gorm:"column:course_status" mapstructure:"courseStatus"`
	IsDeleted    int        `json:"isDeleted" gorm:"column:is_deleted" mapstructure:"isDeleted"`
}

const (
	COURSE_TITLE_MAX   int = 100   //课程标题最长100字
	COURSE_CONTENT_MAX int = 100   //课程内容简介最长100字
	COURSE_PRICE_MAX   int = 10000 //课程价格最长10000元

	COURSE_STATUS_PUBLISHED   int = 1 //已发布
	COURSE_STATUS_UNPUBLISHED int = 0 //未发布
)

// 指定Course对应的表名
func (Course) TableName() string {
	return "course"
}
