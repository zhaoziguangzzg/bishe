package main

import (
	"bishe/internal/app/knowledge_sharing/controller"
	"bishe/internal/app/knowledge_sharing/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	var err error
	err = service.LoggerInit()
	if err != nil {
		panic(err)
	}
	//Logger, err = zap.NewDevelopment()
	defer service.SyncLogger()

	err = service.LoadConfig()
	if err != nil {
		service.Logger.Error("LoadConfig err", zap.Error(err))
		panic(err)
	}

	// 初始化数据库
	err = service.ServiceInitDB(service.Cfg.Database.Dsn)
	if err != nil {
		service.Logger.Error("InitDB err", zap.Error(err))
		panic(err)
	}

	//初始化 Redis
	service.ServiceInitRedis(service.Cfg.Redis.Addr, service.Cfg.Redis.Password, service.Cfg.Redis.DB)

	// //初始化 kafka
	// err = service.ServiceInitKafka()
	// if err != nil {
	// 	service.Logger.Error("InitKafka err", zap.Error(err))
	// 	panic(err)
	// }
	// defer service.Closekafka()

	//创建 Gin 路由引擎
	r := gin.Default()

	r.Static("/static", "./static")

	// 注册路由
	r.POST("/api/user/add", controller.AddUserHandler) //绑定路径和函数，当客户端请求路径为""时使用这个函数处理请求
	r.GET("/api/user", controller.GetUserHandler)

	// 启动服务器
	service.Logger.Info("The server started at port", zap.String("port", "8080"))
	service.Logger.Error("Default error", zap.Error(r.Run(":8080")))
}
