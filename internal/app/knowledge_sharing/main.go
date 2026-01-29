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

	//用户模块
	r.POST("/api/user/add", controller.AddUserHandler)       //绑定路径和函数，当客户端请求路径为""时使用这个函数处理请求
	r.POST("/api/user/login", controller.UserLoginHandler)   //用户登录
	r.GET("/api/user/get", controller.GetUserHandler)        //获取某用户信息
	r.POST("/api/user/update", controller.UpdateUserHandler) //更新用户信息
	r.GET("/api/user/logout", controller.UserLogoutHandler)  //用户退出登录

	//圈子模块
	r.POST("/api/circle/add", controller.AddCircleHandler)             //创建圈子
	r.GET("/api/circle/all", controller.GetAllCircleHandler)           //获取全部圈子
	r.GET("/api/circle/get", controller.GetCircleHandler)              //获取圈子详情
	r.GET("/api/circle/create", controller.GetUserCreateCircleHandler) //获取用户创建的圈子
	r.GET("/api/circle/join", controller.GetUserJoinCircleHandler)     //获取用户加入的圈子
	r.GET("/api/circle/charge", controller.GetChargeCircleRankHandler) //获取全部付费圈子
	r.POST("/api/circle/update", controller.UpdateCircleHandler)       //更新圈子信息

	//用户加入圈子
	r.POST("/api/usercircle/add", controller.AddUserCircleJoinHandle) //创建用户加入圈子
	r.POST("/api/usercircle/quit", controller.UserQuitCircleHandler)  //用户退出圈子

	//文章模块
	r.POST("/api/essay/add", controller.AddEssayHandler)       //创建文章
	r.GET("/api/essay/all", controller.GetUserAllEssayHandler) //获取用户全部文章
	r.GET("/api/essay/get", controller.GetEssayHandler)        //查看文章

	// 启动服务器
	service.Logger.Info("The server started at port", zap.String("port", "8080"))
	service.Logger.Error("Default error", zap.Error(r.Run(":8080")))
}
