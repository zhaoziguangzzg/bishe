package model

import "time"

/*
create table lesson(
  id bigint unsigned not null auto_increment comment 'id',
  course_id bigint unsigned not null default 0 comment '课程id',
  title varchar(100) not null default '' comment '课时标题',
  content text character set utf8mb4 collate utf8mb4_0900_ai_ci not null comment '课时内容',
  create_at datetime not null default current_timestamp comment '创建时间',
  update_at DATETIME not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
  `is_deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '1删除，0未删除',
  primary key (`id`)
  ) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 comment='课时表';
*/

//Lesson 定义课时结构体
type Lesson struct {
	//课时id
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//课程id
	CourseId int `json:"courseId" gorm:"column:course_id" mapstructure:"courseId"`
	//课时标题
	Title string `json:"title" gorm:"column:title" mapstructure:"title"`
	//课时内容
	Content string `json:"content" gorm:"column:content" mapstructure:"content"`

	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	IsDeleted   int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	LESSON_TITLE_MAX   int = 100   //课时标题最长100字
	LESSON_CONTENT_MAX int = 65535 //课时内容最长65535字

)

// 指定Lesson对应的表名
func (Lesson) TableName() string {
	return "lesson"
}
