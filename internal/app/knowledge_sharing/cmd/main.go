package main

import (
	cmd "bishe/internal/app/knowledge_sharing/cmd/purchase"
	"bishe/internal/app/knowledge_sharing/service"
)

func main() {

	err := service.ServiceInit()
	if err != nil {
		panic(err)
	}
	//main结束之前将日志写到文件
	defer service.SyncLogger()

	cmd.PurchaseExpire()
}
