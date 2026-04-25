package cmd

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// 查询30分钟后未支付的订单
func PurchaseExpire() {
	sigChan := make(chan os.Signal, 1)
	//windows收不到命令，只能收到 Ctrl+C / Ctrl+Break,不是sleep，是收不到信号
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

purchase:
	for {
		service.Logger.Info("UpdatePurchaseStatusById")

		select {
		case sig := <-sigChan:
			//当收到信号时，记录日志，结束循环
			service.Logger.Info("PurchaseExpire get sig", zap.Any("sig", sig))
			break purchase
		default:
			//没有信号，继续for循环
		}

		service.Logger.Info("UpdatePurchaseStatusById2")

		t := time.Now().Add(-30 * time.Minute)
		purchases, err := service.GetPurchaseByStatusTime(model.PURCHASE_STATUS_UNPAID, t, 100)
		if err != nil {
			service.Logger.Error("GetPurchaseByStatusTime err", zap.Error(err))
		}

		if len(purchases) == 0 {
			service.Logger.Info("UpdatePurchaseStatusById3")
			time.Sleep(time.Second * 2)
			// <-time.After(time.Second * 2)

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
