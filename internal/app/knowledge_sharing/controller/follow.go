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
	uid := service.GetUidFromContext(c)

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

	//从context获取用户名
	userName := service.GetNameFromContext(c)

	//查询用户的关注
	follow, err := service.GetUserFollow(uid, followerId)
	if err != nil {
		service.Logger.Error("GetUserFollow", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	nowTime := time.Now()

	//仅 不存在，存在状态为删除 两种
	if follow == nil {

		newFollow := &model.Follow{ //其中包含自动生成的id
			FanId:        uid,
			FollowerId:   followerId,
			FollowTime:   &nowTime,
			CreateAt:     &nowTime,
			UpdateAt:     &nowTime,
			FollowStatus: model.FOLLOW_STATUS_NORMAL,
		}

		err = service.CreateUserFollow(newFollow)
		if err != nil {
			service.Logger.Error("CreateUserFollow err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		typei := model.NOTICE_TYPE_FOLLOW
		//TODO 异步处理

		noticeMsg := &model.NoticeMsg{
			Type:     typei,
			Uid:      followerId,
			Time:     nowTime.Unix(),
			UserName: userName,
		}

		_, _, err := service.ProduceKafkaNoticeMessage(noticeMsg)
		if err != nil {
			service.Logger.Error("ProduceKafkaNoticeMessage err", zap.Error(err))
			err = nil
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
	uid := service.GetUidFromContext(c)

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
	uid := service.GetUidFromContext(c)

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
	uid := service.GetUidFromContext(c)

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
	uid := service.GetUidFromContext(c)

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
