package model

import "time"

//订单结构体
type Orders struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//用户id
	Uid int `json:"uid" gorm:"column:uid" mapstructure:"uid"`
	//圈子id
	Cid int `json:"cid" gorm:"column:cid" mapstructure:"cid"`
	//圈子价格
	Price int `json:"price" gorm:"column:price" mapstructure:"price"`
	//折扣金额
	Discount int `json:"discount" gorm:"column:discount" mapstructure:"discount"`

	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	OrderStatus int        `json:"orderStatus" gorm:"column:order_status" mapstructture:"orderStatus"`
}

const (
	ORDER_STATUS_WAIT   int = 0 //待支付
	ORDER_STATUS_PAID   int = 1 //已支付
	ORDER_STATUS_CANCEL int = 2 //已取消
)

// 指定Orders对应的表名
func (Orders) TableName() string {
	return "orders"
}
