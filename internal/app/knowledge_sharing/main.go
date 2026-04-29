package main

import (
	"bishe/internal/app/knowledge_sharing/controller"
	"bishe/internal/app/knowledge_sharing/middleware"
	"bishe/internal/app/knowledge_sharing/service"
	"html/template"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {

	err := service.ServiceInit()
	if err != nil {
		panic(err)
	}
	//main结束之前将日志写到文件
	defer service.SyncLogger()

	//创建 Gin 路由引擎
	r := gin.Default()
	// 给 Gin 设置自定义模板引擎
	r.SetHTMLTemplate(tpl)

	r.Static("/img", "web/img")

	//页面路由
	r.GET("/", controller.IndexPageHandler)
	r.GET("/page/user/login", controller.LoginPageHandler)
	r.GET("/page/user/register", controller.RegisterPageHandler)

	userPage := r.Group("/page/user")
	userPage.Use(middleware.MiddlewareUserLoginPage())
	userPage.GET("/profile", controller.ProfilePageHandler)
	userPage.GET("/edit", controller.EditPageHandler)
	userPage.GET("/edit-password", controller.EditPasswordPageHandler)

	r.POST("/api/user/login", controller.UserLoginHandler) //用户登录
	r.POST("/api/user/add", controller.AddUserHandler)     //绑定路径和函数，当客户端请求路径为""时使用这个函数处理请求

	userApi := r.Group("/api/user")
	userApi.Use(middleware.MiddlewareUserLoginApi())
	userApi.POST("/update", controller.UpdateUserHandler)                 //更新用户信息
	userApi.POST("/updatepassword", controller.UpdateUserPasswordHandler) //更新用户密码
	userApi.GET("/get", controller.GetUserHandler)                        //获取某用户信息
	userApi.GET("/get-by-id", controller.GetUserByIdHandler)              //通过uid获取用户信息
	userApi.GET("/logout", controller.UserLogoutHandler)                  //用户退出登录

	//圈子模块
	//圈子列表（排行榜更多）
	r.GET("/page/circle/list", controller.CircleListPageHandler)

	circleUserLoginPage := r.Group("/page/circle")
	circleUserLoginPage.Use(middleware.MiddlewareUserLoginPage())
	//创建圈子页面
	circleUserLoginPage.GET("/add", controller.AddCirclePageHandler)
	//获取圈子详情页面
	circleUserLoginPage.GET("/detail", controller.CircleDetailPageHandler)
	//加入的圈子页面
	circleUserLoginPage.GET("/index", controller.CircleIndexPageHandler)
	//创建的圈子页面
	circleUserLoginPage.GET("/create", controller.CircleCreatePageHandler)

	circleOwnerPage := r.Group("/page/circle/")
	circleOwnerPage.Use(middleware.MiddlewareUserLoginPage(), middleware.MiddlewareIsJoinCirclePage(),
		middleware.MiddlewareIsCircleOwnerPage())
	//修改圈子页面
	circleOwnerPage.GET("edit", controller.EditCirclePageHandler)

	circleUserLoginApi := r.Group("/api/circle")
	circleUserLoginApi.Use(middleware.MiddlewareUserLoginApi())
	//创建圈子
	circleUserLoginApi.POST("/add", controller.AddCircleHandler)
	//获取圈子详情
	circleUserLoginApi.GET("/get", controller.GetCircleHandler)
	//获取用户加入的圈子
	circleUserLoginApi.GET("/join", controller.GetUserJoinCircleHandler)
	//获取全部圈子
	circleUserLoginApi.GET("/all", controller.GetAllCircleHandler)
	//获取全部付费圈子
	circleUserLoginApi.GET("/charge", controller.GetChargeCircleRankHandler)
	//获取全部免费圈子
	circleUserLoginApi.GET("/free", controller.GetFreeCircleRankHandler)
	//获取用户创建的圈子
	circleUserLoginApi.GET("/create", controller.GetUserCreateCircleHandler)

	circleOwnerApi := r.Group("/api/circle")
	circleOwnerApi.Use(middleware.MiddlewareUserLoginApi(), middleware.MiddlewareIsJoinCircleApi(),
		middleware.MiddlewareIsCircleOwnerApi())
	//更新圈子信息
	circleOwnerApi.POST("/update", controller.UpdateCircleHandler)

	//用户加入圈子
	r.POST("/api/usercircle/add", middleware.MiddlewareUserLoginApi(), controller.AddUserCircleJoinHandle) //创建用户加入圈子
	//用户退出圈子
	r.POST("/api/usercircle/quit", middleware.MiddlewareUserLoginApi(), middleware.MiddlewareIsJoinCircleApi(),
		controller.QuitCircleHandler)

	//文章模块
	//创建文章页面
	essayUserLoginPage := r.Group("/page/essay")
	essayUserLoginPage.Use(middleware.MiddlewareUserLoginPage())
	essayUserLoginPage.GET("/add", controller.AddEssayPageHandler)
	essayUserLoginPage.GET("/detail", controller.EssayDetailPageHandler)
	essayUserLoginPage.GET("/search", controller.SearchEssayPageHandler) //获取全部搜索记录
	essayUserLoginPage.GET("/edit", controller.EditEssayPageHandler)

	essayUserLoginApi := r.Group("/api/essay")
	essayUserLoginApi.Use(middleware.MiddlewareUserLoginApi())
	essayUserLoginApi.POST("/add", controller.AddEssayHandler)                         //创建文章
	essayUserLoginApi.POST("/delete", controller.DeletedEssayByUpdateIsDeletedHandler) //删除文章
	essayUserLoginApi.GET("/get", controller.GetEssayHandler)                          //查看文章
	essayUserLoginApi.GET("/circle-all", controller.GetCircleAllEssayHandler)          //获取圈子全部文章
	essayUserLoginApi.GET("/user-all", controller.GetUserAllEssayHandler)              //获取用户全部文章
	essayUserLoginApi.GET("/user-all-by-uid", controller.GetUserAllEssayByUidHandler)  //获取根据uid用户全部文章
	essayUserLoginApi.POST("/update", controller.UpdateEssayHandler)                   //更新文章

	//周刊
	r.POST("/api/essay/update-weekly", middleware.MiddlewareUserLoginApi(), middleware.MiddlewareIsJoinCircleApi(),
		middleware.MiddlewareIsCircleOwnerApi(), controller.UpdateEssayWeeklyHandler) //将文章添加周刊
	r.GET("/api/essay/get-weekly", middleware.MiddlewareUserLoginApi(), controller.GetEssayWeeklylistHandler) //获取文章周刊
	//精粹
	r.POST("/api/essay/update-essence", middleware.MiddlewareUserLoginApi(), middleware.MiddlewareIsJoinCircleApi(),
		middleware.MiddlewareIsCircleOwnerApi(), controller.UpdateEssayEssenceHandler) //将文章添加精粹
	r.GET("/api/essay/get-essence", middleware.MiddlewareUserLoginApi(), controller.GetEssayEssonceHandler) //获取文章精粹

	//课程
	userCourseLoginPage := r.Group("/page/course")
	userCourseLoginPage.Use(middleware.MiddlewareUserLoginPage())
	//课程首页
	userCourseLoginPage.GET("/index", controller.CourseIndexPageHandler)
	//课程详情页面
	userCourseLoginPage.GET("/detail", controller.CourseDetailPageHandler)

	courseAuthorPage := r.Group("/page/course")
	courseAuthorPage.Use(middleware.MiddlewareUserLoginPage())
	//修改课程页面
	userCourseLoginPage.GET("/edit", controller.EditCoursePageHandler)

	userCourseLoginApi := r.Group("/api/course")
	userCourseLoginApi.Use(middleware.MiddlewareUserLoginApi())
	userCourseLoginApi.POST("/add", controller.AddCourseHandler)                 //添加课程
	userCourseLoginApi.GET("/all", controller.GetAllCourseHandler)               //获取全部课程
	userCourseLoginApi.GET("/user-all", controller.GetUserAllCourseByUidHandler) //获取用户发布的课程
	userCourseLoginApi.GET("/search", controller.GetCourseByTitleHandler)        //获取全部搜索记录
	userCourseLoginApi.GET("/api/course/get", controller.GetCourseHandler)       //获取课程详情

	//课时
	lessonOwnerApi := r.Group("/api/lesson")
	lessonOwnerApi.Use(middleware.MiddlewareUserLoginApi())
	lessonOwnerApi.POST("/add", controller.AddLessonHandler)         //添加课时
	lessonOwnerApi.GET("/get", controller.GetLessonHandler)          //获取课时详情
	lessonOwnerApi.GET("/all", controller.GetCourseAllLessonHandler) //获取课程全部课时

	lessonOwnerPage := r.Group("/page/lesson")
	lessonOwnerPage.Use(middleware.MiddlewareUserLoginPage())
	//创建课时页面
	lessonOwnerPage.GET("/add", controller.AddLessonPageHandler)
	//课时详情页面
	lessonOwnerPage.GET("/detail", controller.LessonDetailPageHandler)

	//买课
	userPurchaseLoginApi := r.Group("/api/purchase")
	userPurchaseLoginApi.Use(middleware.MiddlewareUserLoginApi())
	userPurchaseLoginApi.POST("/add", controller.AddPurchaseHandler)        //购买课程
	userPurchaseLoginApi.GET("/all", controller.GetUserPurchaseListHandler) //获取用户购买记录
	userPurchaseLoginApi.GET("/get", controller.GetPurchaseHandler)         //获取购买记录
	//TODO 支付，取消，退款
	userPurchaseLoginApi.POST("/update", controller.UpdatePurchaseStatusHandler) //更新购买记录状态购买课程

	//点赞
	userLikeLoginApi := r.Group("/api/like")
	userLikeLoginApi.Use(middleware.MiddlewareUserLoginApi())
	userLikeLoginApi.POST("/add", controller.AddUserEssayLikeHandler)          //添加点赞
	userLikeLoginApi.POST("/cancel", controller.CancelUserEssayLikeHandler)    //更新点赞删除
	userLikeLoginApi.GET("/get", controller.GetUserEssayLikeHandler)           //获取点赞
	userLikeLoginApi.GET("/all", controller.GetUserAllLikeHandler)             //获取用户点赞
	userLikeLoginApi.GET("/all-by-uid", controller.GetUserAllLikeByUidHandler) //获取根据uid用户点赞

	//收藏夹
	userFavoriteLoginApi := r.Group("/api/favorite")
	userFavoriteLoginApi.Use(middleware.MiddlewareUserLoginApi())
	userFavoriteLoginApi.POST("/add", controller.AddFavoriteHandler)                         //添加收藏夹
	userFavoriteLoginApi.POST("/update", controller.UpdateFavoriteTitleHandler)              //修改收藏夹名
	userFavoriteLoginApi.POST("/delete", controller.DeletedFavoriteByUpdateIsDeletedHandler) //删除收藏夹
	userFavoriteLoginApi.GET("/get", controller.GetFavoriteHandler)                          //获取收藏夹
	userFavoriteLoginApi.GET("/all", controller.GetUserAllFavoriteHandler)                   //获取用户全部收藏夹
	userFavoriteLoginApi.GET("/all-by-uid", controller.GetUserAllFavoriteByUidHandler)       //获取根据uid用户全部收藏夹

	//TODO 去除唯一键，新建时判断该数量=1，就不能新建

	//收藏
	userCollectLoginApi := r.Group("/api/collect")
	userCollectLoginApi.Use(middleware.MiddlewareUserLoginApi())
	userCollectLoginApi.POST("/add", controller.AddUserEssayCollectHandler)   //添加收藏
	userCollectLoginApi.POST("/cancel", controller.CancelEssayCollectHandler) //更新收藏删除状态
	userCollectLoginApi.GET("/get", controller.GetEssayCollectHandler)        //获取收藏
	userCollectLoginApi.GET("/all", controller.GetUserAllCollectHandler)      //获取用户收藏夹的全部收藏

	//评论
	userCommentLoginApi := r.Group("/api/comment")
	userCommentLoginApi.Use(middleware.MiddlewareUserLoginApi())
	userCommentLoginApi.POST("/add", controller.AddUserEssayCommentHandle)                 //创建用户评论
	userCommentLoginApi.POST("/delete", controller.DeletedCommentByUpdateIsDeletedHandler) //删除评论
	userCommentLoginApi.GET("/essayall", controller.GetEssayAllCommentHandle)              //获取文章全部评论
	userCommentLoginApi.GET("/userall", controller.GetUserAllCommentHandler)               //获取用户全部评论
	userCommentLoginApi.GET("/user-by-uid", controller.GetUserAllCommentByUidHandler)      //获取根据uid用户全部评论

	//关注
	userFollowLoginApi := r.Group("/api/follow")
	userFollowLoginApi.Use(middleware.MiddlewareUserLoginApi())
	userFollowLoginApi.POST("/add", controller.AddUserFollowHandler)       //添加关注
	userFollowLoginApi.POST("/cancel", controller.CancelUserFollowHandler) //更新关注删除
	userFollowLoginApi.GET("/get", controller.GetUserFollowHandler)        //获取用户关注
	userFollowLoginApi.GET("/all", controller.GetUserAllFollowHandler)     //获取用户关注列表
	userFollowLoginApi.GET("/allfan", controller.GetUserAllFanHandler)     //获取用户粉丝列表

	//举报
	r.GET("/page/accusation/edit", middleware.MiddlewareAdminUserLoginPage(), controller.AccusationEditPageHandler)
	r.POST("/api/accusation/update", middleware.MiddlewareAdminUserLoginApi(), controller.UpdateAccusationStatusHandler) //更新举报状态
	userAccusationLoginApi := r.Group("/api/accusation")
	userAccusationLoginApi.Use(middleware.MiddlewareUserLoginApi())
	userAccusationLoginApi.POST("/add", controller.AddUserAccusationEssayHandler)     //创建举报
	userAccusationLoginApi.GET("/all", controller.GetAllAccusationEssayHandler)       //获取全部未审核举报
	userAccusationLoginApi.GET("/get", controller.GetEssayContentByAccusationHandler) //获取举报内容文章

	//反馈
	userFeedbackLoginPage := r.Group("/page/feedback")
	userFeedbackLoginPage.Use(middleware.MiddlewareUserLoginPage())
	userFeedbackLoginPage.POST("/add", controller.AddUserFeedbackHandler) //创建反馈
	userFeedbackLoginPage.GET("/all", controller.GetAllFeedbackHandler)   //获取全部未处理反馈
	r.GET("/page/feedback/edit", middleware.MiddlewareAdminUserLoginPage(), controller.FeedbackEditPageHandler)

	userFeedbackLoginApi := r.Group("/api/feedback")
	userFeedbackLoginApi.Use(middleware.MiddlewareUserLoginApi())
	userFeedbackLoginApi.POST("/add", controller.AddUserFeedbackHandler)                                                   //创建反馈
	userFeedbackLoginApi.GET("/all", controller.GetAllFeedbackHandler)                                                     //获取全部未处理反馈
	userFeedbackLoginApi.GET("/get", controller.GetFeedbackContentHandler)                                                 //获取反馈
	userFeedbackLoginApi.GET("/get-by-uid", controller.GetFeedbackByUidHandler)                                            //获取用户反馈
	userFeedbackLoginApi.POST("/update", middleware.MiddlewareAdminUserLoginApi(), controller.UpdateFeedbackStatusHandler) //更新反馈状态

	//私信
	userChatLoginPage := r.Group("/page/chat")
	userChatLoginPage.Use(middleware.MiddlewareUserLoginPage())
	userChatLoginPage.POST("/add", controller.AddChatHandler)          //添加私信
	userChatLoginPage.GET("/detail", controller.ChatDetailPageHandler) //获取私信详情页面

	userChatLoginApi := r.Group("/api/chat")
	userChatLoginApi.Use(middleware.MiddlewareUserLoginApi())
	userChatLoginApi.POST("/add", controller.AddChatHandler)    //添加私信
	userChatLoginApi.GET("/get", controller.GetChatListHandler) //获取私信记录

	//联系人
	r.GET("/api/contact/all", middleware.MiddlewareUserLoginApi(), controller.GetChatContactListHandler) //获取最近联系人列表

	//通知
	userNoticeLoginPage := r.Group("/page/notice")
	userNoticeLoginPage.Use(middleware.MiddlewareUserLoginPage())
	userNoticeLoginPage.GET("/detail", controller.NoticeDetailPageHandler)
	//获取通知详情列表
	userNoticeLoginPage.GET("/index", controller.NoticeIndexPageHandler)

	userNoticeLoginApi := r.Group("/api/notice")
	userNoticeLoginApi.Use(middleware.MiddlewareUserLoginApi())
	userNoticeLoginApi.GET("/all", controller.GetNoticeListHandler)           //获取通知列表
	userNoticeLoginApi.GET("/alltype", controller.GetNoticeListByTypeHandler) //获取某类型通知列表

	//统计
	userStatLoginPage := r.Group("/page/stat")
	userStatLoginPage.Use(middleware.MiddlewareUserLoginPage())
	userStatLoginPage.GET("/index", controller.StatIndexPageHandler) //获取用户数据首页统计数据

	userStatLoginApi := r.Group("/api/stat")
	userStatLoginApi.Use(middleware.MiddlewareUserLoginApi())
	userStatLoginApi.GET("/by-time", controller.GetUserStatDetailsListByTimeHandler) //获取用户全部数据详情
	userStatLoginApi.GET("/map", controller.GetUserStatMapHandler)                   //获取用户某类数据数量
	userStatLoginApi.GET("/map-by-uid", controller.GetUserStatMapByUidHandler)       //获取用户某类数据数量

	//等级
	r.GET("/api/levelrecord/all", controller.GetUserCircleLevelAllRecordHandler) //获取用户在圈子全部等级详情

	//管理员用户
	adminPage := r.Group("/page/admin")
	adminPage.Use(middleware.MiddlewareAdminUserLoginPage())
	adminPage.GET("/index", controller.AdminIndexPageHandler)
	adminPage.GET("/edit", controller.AdminEditPageHandler)

	//角色权限
	adminRolePage := r.Group("/page/adminrole")
	adminRolePage.Use(middleware.MiddlewareAdminUserLoginPage())
	adminRolePage.GET("/edit", controller.AdminRoleEditPageHandler)

	r.POST("/api/adminuser/add", controller.AddAdminUserHandler)      //添加管理员用户信息
	r.POST("/api/adminuser/login", controller.AdminUserLoginHandler)  //管理员用户登录
	r.GET("/api/adminuser/logout", controller.AdminUserLogoutHandler) //管理员用户退出登录

	adminApi := r.Group("/api/adminuser")
	adminApi.Use(middleware.MiddlewareAdminUserLoginApi())
	adminApi.POST("/update", controller.UpdateAdminUserHandler)          //更新管理员用户信息
	adminApi.GET("/get", controller.GetAdminUserHandler)                 //获取某管理员用户信息
	adminApi.POST("/delete", controller.DeleteAdminUserHandler)          // 删除管理员用户
	adminApi.POST("/update-role", controller.UpdateAdminUserRoleHandler) //更新管理员用户角色

	//  菜单权限
	adminUserMenuLoginApi := r.Group("/api/adminmenu")
	adminUserMenuLoginApi.Use(middleware.MiddlewareAdminUserLoginApi())
	adminUserMenuLoginApi.GET("/all", controller.GetAllMenuHandler)     //获取全部菜单
	adminUserMenuLoginApi.POST("/add", controller.AddMenuHandler)       //添加菜单
	adminUserMenuLoginApi.POST("/delete", controller.DeleteMenuHandler) //删除菜单

	//  角色权限
	adminUserRoleLoginApi := r.Group("/api/adminrole")
	adminUserRoleLoginApi.Use(middleware.MiddlewareAdminUserLoginApi())
	adminUserRoleLoginApi.GET("/all", controller.GetAllRoleHandler)     //获取全部角色
	adminUserRoleLoginApi.GET("/get", controller.GetRoleHandler)        //获取角色详情
	adminUserRoleLoginApi.POST("/add", controller.AddRoleHandler)       //添加角色
	adminUserRoleLoginApi.POST("/update", controller.UpdateRoleHandler) //更新角色
	adminUserRoleLoginApi.POST("/delete", controller.DeleteRoleHandler) //删除角色

	//公告
	userAnnounceLoginPage := r.Group("/page/announce")
	userAnnounceLoginPage.Use(middleware.MiddlewareAdminUserLoginPage())
	userAnnounceLoginPage.GET("/edit", controller.AnnounceEditPageHandler)

	userAnnounceLoginApi := r.Group("/api/announce")
	userAnnounceLoginApi.Use(middleware.MiddlewareAdminUserLoginApi())
	userAnnounceLoginApi.POST("/add", controller.AddAnnounceHandler)                         //创建公告
	userAnnounceLoginApi.GET("/all-time", controller.GetAllAnnounceByTimeHandler)            //获取全部公告
	userAnnounceLoginApi.GET("/all", controller.GetAllAnnounceHandler)                       //获取全部公告
	userAnnounceLoginApi.GET("/get", controller.GetAnnounceHandler)                          //查看公告
	userAnnounceLoginApi.POST("/update", controller.UpdateAnnounceHandler)                   //更新公告
	userAnnounceLoginApi.POST("/delete", controller.DeletedAnnounceByUpdateIsDeletedHandler) //删除公告

	//广告
	userAdvertLoginPage := r.Group("/page/advert")
	userAdvertLoginPage.Use(middleware.MiddlewareAdminUserLoginPage())
	userAdvertLoginPage.GET("/edit", controller.AdvertEditPageHandler)

	adminAdvertLoginApi := r.Group("/api/advert")
	adminAdvertLoginApi.Use(middleware.MiddlewareAdminUserLoginApi())
	adminAdvertLoginApi.POST("/add", controller.AddAdvertHandler) //创建广告
	adminAdvertLoginApi.GET("/all", controller.GetAllAdvertHandler)
	adminAdvertLoginApi.GET("/get", controller.GetAdvertHandler)                          //查看广告
	adminAdvertLoginApi.POST("/update", controller.UpdateAdvertHandler)                   //更新广告
	adminAdvertLoginApi.POST("/delete", controller.DeletedAdvertByUpdateIsDeletedHandler) //删除广告

	r.GET("/all-time", middleware.MiddlewareUserLoginApi(), controller.GetAllAdvertByTimeHandler) //获取全部广告

	//搜索
	apiSearchLoginApi := r.Group("/api/search")
	apiSearchLoginApi.Use(middleware.MiddlewareUserLoginApi())
	apiSearchLoginApi.GET("/circle", controller.GetCircleByTitleHandler) //搜索圈子
	apiSearchLoginApi.GET("/essay", controller.GetEssayByTitleHandler)   //搜索文章

	//支付

	r.GET("/page/orders/index", middleware.MiddlewareUserLoginPage(), controller.OrdersIndexPageHandler) //获取用户订单首页列表

	userOrdersLoginApi := r.Group("/api/orders")
	userOrdersLoginApi.Use(middleware.MiddlewareUserLoginApi())
	userOrdersLoginApi.POST("/add", controller.AddOrdersHandler)       //创建支付
	userOrdersLoginApi.GET("/all", controller.GetUserAllOrdersHandler) //获取用户全部支付
	userOrdersLoginApi.GET("/get", controller.GetOrdersHandler)        //查看支付
	//TODO 支付，取消，退款
	userOrdersLoginApi.POST("/update", controller.UpdateUserOrdersHandler)      //用户支付更新
	userOrdersLoginApi.GET("/getorders", controller.GetUserOrdersCircleHandler) //获取需要支付

	userContactLoginApi := r.Group("/api/contact")
	userContactLoginApi.Use(middleware.MiddlewareUserLoginApi())
	userContactLoginApi.POST("/add", controller.AddUserContactHandler)       //添加联系人
	userContactLoginApi.POST("/delete", controller.DeleteUserContactHandler) //删除联系人
	//userContactLoginApi.GET("/all", controller.GetUserAllContactHandler)     //获取用户全部联系人
	userContactLoginApi.GET("/get", controller.GetUserContactHandler) //获取联系人

	// 启动服务器
	service.Logger.Info("The server started at port", zap.String("port", "8080"))
	service.Logger.Error("Default error", zap.Error(r.Run(":8080")))
}

// 全局模板
var tpl *template.Template

func init() {
	// 根目录
	root := "./web/views"
	tpl = template.New("")

	// 遍历所有 HTML
	filepath.Walk(root, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() && filepath.Ext(path) == ".html" {
			// --------------------------
			// 🔥 核心：强制模板名 = user/edit.html
			// --------------------------
			rel, _ := filepath.Rel(root, path)
			name := filepath.ToSlash(rel)

			// 读文件内容
			content, _ := os.ReadFile(path)

			// 解析
			template.Must(tpl.New(name).Parse(string(content)))
		}
		return nil
	})
}
