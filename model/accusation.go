package model

import "time"

// Accusation 定义文章结构体
type Accusation struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//举报的用户id
	UserId int `json:"userId" gorm:"column:user_id" mapstructure:"userId"`
	//被举报文章id
	EssayId int `json:"essayId" gorm:"column:essay_id" mapstructure:"essayId"`
	//举报内容
	Content string `json:"content" gorm:"column:content" mapstructure:"content"`
	//举报时间
	AccusationTime *time.Time `json:"accusationTime" gorm:"column:accusation_time" mapstructure:"accusationTime"`

	CreateAt         *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr      string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt         *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr      string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	AccusationStatus int        `json:"accusationStatus" gorm:"column:accusation_status" mapstructure:"accusationStatus"`
	IsDeleted        int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	ACCUSATION_STATUS_WAIT    int = 0 //待审核
	ACCUSATION_STATUS_NORMAL  int = 1 //无违规
	ACCUSATION_STATUS_VIOLATE int = 2 //有违规

	ACCUSATION_NOT_DELETED int = 0 //未被删除
	ACCUSATION_IS_DELETED  int = 1 //被删除

	ACCUSATION_MAX_CONTENT int = 100 //举报内容最长100字

)

// 指定Accusation对应的表名
func (Accusation) TableName() string {
	return "accusation"
}
