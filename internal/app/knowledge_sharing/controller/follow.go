package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 添加关注
func AddUserFollowHandler(c *gin.Context) {
	uid, userName := service.GetUserFromCookie(c)
	if uid == 0 || userName == "" {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	followerIdStr := c.PostForm("followerId")
	if followerIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	followerId, err := strconv.Atoi(followerIdStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	//查询用户的关注
	follow, err := service.GetUserFollow(uid, followerId)
	if err != nil {
		service.Logger.Error("GetUserFollow", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	createTime := time.Now()

	//仅 不存在，存在状态为删除 两种
	if follow == nil {

		newFollow := &model.Follow{ //其中包含自动生成的id
			FanId:        uid,
			FollowerId:   followerId,
			FollowTime:   &createTime,
			CreateAt:     &createTime,
			UpdateAt:     &createTime,
			FollowStatus: model.FOLLOW_STATUS_NORMAL,
		}

		err = service.CreateUserFollow(newFollow)
		if err != nil {
			service.Logger.Error("CreateUserFollow err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		content := "又有新用户" + userName + "关注啦"
		typei := model.STAT_TYPE_FOLLOW

		//添加通知
		err = service.UserAddNotice(followerId, content, typei, createTime)
		if err != nil {
			service.Logger.Error("UserAddNotice err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

	} else {

		//未关注状态转关注
		affectRows, err := service.UpdateUserFollowNotToIs(uid, followerId)
		if err != nil || affectRows == 0 {
			service.Logger.Error("UpdateUserFollowNotToIs err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

	}

	nowTime := time.Now()

	//更新关注数据总数和详情
	err = service.UpdateStatAndStatDetail(uid, model.STAT_TYPE_FOLLOW, model.STAT_DETAILS_STATUS_INCR, nowTime)
	if err != nil {
		service.Logger.Error("UpdateStatAndStatDetail err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//更新被关注数据总数和详情
	err = service.UpdateStatAndStatDetail(followerId, model.STAT_TYPE_FAN, model.STAT_DETAILS_STATUS_INCR, nowTime)
	if err != nil {
		service.Logger.Error("UpdateStatAndStatDetail err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 取消关注
func CancelUserFollowHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	followerIdStr := c.PostForm("followerId")
	if followerIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	followerId, err := strconv.Atoi(followerIdStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	//取关
	affectRows, err := service.UpdateUserFollowIsToNot(uid, followerId)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateUserFollowIsToNot err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 用户关注状态
func GetUserFollowHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	followerIdStr := c.Query("followerId")
	if followerIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	followerId, err := strconv.Atoi(followerIdStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	//查询用户的关注
	follow, err := service.GetUserFollowByStatus(uid, followerId)
	if err != nil {
		service.Logger.Error("GetUserFollow", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	var isFollowed bool

	if follow != nil {
		isFollowed = true
	}

	data := map[string]bool{
		"isFollowed": isFollowed,
	}

	MakeApiResponseSuccess(c, data)
}

// 用户关注列表
func GetUserAllFollowHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pageSize := 10

	//获取用户follow列表
	users, err := service.GetUserFollowListByUid(uid, page, pageSize)
	if err != nil {
		service.Logger.Error("GetUserFollowListByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if users == nil {
		users = make([]model.User, 0)
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"user": users,
	})
}

// 用户粉丝列表
func GetUserAllFanHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pageSize := 10

	//获取用户被关注列表
	users, err := service.GetUserFanListByUid(uid, page, pageSize)
	if err != nil {
		service.Logger.Error("GetUserFanListByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if users == nil {
		users = make([]model.User, 0)
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"user": users,
	})
}
