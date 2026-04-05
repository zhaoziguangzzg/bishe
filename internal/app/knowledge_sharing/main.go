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
	r.POST("/api/user/update", controller.UpdateUserHandler) //更新用户信息
	r.POST("/api/user/login", controller.UserLoginHandler)   //用户登录
	r.GET("/api/user/get", controller.GetUserHandler)        //获取某用户信息
	r.GET("/api/user/logout", controller.UserLogoutHandler)  //用户退出登录

	//圈子模块
	r.POST("/api/circle/add", controller.AddCircleHandler)             //创建圈子
	r.POST("/api/circle/update", controller.UpdateCircleHandler)       //更新圈子信息
	r.GET("/api/circle/all", controller.GetAllCircleHandler)           //获取全部圈子
	r.GET("/api/circle/get", controller.GetCircleHandler)              //获取圈子详情
	r.GET("/api/circle/create", controller.GetUserCreateCircleHandler) //获取用户创建的圈子
	r.GET("/api/circle/join", controller.GetUserJoinCircleHandler)     //获取用户加入的圈子
	r.GET("/api/circle/charge", controller.GetChargeCircleRankHandler) //获取全部付费圈子
	r.GET("/api/circle/free", controller.GetFreeCircleRankHandler)     //获取全部免费圈子

	//用户加入圈子
	r.POST("/api/usercircle/add", controller.AddUserCircleJoinHandle) //创建用户加入圈子
	r.POST("/api/usercircle/quit", controller.QuitCircleHandler)      //用户退出圈子

	//文章模块
	r.POST("/api/essay/add", controller.AddEssayHandler)                         //创建文章
	r.POST("/api/essay/update", controller.UpdateEssayHandler)                   //更新文章
	r.POST("/api/essay/delete", controller.DeletedEssayByUpdateIsDeletedHandler) //删除文章
	r.GET("/api/essay/get", controller.GetEssayHandler)                          //查看文章
	r.GET("/api/essay/circleall", controller.GetCircleAllEssayHandler)           //获取圈子全部文章
	r.GET("/api/essay/userall", controller.GetUserAllEssayHandler)               //获取用户全部文章

	//周刊
	r.POST("/api/essay/addweekly", controller.AddEssayWeeklyHandler)    //将文章添加周刊
	r.GET("/api/essay/getweekly", controller.GetEssayWeeklylistHandler) //获取文章周刊
	//精粹
	r.POST("/api/essay/addessence", controller.AddEssayEssenceHandler) //将文章添加精粹
	r.GET("/api/essay/getessence", controller.GetEssayEssonceHandler)  //获取文章精粹

	//点赞
	r.POST("/api/like/add", controller.AddUserEssayLikeHandler)       //添加点赞
	r.POST("/api/like/cancel", controller.CancelUserEssayLikeHandler) //更新点赞删除
	r.GET("/api/like/get", controller.GetUserEssayLikeHandler)        //获取点赞
	r.GET("/api/like/all", controller.GetUserAllLikeHandler)          //获取用户点赞

	//收藏夹
	r.POST("/api/favorite/add", controller.AddFavoriteHandler)                         //添加收藏夹
	r.POST("/api/favorite/update", controller.UpdateFavoriteTitleHandler)              //修改收藏夹名
	r.POST("/api/favorite/delete", controller.DeletedFavoriteByUpdateIsDeletedHandler) //删除收藏夹
	r.GET("/api/favorite/get", controller.GetFavoriteHandler)                          //获取收藏夹
	r.GET("/api/favorite/all", controller.GetUserAllFavoriteHandler)                   //获取用户全部收藏夹
	//TODO 去除唯一键，新建时判断该数量=1，就不能新建

	//收藏
	r.POST("/api/collect/add", controller.AddUserEssayCollectHandler)   //添加收藏
	r.POST("/api/collect/cancel", controller.CancelEssayCollectHandler) //更新收藏删除状态
	r.GET("/api/collect/get", controller.GetEssayCollectHandler)        //获取收藏
	r.GET("/api/collect/all", controller.GetUserAllCollectHandler)      //获取用户收藏夹的全部收藏

	//评论
	r.POST("/api/comment/add", controller.AddUserEssayCommentHandle)                 //创建用户评论
	r.POST("/api/comment/delete", controller.DeletedCommentByUpdateIsDeletedHandler) //删除评论
	r.GET("/api/comment/essayall", controller.GetEssayAllCommentHandle)              //获取文章全部评论
	r.GET("/api/comment/userall", controller.GetUserAllCommentHandler)               //获取用户全部评论

	//关注
	r.POST("/api/follow/add", controller.AddUserFollowHandler)       //添加关注
	r.POST("/api/follow/cancel", controller.CancelUserFollowHandler) //更新关注删除
	r.GET("/api/follow/get", controller.GetUserFollowHandler)        //获取用户关注
	r.GET("/api/follow/all", controller.GetUserAllFollowHandler)     //获取用户关注列表
	r.GET("/api/follow/allfan", controller.GetUserAllFanHandler)     //获取用户粉丝列表

	//举报
	r.POST("/api/accusation/add", controller.AddUserAccusationEssayHandler)     //创建举报
	r.GET("/api/accusation/all", controller.GetAllAccusationEssayHandler)       //获取全部未审核举报
	r.GET("/api/accusation/get", controller.GetEssayContentByAccusationHandler) //获取举报内容文章
	r.POST("/api/accusation/update", controller.UpdateAccusationStatusHandler)  //更新举报状态

	//反馈
	r.POST("/api/feedback/add", controller.AddUserFeedbackHandler)         //创建反馈
	r.GET("/api/feedback/all", controller.GetAllFeedbackHandler)           //获取全部未处理反馈
	r.GET("/api/feedback/get", controller.GetFeedbackContentHandler)       //获取反馈
	r.POST("/api/feedback/update", controller.UpdateFeedbackStatusHandler) //更新反馈状态

	//私信
	r.POST("/api/chat/add", controller.AddChatHandler)    //添加私信
	r.GET("/api/chat/get", controller.GetChatListHandler) //获取私信记录

	//联系人
	r.GET("/api/contact/all", controller.GetChatContactListHandler) //获取最近联系人列表

	//通知
	r.GET("/api/notice/all", controller.GetNoticeListHandler)           //获取通知列表
	r.GET("/api/notice/alltype", controller.GetNoticeListByTypeHandler) //获取某类型通知列表

	//统计
	r.GET("/api/stat/all", controller.GetUserStatListHandler) //获取全部统计数据列表

	//等级
	r.GET("/api/levelrecord/all", controller.GetUserCircleLevelAllRecordHandler) //获取用户在圈子全部等级详情

	//管理员用户
	r.POST("/api/adminuser/add", controller.AddAdminUserHandler)       //添加管理员用户信息
	r.POST("/api/adminuser/update", controller.UpdateAdminUserHandler) //更新管理员用户信息
	r.POST("/api/adminuser/login", controller.AdminUserLoginHandler)   //管理员用户登录
	r.GET("/api/adminuser/get", controller.GetAdminUserHandler)        //获取某管理员用户信息
	r.GET("/api/adminuser/logout", controller.AdminUserLogoutHandler)  //管理员用户退出登录
	r.POST("/api/adminuser/delete", controller.DeleteAdminUserHandler) // 删除管理员用户
	// TODO 角色
	// TODO 菜单权限

	//公告
	r.POST("/api/announce/add", controller.AddAnnounceHandler)                         //创建公告
	r.GET("/api/announce/all", controller.GetAllAnnounceHandler)                       //获取全部公告
	r.GET("/api/announce/get", controller.GetAnnounceHandler)                          //查看公告
	r.POST("/api/announce/update", controller.UpdateAnnounceHandler)                   //更新公告
	r.POST("/api/announce/delete", controller.DeletedAnnounceByUpdateIsDeletedHandler) //删除公告

	//广告
	r.POST("/api/advert/add", controller.AddAdvertHandler)                         //创建广告
	r.GET("/api/advert/all", controller.GetAllAdvertHandler)                       //获取全部广告
	r.GET("/api/advert/get", controller.GetAdvertHandler)                          //查看广告
	r.POST("/api/advert/update", controller.UpdateAdvertHandler)                   //更新广告
	r.POST("/api/advert/delete", controller.DeletedAdvertByUpdateIsDeletedHandler) //删除广告

	//联系人列表（send-receive）
	r.POST("/api/contact/add", controller.AddUserContactHandler)       //添加联系人
	r.POST("/api/contact/delete", controller.DeleteUserContactHandler) //删除联系人
	//r.GET("/api/contact/all", controller.GetUserAllContactHandler)     //获取用户全部联系人
	r.GET("/api/contact/get", controller.GetUserContactHandler) //获取联系人

	//TODO 与某人消息记录
	r.GET("/api/information/send", controller.GetUserSendInformationHandler)       //获取用户发送的各消息
	r.GET("/api/information/receive", controller.GetUserReceiveInformationHandler) //获取用户接收的各消息

	// 启动服务器
	service.Logger.Info("The server started at port", zap.String("port", "8080"))
	service.Logger.Error("Default error", zap.Error(r.Run(":8080")))
}
