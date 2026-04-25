package cmd

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"time"

	"go.uber.org/zap"
)

// 查询30分钟后未支付的订单
func PurchaseExpire() {

	for {
		t := time.Now().Add(-30 * time.Minute)
		purchases, err := service.GetPurchaseByStatusTime(model.PURCHASE_STATUS_UNPAID, t, 100)
		if err != nil {
			service.Logger.Error("GetPurchaseByStatusTime err", zap.Error(err))
		}

		if len(purchases) == 0 {
			time.Sleep(5 * time.Second)
			continue
		}

		for _, v := range purchases {
			purchaseStatus := v.PurchaseStatus
			id := v.Id
			newStatus, err := service.MakePurchaseStatus(purchaseStatus, model.PURCHASE_ACTION_EXPIRE)
			if err != nil {
				service.Logger.Error("MakePurchaseStatus", zap.Int("id", id), zap.Int("status", purchaseStatus), zap.Error(err))
				continue
			}

			affectRows, err := service.UpdatePurchaseStatusById(id, purchaseStatus, newStatus)
			if affectRows == 0 || err != nil {
				service.Logger.Error("UpdatePurchaseStatusById", zap.Int("id", id), zap.Error(err))
				continue
			}

			service.Logger.Info("UpdatePurchaseStatusById", zap.Int("id", id), zap.Int("newStatus", newStatus))
		}

	}

}
