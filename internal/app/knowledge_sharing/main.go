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
	defer service.Closekafka()

	//创建 Gin 路由引擎
	r := gin.Default()
	// r.SetTrustedProxies(nil)
	// 给 Gin 设置自定义模板引擎
	r.SetHTMLTemplate(tpl)
	r.Static("/img", "web/img")

	//页面路由
	page := r.Group("/page")
	page.Use(middleware.PageUserLogin())
	api := r.Group("/api")
	api.Use(middleware.ApiUserLogin())

	r.GET("/", controller.IndexPageHandler)

	//注册页面
	r.GET("/page/user/register", controller.RegisterPageHandler)
	r.POST("/api/user/add", controller.AddUserHandler)
	//用户登录页面
	r.GET("/page/user/login", controller.LoginPageHandler)
	r.POST("/api/user/login", controller.UserLoginHandler)

	//用户退出登录
	api.GET("/user/logout", controller.UserLogoutHandler)

	//更新用户
	page.GET("/user/edit", controller.EditPageHandler)
	//更新用户信息
	api.POST("/user/update", controller.UpdateUserHandler)

	//更新用户密码
	page.GET("/user/edit-password", controller.EditPasswordPageHandler)
	api.POST("/user/updatepassword", controller.UpdateUserPasswordHandler)

	page.GET("/user/profile", controller.ProfilePageHandler)
	//获取某用户信息
	api.GET("/user/get", controller.GetUserHandler)
	//通过uid获取用户信息
	api.GET("/user/get-by-id", controller.GetUserByIdHandler)
	//获取用户喜欢的文章
	api.GET("/like/all", controller.GetUserAllLikeHandler)
	api.GET("/like/all-by-uid", controller.GetUserAllLikeByUidHandler)
	//获取用户全部收藏夹
	api.GET("/favorite/all", controller.GetUserAllFavoriteHandler)
	//获取根据uid用户全部收藏夹
	api.GET("/favorite/all-by-uid", controller.GetUserAllFavoriteByUidHandler)
	//用户收藏的文章列表
	api.GET("/collect/all", controller.GetUserAllCollectHandler)
	//获取用户全部评论
	api.GET("/comment/userall", controller.GetUserAllCommentHandler)
	//获取根据uid用户全部评论
	api.GET("/comment/user-by-uid", controller.GetUserAllCommentByUidHandler)
	//添加关注
	api.POST("/follow/add", controller.AddUserFollowHandler)
	//更新关注删除
	api.POST("/follow/cancel", controller.CancelUserFollowHandler)
	//获取用户关注
	api.GET("/follow/get", controller.GetUserFollowHandler)
	//获取用户关注列表
	api.GET("/follow/all", controller.GetUserAllFollowHandler)
	//获取用户粉丝列表
	api.GET("/follow/allfan", controller.GetUserAllFanHandler)

	//圈子模块
	//创建圈子页面
	page.GET("/circle/add", controller.AddCirclePageHandler)
	//创建圈子
	api.POST("/circle/add", controller.AddCircleHandler)

	//修改圈子页面
	page.GET("/circle/edit",
		middleware.PageIsJoinCircle(),
		middleware.PageIsCircleOwner(),
		controller.EditCirclePageHandler)
	//更新圈子信息
	api.POST("/circle/update",
		middleware.ApiIsJoinCircle(),
		middleware.ApiIsCircleOwner(),
		controller.UpdateCircleHandler)

	//获取圈子详情页面
	page.GET("/circle/detail", controller.CircleDetailPageHandler)
	//获取圈子详情
	api.GET("/circle/get", controller.GetCircleHandler)
	//获取是否需要续费
	api.GET("/circle/need-new", controller.GetUserOrdersCircleHandler)
	//用户加入圈子
	api.POST("/usercircle/add", controller.AddUserCircleJoinHandle)
	//创建圈子订单
	api.POST("/orders/add", controller.AddOrdersHandler)
	//更新订单状态
	api.POST("/orders/update", controller.UpdateUserOrdersHandler)

	//用户退出圈子
	api.POST("/usercircle/quit",
		middleware.ApiIsJoinCircle(),
		controller.QuitCircleHandler)

	//加入的圈子页面
	page.GET("/circle/index", controller.CircleIndexPageHandler)
	//获取用户加入的圈子
	api.GET("/circle/join", controller.GetUserJoinCircleHandler)

	//创建的圈子页面
	page.GET("/circle/create", controller.CircleCreatePageHandler)
	//获取用户创建的圈子
	api.GET("/circle/create", controller.GetUserCreateCircleHandler)

	//圈子列表（排行榜更多）
	//TODO 展示广告
	r.GET("/page/circle/list", controller.CircleListPageHandler)
	//获取全部圈子
	api.GET("/circle/all", controller.GetAllCircleHandler)
	//获取全部付费圈子
	api.GET("/circle/charge", controller.GetChargeCircleRankHandler)
	//获取全部免费圈子
	api.GET("/circle/free", controller.GetFreeCircleRankHandler)

	//搜索圈子
	api.GET("/search/circle", controller.GetCircleByTitleHandler)

	//获取全部广告
	api.GET("/advert/all-time", controller.GetAllAdvertByTimeHandler)

	//文章模块
	//创建文章页面
	page.GET("/essay/add", controller.AddEssayPageHandler)
	api.POST("/essay/add", controller.AddEssayHandler)

	page.GET("/essay/edit", controller.EditEssayPageHandler)
	//更新文章
	api.POST("/essay/update", controller.UpdateEssayHandler)
	//删除文章
	api.POST("/essay/delete", controller.DeletedEssayByUpdateIsDeletedHandler)
	//将文章添加周刊
	api.POST("/essay/update-weekly",
		middleware.ApiIsJoinCircle(),
		middleware.ApiIsCircleOwner(),
		controller.UpdateEssayWeeklyHandler)
	//将文章添加精粹
	api.POST("/essay/update-essence",
		middleware.ApiIsJoinCircle(),
		middleware.ApiIsCircleOwner(),
		controller.UpdateEssayEssenceHandler)

	//查看文章详情页面
	page.GET("/essay/detail", controller.EssayDetailPageHandler)
	//查看文章
	api.GET("/essay/get", controller.GetEssayHandler)
	//获取文章全部评论
	api.GET("/comment/essayall", controller.GetEssayAllCommentHandle)

	//搜索文章
	page.GET("/essay/search", controller.SearchEssayPageHandler)
	api.GET("/search/essay", controller.GetEssayByTitleHandler)
	//获取圈子全部文章 /page/circle/index
	api.GET("/essay/circle-all", controller.GetCircleAllEssayHandler)
	//获取周刊文章
	api.GET("/essay/get-weekly", controller.GetEssayWeeklylistHandler)
	//获取精粹文章
	api.GET("/essay/get-essence", controller.GetEssayEssonceHandler)

	//获取用户全部文章 /page/user/profile
	api.GET("/essay/user-all", controller.GetUserAllEssayHandler)
	//获取根据uid用户全部文章 /page/user/profile?uid=?
	api.GET("/essay/user-all-by-uid", controller.GetUserAllEssayByUidHandler)

	//点赞 /page/essay/detail   /page/circle/index
	api.POST("/like/add", controller.AddUserEssayLikeHandler)
	api.POST("/like/cancel", controller.CancelUserEssayLikeHandler)
	api.GET("/like/get", controller.GetUserEssayLikeHandler)

	//收藏夹 /page/circle/index /page/essay/detail
	//添加收藏夹
	api.POST("/favorite/add", controller.AddFavoriteHandler)
	//修改收藏夹名
	api.POST("/favorite/update", controller.UpdateFavoriteTitleHandler)
	//删除收藏夹
	api.POST("/favorite/delete", controller.DeletedFavoriteByUpdateIsDeletedHandler)
	//获取收藏夹
	api.GET("/favorite/get", controller.GetFavoriteHandler)

	//收藏
	//添加收藏 /page/circle/index /page/essay/detail
	api.POST("/collect/add", controller.AddUserEssayCollectHandler)
	//更新收藏删除状态
	api.POST("/collect/cancel", controller.CancelEssayCollectHandler)
	//获取文章是否收藏
	api.GET("/collect/get", controller.GetEssayCollectHandler)

	//评论
	//创建用户评论 /page/circle/index /page/essay/detail
	api.POST("/comment/add", controller.AddUserEssayCommentHandle)
	//删除评论
	api.POST("/comment/delete", controller.DeletedCommentByUpdateIsDeletedHandler)

	//举报 /page/circle/index /page/essay/detail
	//创建举报
	api.POST("/accusation/add", controller.AddUserAccusationEssayHandler)

	//私信
	//获取私信首页
	page.GET("/chat/index", controller.ChatIndexPageHandler)
	//获取最近联系人列表
	api.GET("/contact/all", controller.GetChatContactListHandler)
	//api.POST("/contact/add", controller.AddUserContactHandler)       //添加联系人
	//api.POST("/contact/delete", controller.DeleteUserContactHandler) //删除联系人
	//api.GET("/contact/get", controller.GetUserContactHandler)        //获取联系人
	// page.GET("/chat/detail", controller.ChatDetailPageHandler)    //获取私信详情页面
	//添加私信
	api.POST("/chat/add", controller.AddChatHandler)
	//获取私信记录
	api.GET("/chat/get", controller.GetChatListHandler)

	//反馈
	//反馈首页
	page.GET("/feedback/index", controller.FeedbackIndexPageHandler)
	//创建反馈
	api.POST("/feedback/add", controller.AddUserFeedbackHandler)
	//获取用户反馈
	api.GET("/feedback/get-by-uid", controller.GetFeedbackByUidHandler)
	//获取反馈详情
	page.GET("/feedback/detail", controller.FeedbackDetailPageHandler)
	//获取反馈
	api.GET("/feedback/get", controller.GetFeedbackContentHandler)

	//通知
	page.GET("/notice/index", controller.NoticeIndexPageHandler)
	//获取通知列表
	api.GET("/notice/all", controller.GetNoticeListHandler)
	//获取某类型通知列表
	api.GET("/notice/alltype", controller.GetNoticeListByTypeHandler)
	//page.GET("/notice/detail", controller.NoticeDetailPageHandler)

	//统计
	page.GET("/stat/index", controller.StatIndexPageHandler)
	//获取用户全部数据详情
	api.GET("/stat/by-time", controller.GetUserStatDetailsListByTimeHandler)
	//获取用户某类数据数量
	api.GET("/stat/map", controller.GetUserStatMapHandler)
	//获取用户某类数据数量
	api.GET("/stat/map-by-uid", controller.GetUserStatMapByUidHandler)

	//圈子订单
	page.GET("/orders/index", controller.OrdersIndexPageHandler)
	//获取用户全部支付
	api.GET("/orders/all", controller.GetUserAllOrdersHandler)
	//查看支付
	api.GET("/orders/get", controller.GetOrdersHandler)

	//等级
	//api.GET("/levelrecord/all", controller.GetUserCircleLevelAllRecordHandler)

	//课程
	page.GET("/course/index", controller.CourseIndexPageHandler)
	//获取用户发布的课程
	api.GET("/course/user-all", controller.GetUserAllCourseByUidHandler)
	//获取全部课程
	api.GET("/course/all", controller.GetAllCourseHandler)
	//获取全部搜索记录
	api.GET("/course/search", controller.GetCourseByTitleHandler)
	//获取用户购买记录
	api.GET("/purchase/all", controller.GetUserPurchaseListHandler)

	//添加课程页面
	page.GET("/course/add", controller.AddCoursePageHandler)
	//添加课程
	api.POST("/course/add", controller.AddCourseHandler)

	//课程详情页面
	page.GET("/course/detail", controller.CourseDetailPageHandler)
	//获取课程详情
	api.GET("/course/get", controller.GetCourseHandler)
	//获取购买记录
	api.GET("/purchase/get", controller.GetPurchaseHandler)
	//购买课程
	api.POST("/purchase/add", controller.AddPurchaseHandler)
	//更新购买记录状态购买课程
	api.POST("/purchase/pay", controller.UpdatePurchaseStatusHandler)

	//修改课程页面
	page.GET("/course/edit", controller.EditCoursePageHandler)
	//修改课程
	api.POST("/course/update", controller.UpdateCourseHandler)

	//课时
	//课时页面
	page.GET("/lesson/add", controller.AddLessonPageHandler)
	//添加课时
	api.POST("/lesson/add", controller.AddLessonHandler)

	//课时详情 页面
	page.GET("/lesson/detail", controller.LessonDetailPageHandler)
	//获取课程全部课时
	api.GET("/lesson/all", controller.GetCourseAllLessonHandler)
	//获取课时详情
	api.GET("/lesson/get", controller.GetLessonHandler)
	// 作者修改课时
	page.GET("/lesson/edit", controller.EditLessonPageHandler)
	//修改课时
	api.POST("/lesson/update", controller.UpdateLessonHandler)

	//公告
	//公告列表页
	page.GET("/announce/list", controller.AnnounceListPageHandler)
	//获取全部公告
	api.GET("/announce/all", controller.GetAllAnnounceHandler)
	//公告详情页
	page.GET("/announce/detail", controller.AnnounceDetailPageHandler)
	//查看公告
	api.GET("/announce/get", controller.GetAnnounceHandler)

	//管理员用户登录
	r.POST("/api/adminuser/login", controller.AdminUserLoginHandler)
	//管理员用户退出登录
	r.GET("/api/adminuser/logout", controller.AdminUserLogoutHandler)

	//管理员用户
	// admin/page   admin/api
	adminPage := r.Group("/page")
	adminPage.Use(middleware.PageAdminUserLogin())
	adminApi := r.Group("/api")
	adminApi.Use(middleware.ApiAdminUserLogin())

	adminPage.GET("/admin/index", controller.AdminIndexPageHandler)

	// 添加菜单页
	adminPage.GET("/adminmenu/add", controller.MenuAddPageHandler)
	//添加菜单
	adminApi.POST("/adminmenu/add", controller.AddMenuHandler)

	// 修改菜单页
	adminPage.GET("/adminmenu/edit", controller.MenuEditPageHandler)
	//修改菜单
	adminApi.POST("/adminmenu/update", controller.UpdateMenuHandler)

	// 菜单列表页
	adminPage.GET("/adminmenu/list", controller.MenuListPageHandler)
	adminApi.GET("/adminmenu/all", controller.GetAllMenuHandler)
	//删除菜单
	adminApi.POST("/adminmenu/delete", controller.DeleteMenuHandler)

	//角色权限
	//添加角色页
	adminPage.GET("/adminrole/add", controller.RoleAddPageHandler)
	//添加角色
	adminApi.POST("/adminrole/add", controller.AddRoleHandler)

	//修改角色页
	adminPage.GET("/adminrole/edit", controller.AdminRoleEditPageHandler)
	//更新角色
	adminApi.POST("/adminrole/update", controller.UpdateRoleHandler)

	// 角色列表页
	adminPage.GET("/adminrole/list", controller.RoleListPageHandler)
	//获取全部角色
	adminApi.GET("/adminrole/all", controller.GetAllRoleHandler)
	//删除角色
	adminApi.POST("/adminrole/delete", controller.DeleteRoleHandler)
	//角色详情页
	adminPage.GET("/adminrole/detail", controller.RoleDetailPageHandler)
	//获取角色详情
	adminApi.GET("/adminrole/get", controller.GetRoleHandler)

	// 添加管理员页
	adminPage.GET("/adminuser/add", controller.AdminUserAddPageHandler)
	adminApi.POST("/adminuser/add", controller.AddAdminUserHandler)

	// 更新管理员页
	adminPage.GET("/adminuser/edit", controller.AdminEditPageHandler)
	//更新管理员用户信息
	adminApi.POST("/adminuser/update", controller.UpdateAdminUserHandler)
	//修改管理员用户角色页
	adminPage.GET("/adminuser/role", controller.AdminUserRolePageHandler)
	//更新管理员用户角色
	adminApi.POST("/adminuser/update-role", controller.UpdateAdminUserRoleByUidHandler)

	// 管理员列表页
	adminPage.GET("/adminuser/list", controller.AdminUserListPageHandler)
	//获取全部管理员用户
	adminApi.GET("/adminuser/all", controller.GetAllAdminUserHandler)
	// 删除管理员用户
	adminApi.POST("/adminuser/delete", controller.DeleteAdminUserHandler)
	//查看管理员详情页
	adminPage.GET("/adminuser/detail", controller.AdminUserDetailPageHandler)
	//获取某管理员用户信息
	adminApi.GET("/adminuser/get", controller.GetAdminUserHandler)

	//举报列表页
	adminPage.GET("/adminaccusation/list", controller.AdminAccusationListPageHandler)
	//获取全部未审核举报
	adminApi.GET("/adminaccusation/all", controller.GetAllAccusationEssayHandler)

	//修改举报页
	adminPage.GET("/adminaccusation/edit", controller.AccusationEditPageHandler)
	//获取举报内容文章
	adminApi.GET("/adminaccusation/get", controller.GetEssayContentByAccusationHandler)
	//更新举报状态
	adminApi.POST("/adminaccusation/update", controller.UpdateAccusationStatusHandler)

	//反馈列表页
	adminPage.GET("/adminfeedback/list", controller.AdminFeedbackListPageHandler)
	//获取全部反馈
	adminApi.GET("/adminfeedback/all", controller.GetAllFeedbackHandler)
	//处理反馈页
	adminPage.GET("/adminfeedback/edit", controller.FeedbackEditPageHandler)
	//获取反馈
	adminApi.GET("/adminfeedback/get", controller.GetFeedbackContentHandler)
	//更新反馈状态
	adminApi.POST("/adminfeedback/update", controller.UpdateFeedbackStatusHandler)

	//公告
	//添加公告页
	adminPage.GET("/adminannounce/add", controller.AdminAnnounceAddPageHandler)
	//创建公告
	adminApi.POST("/adminannounce/add", controller.AddAnnounceHandler)
	//公告列表页
	adminPage.GET("/adminannounce/list", controller.AdminAnnounceListPageHandler)
	//adminApi.GET("/announce/all-time", controller.GetAllAnnounceByTimeHandler) //获取全部公告
	//获取全部公告
	adminApi.GET("/adminannounce/all", controller.GetAllAnnounceHandler)
	//修改公告页
	adminPage.GET("/adminannounce/edit", controller.AnnounceEditPageHandler)
	//查看公告
	adminApi.GET("/adminannounce/get", controller.GetAnnounceHandler)
	//更新公告
	adminApi.POST("/adminannounce/update", controller.UpdateAnnounceHandler)
	//删除公告
	adminApi.POST("/adminannounce/delete", controller.DeletedAnnounceByUpdateIsDeletedHandler)

	//广告
	//添加广告页
	adminPage.GET("/adminadvert/add", controller.AdminAdvertAddPageHandler)
	//创建广告
	adminApi.POST("/adminadvert/add", controller.AddAdvertHandler)
	//广告列表页
	adminPage.GET("/adminadvert/list", controller.AdminAdvertListPageHandler)
	adminApi.GET("/adminadvert/all", controller.GetAllAdvertHandler)
	//修改广告页
	adminPage.GET("/adminadvert/edit", controller.AdvertEditPageHandler)
	//查看广告
	adminApi.GET("/adminadvert/get", controller.GetAdvertHandler)
	//更新广告
	adminApi.POST("/adminadvert/update", controller.UpdateAdvertHandler)
	//删除广告
	adminApi.POST("/adminadvert/delete", controller.DeletedAdvertByUpdateIsDeletedHandler)

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

	// 遍历所有HTML
	filepath.Walk(root, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() && filepath.Ext(path) == ".html" {
			// --------------------------
			// 核心：强制模板名 = user/edit.html
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
