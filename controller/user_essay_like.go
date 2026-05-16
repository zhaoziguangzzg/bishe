package controller

import (
	"bishe/model"
	"bishe/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 用户在文章的喜欢
func AddUserEssayLikeHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	eidStr := c.PostForm("eid")
	if eidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	lockKey := "user-add-like" + strconv.Itoa(uid) + "-" + strconv.Itoa(eid)
	lockValue, locked, err := service.Lock(c, lockKey, 5*time.Second)
	if err != nil {
		service.Logger.Error("Lock err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if !locked {
		MakeApiResponseError(c, CODE_LOCKED)
		return
	}

	defer service.Unlock(c, lockKey, lockValue)

	//根据eid获取文章
	essay, err := service.GetEssayByEid(eid)
	if err != nil {
		service.Logger.Error("GetEssayByEid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if essay == nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	authorId := essay.AuthorId
	cid := essay.CircleId

	var relateId int

	//查询对文章的点赞
	like, err := service.GetUserEssayLike(uid, eid)
	if err != nil {
		service.Logger.Error("GetUserEssayLike", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	nowTime := time.Now()

	//仅 不存在，存在状态为删除 两种
	if like == nil {

		newUserEssayLike := &model.UserEssayLike{ //其中包含自动生成的id
			UserId:     uid,
			EssayId:    eid,
			CreateAt:   &nowTime,
			UpdateAt:   &nowTime,
			LikeStatus: model.LIKE_STATUS_NORMAL,
		}

		err = service.CreateUserEssayLike(newUserEssayLike)
		if err != nil {
			service.Logger.Error("CreateUserEssayLike err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		typei := model.NOTICE_TYPE_LIKE

		noticeMsg := &model.NoticeMsg{
			Type:    typei,
			Time:    nowTime.Unix(),
			EssayId: eid,
			LikeUid: uid,
		}

		_, _, err := service.ProduceKafkaNoticeMessage(noticeMsg)
		if err != nil {
			service.Logger.Error("ProduceKafkaNoticeMessage err", zap.Error(err))
			err = nil
		}

		relateId = newUserEssayLike.Id
		score := 2
		//点赞 更新等级分数 增加2
		affectRows, err := service.UpdateLevelScoreByUidCid(uid, cid, score)
		if err != nil || affectRows == 0 {
			service.Logger.Error("UpdateLevelScoreByUidCid err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		//增加等级详情
		err = service.UserAddLevelScoreRecord(uid, cid, score, relateId, model.LEVEL_SCORE_RECORD_TYPE_LIKE, nowTime)
		if err != nil {
			service.Logger.Error("UserAddLevelScoreRecord err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		score = 3
		//被点赞 更新等级分数 增加3
		affectRows, err = service.UpdateLevelScoreByUidCid(authorId, cid, score)
		if err != nil || affectRows == 0 {
			service.Logger.Error("UpdateLevelScoreByUidCid err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		//增加等级详情
		err = service.UserAddLevelScoreRecord(authorId, cid, score, relateId, model.LEVEL_SCORE_RECORD_TYPE_LIKED, nowTime)
		if err != nil {
			service.Logger.Error("UserAddLevelScoreRecord err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

	} else {

		//不喜欢转喜欢
		affectRows, err := service.UpdateUserEssayLikeNotToIs(uid, eid)
		if err != nil || affectRows == 0 {
			service.Logger.Error("UpdateUserEssayLikeNotToIs err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

	}

	//更新被点赞数据总数和详情
	err = service.UpdateStatAndStatDetail(essay.AuthorId, model.STAT_TYPE_LIKED, model.STAT_DETAILS_STATUS_INCR, nowTime)
	if err != nil {
		service.Logger.Error("UpdateStatAndStatDetail err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//更新文章点赞数
	affectRows, err := service.UpdateEssayLikeNum(eid, 1)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateEssayLikeNum err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 取消用户喜欢
func CancelUserEssayLikeHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	eidStr := c.PostForm("eid")
	if eidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	//喜欢转不喜欢
	affectRows, err := service.UpdateUserEssayLikeIsToNot(uid, eid)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateUserEssayLikeIsToNot err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//更新文章点赞数
	affectRows, err = service.UpdateEssayLikeNum(eid, -1)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateEssayLikeNum err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 获取用户文章是否喜欢
func GetUserEssayLikeHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	eidStr := c.Query("eid")
	if eidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	like, err := service.GetUserEssayLike(uid, eid)
	if err != nil {
		service.Logger.Error("GetUserEssayLike", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	isLike := false
	if like != nil && like.LikeStatus == model.LIKE_STATUS_NORMAL {
		isLike = true
	}

	data := map[string]interface{}{
		"isLike": isLike,
	}

	MakeApiResponseSuccess(c, data)
}

// 获取用户全部喜欢
func GetUserAllLikeHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pageSize := 10

	//获取用户全部like
	userEssays, err := service.GetUserAllLikeEssayByUid(uid, page, pageSize)
	if err != nil {
		service.Logger.Error("GetUserAllLikeEssayByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(userEssays) == 0 {
		userEssays = make([]model.UserEssay, 0)
	}

	data := map[string]interface{}{
		"essays": userEssays,
	}

	MakeApiResponseSuccess(c, data)
}

// 获取根据uid用户全部喜欢
func GetUserAllLikeByUidHandler(c *gin.Context) {
	uidStr := c.Query("uid")
	if uidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pageSize := 10

	//获取用户全部like
	userEssays, err := service.GetUserAllLikeEssayByUid(uid, page, pageSize)
	if err != nil {
		service.Logger.Error("GetUserAllLikeEssayByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(userEssays) == 0 {
		userEssays = make([]model.UserEssay, 0)
	}

	data := map[string]interface{}{
		"essays": userEssays,
	}

	MakeApiResponseSuccess(c, data)
}

// func UpdateUserEssayInteract(c *gin.Context) {

// 	uidStr := c.PostForm("user_id")
// 	circleIdStr := c.PostForm("circle_id")
// 	essayIdStr := c.PostForm("essay_id")
// 	favorite := c.PostForm("favorite")

// 	if uidStr == "" {
// 		MakeApiResponseError(c, CODE_PARAMS_ERROR)
// 		return
// 	}

// 	if circleIdStr == "" {
// 		MakeApiResponseError(c, CODE_PARAMS_ERROR)
// 		return
// 	}

// 	if essayIdStr == "" {
// 		MakeApiResponseError(c, CODE_PARAMS_ERROR)
// 		return
// 	}

// 	uid, err := strconv.Atoi(uidStr)
// 	if err != nil {
// 		MakeApiResponseError(c, CODE_PARAMS_ERROR)
// 		return
// 	}

// 	circleId, err := strconv.Atoi(circleIdStr)
// 	if err != nil {
// 		MakeApiResponseError(c, CODE_PARAMS_ERROR)
// 		return
// 	}

// 	essayId, err := strconv.Atoi(essayIdStr)
// 	if err != nil {
// 		MakeApiResponseError(c, CODE_PARAMS_ERROR)
// 		return
// 	}

// 	userEssayInteract, err := service.GetUserEssayInteract(uid, circleId, essayId)
// 	if err != nil {
// 		service.Logger.Error("GetUserEssayInteract err", zap.Error(err))
// 		MakeApiResponseErrorDefault(c)
// 		return
// 	}

// 	result := service.UpdateUserEssayInteractLike(userEssayInteract, uid, circleId, essayId)
// 	if result.Error != nil {
// 		service.Logger.Error("UpdateUserEssayInteractLike err", zap.Error(result.Error))
// 		MakeApiResponseErrorDefault(c)
// 		return
// 	}

// 	result = service.UpdateUserEssayInteractCollect(userEssayInteract, uid, circleId, essayId, favorite)
// 	if result.Error != nil {
// 		service.Logger.Error("UpdateUserEssayInteractCollect err", zap.Error(result.Error))
// 		MakeApiResponseErrorDefault(c)
// 		return
// 	}

// }
