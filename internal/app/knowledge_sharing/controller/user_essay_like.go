package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 用户在文章的喜欢
func AddUserEssayLikeHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	eidStr := c.Query("eid")
	if eidStr == "" {
		service.Logger.Error("Geteid err", zap.String("err", "get eid err"))
		MakeApiResponseErrorParams(c)
		return
	}

	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		service.Logger.Error("Atoi eidStr err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
	}

	//查询对文章的点赞
	like, err := service.GetUserEssayLike(uid, eid)
	if err != nil {
		service.Logger.Error("GetUserEssayLike", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if like == nil {
		createTime := time.Now()

		newUserEssayLike := &model.UserEssayLike{ //其中包含自动生成的id
			UserId:   uid,
			EssayId:  eid,
			CreateAt: &createTime,
			UpdateAt: &createTime,
		}

		err = service.CreateUserEssayLike(newUserEssayLike)
		if err != nil {
			service.Logger.Error("CreateUserEssayLike err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		MakeApiResponseSuccessDefault(c)
		return
	} else {
		if like.IsDeleted == model.LIKE_NOT_DELETED {
			MakeApiResponseError(c, CODE_LIKE_EXIST)
			return
		}

		// 返回响应
		MakeApiResponseError(c, CODE_LIKE_DELETED)
	}
}

// 取消用户喜欢
func CancelUserEssayLikeHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	eidStr := c.Query("eid")
	if eidStr == "" {
		service.Logger.Error("Geteid err", zap.String("err", "get eid err"))
		MakeApiResponseErrorParams(c)
		return
	}

	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		service.Logger.Error("Atoi eidStr err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
	}

	//TODO 获取喜欢

	//喜欢转不喜欢
	affectRows, err := service.UpdateUserEssayLikeIsToNot(uid, eid)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateUserEssayLikeIsToNot err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 获取用户文章是否喜欢
func GetUserEssayLikeHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	eidStr := c.Query("eid")
	if eidStr == "" {
		service.Logger.Error("Geteid err", zap.String("err", "get eid err"))
		MakeApiResponseErrorParams(c)
		return
	}

	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		service.Logger.Error("Atoi eidStr err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
	}

	like, err := service.GetUserEssayLike(uid, eid)
	if err != nil {
		service.Logger.Error("GetUserEssayLike", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if like == nil {
		MakeApiResponseError(c, CODE_LIKE_NOT_EXIST)
		return
	}

	data := map[string]interface{}{
		"like": like,
	}

	MakeApiResponseSuccess(c, data)
}

// 获取用户全部喜欢
func GetUserAllLikeHandler(c *gin.Context) {
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

	//获取用户全部like
	essays, err := service.GetUserAllLikeEssayByUid(uid, page, pageSize)
	if err != nil {
		service.Logger.Error("GetUserAllLikeEssayByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if essays == nil {
		essays = make([]model.Essay, 0)
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"essays": essays,
	})
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
