package model

import "time"

//支付结构体
type Orders struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//用户id
	Uid int `json:"uid" gorm:"column:uid" mapstructure:"uid"`
	//圈子id
	Cid int `json:"cid" gorm:"column:cid" mapstructure:"cid"`
	//支付金额
	Price int `json:"price" gorm:"column:price" mapstructure:"price"`
	//折扣金额
	Discount float64 `json:"discount" gorm:"column:discount" mapstructure:"discount"`

	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	OrderStatus int        `json:"orderStatus" gorm:"column:order_status" mapstructture:"orderStatus"`
}

const (
	ORDER_STATUS_UNPAID   int = 0 //未支付
	ORDER_STATUS_PAID     int = 1 //已支付
	ORDER_STATUS_CANCELED int = 2 //已取消
	ORDER_STATUS_EXPIRED  int = 3 //已过期
	ORDER_STATUS_REFUND   int = 4 //已退款

	ORDER_ACTION_PAY    int = 1 //支付
	ORDER_ACTION_CANCEL int = 2 //取消
	ORDER_ACTION_EXPIRE int = 3 //过期
	ORDER_ACTION_REFUND int = 4 //退款
)

// 指定Orders对应的表名
func (Orders) TableName() string {
	return "orders"
}
