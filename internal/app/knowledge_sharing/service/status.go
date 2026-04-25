package service

import (
	"bishe/internal/app/knowledge_sharing/model"
	"errors"
)

var orderStatusActionMap = map[int]map[int]int{
	//status
	model.ORDER_STATUS_UNPAID: map[int]int{
		//action
		model.ORDER_ACTION_PAY:    model.ORDER_STATUS_PAID,
		model.ORDER_ACTION_CANCEL: model.ORDER_STATUS_CANCELED,
		model.ORDER_ACTION_EXPIRE: model.ORDER_STATUS_EXPIRED,
	},

	model.ORDER_STATUS_PAID: map[int]int{
		model.ORDER_ACTION_REFUND: model.ORDER_STATUS_REFUND,
	},
}

// 根据status和action转换新status
func MakeOrderStatus(status int, action int) (statusNew int, err error) {
	actionMap, ok := orderStatusActionMap[status]
	if !ok {
		err = errors.New("status err")
		return
	}

	statusNew, ok = actionMap[action]
	if !ok {
		err = errors.New("action err")
		return
	}
	return

	/*
		switch status {
		case model.ORDER_STATUS_UNPAID:
			if action == model.ORDER_ACTION_PAY {
				convStatus = model.ORDER_STATUS_PAID
			} else if action == model.ORDER_ACTION_CANCEL {
				convStatus = model.ORDER_STATUS_CANCELED
			} else if action == model.ORDER_ACTION_EXPIRE {
				convStatus = model.ORDER_STATUS_EXPIRED
			}
		case model.ORDER_STATUS_PAID:
			if action == model.ORDER_ACTION_REFUND {
				convStatus = model.ORDER_STATUS_REFUND
			}
		case model.ORDER_STATUS_CANCELED:
		case model.ORDER_STATUS_EXPIRED:
		case model.ORDER_STATUS_REFUND:
		default:

		}

		return
	*/
}

var purchaseStatusActionMap = map[int]map[int]int{
	//status
	model.PURCHASE_STATUS_UNPAID: map[int]int{
		//action
		model.PURCHASE_ACTION_PAY:    model.PURCHASE_STATUS_PAID,
		model.PURCHASE_ACTION_CANCEL: model.PURCHASE_STATUS_CANCELED,
		model.PURCHASE_ACTION_EXPIRE: model.PURCHASE_STATUS_EXPIRED,
	},

	model.PURCHASE_STATUS_PAID: map[int]int{
		model.PURCHASE_ACTION_REFUND: model.PURCHASE_STATUS_REFUND,
	},
}

// 根据status和action转换新status
func MakePurchaseStatus(status int, action int) (statusNew int, err error) {
	actionMap, ok := purchaseStatusActionMap[status]
	if !ok {
		err = errors.New("status err")
		return
	}

	statusNew, ok = actionMap[action]
	if !ok {
		err = errors.New("action err")
		return
	}

	return
}
