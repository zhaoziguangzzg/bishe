package model

import "time"

// Advert 定义广告结构体
type Advert struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//广告位置
	Position int `json:"position" gorm:"column:position" mapstructure:"position"`
	//跳转链接
	AdvertAddr string `json:"advertAddr" gorm:"column:advert_addr" mapstructure:"advertAddr"`
	//广告内容
	Content string `json:"content" gorm:"column:content" mapstructure:"content"`
	//广告开始时间
	StartTime *time.Time `json:"startTime" gorm:"column:start_time" mapstructure:"startTime"`
	//广告结束时间
	EndTime *time.Time `json:"endTime" gorm:"column:end_time" mapstructure:"endTime"`
	//广告图片
	Img string `json:"img" gorm:"column:img" mapstructure:"img"`

	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	IsDeleted   int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	ADVERT_MAX_ADDR              int = 200 //广告地址最长200字
	ADVERT_MAX_CONTENT           int = 100 //广告内容最长100字
	ADVERT_POSITION_CIRCLE_INDEX int = 1   //圈子首页
	ADVERT_POSITION_COURSE_INDEX int = 2   //课程首页
	ADVERT_POSITION_USER_PROFILE int = 3   //用户个人中心
	ADVERT_POSITION_INDEX        int = 4   //首页
)

// 指定Advert对应的表名
func (Advert) TableName() string {
	return "advert"
}
