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
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	followIdStr := c.Query("followId")
	if followIdStr == "" {
		service.Logger.Error("GetfollowId err", zap.String("err", "get followId err"))
		MakeApiResponseErrorParams(c)
		return
	}

	followId, err := strconv.Atoi(followIdStr)
	if err != nil {
		service.Logger.Error("Atoi followIdStr err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//查询用户的关注
	follow, err := service.GetUserFollow(uid, followId)
	if err != nil {
		service.Logger.Error("GetUserFollow", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//仅 不存在，存在状态为删除 两种
	if follow == nil {
		createTime := time.Now()

		newFollow := &model.Follow{ //其中包含自动生成的id
			FanId:        uid,
			FollowerId:   followId,
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

		MakeApiResponseSuccessDefault(c)
		return
	} else {

		//未关注状态转关注
		affectRows, err := service.UpdateUserFollowNotToIs(uid, followId)
		if err != nil || affectRows == 0 {
			service.Logger.Error("UpdateUserFollowNotToIs err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		MakeApiResponseSuccessDefault(c)
		return
	}
}

// 取消关注
func CancelUserFollowHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	followerIdStr := c.Query("followerId")
	if followerIdStr == "" {
		service.Logger.Error("GetfollowerId err", zap.String("err", "get followerId err"))
		MakeApiResponseErrorParams(c)
		return
	}

	followerId, err := strconv.Atoi(followerIdStr)
	if err != nil {
		service.Logger.Error("Atoi followerIdStr err", zap.Error(err))
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

// 用户关注列表
func GetUserAllFollowHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	page := c.GetInt("page")
	if page < 1 {
		page = 1
	}

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

	page := c.GetInt("page")
	if page < 1 {
		page = 1
	}

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
