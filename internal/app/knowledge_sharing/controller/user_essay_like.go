package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
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

	eid := c.GetInt("eid")
	if eid == 0 {
		service.Logger.Error("geteid err", zap.String("err", "get eid"))
		MakeApiResponseErrorDefault(c)
		return
	}

	createTime := time.Now()

	newUserEssayLike := &model.UserEssayLike{ //其中包含自动生成的id
		UserId:   uid,
		EssayId:  eid,
		CreateAt: &createTime,
		UpdateAt: &createTime,
	}

	err := service.CreateUserEssayLike(newUserEssayLike)
	if err != nil {
		service.Logger.Error("CreateUserEssayLike err", zap.Error(err))
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

	eid := c.GetInt("eid")
	if eid == 0 {
		service.Logger.Error("geteid err", zap.String("err", "get eid"))
		MakeApiResponseErrorDefault(c)
		return
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
	if page == 0 {
		service.Logger.Error("GetInt page err", zap.String("err", "get page err"))
		MakeApiResponseErrorDefault(c)
		return
	}

	pageSize := 10

	//获取用户全部like
	likes, err := service.GetUserAllLikeByUid(uid, page, pageSize)
	if err != nil {
		service.Logger.Error("GetUserAllLikeByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"likes": likes,
	})
}

// 更新用户喜欢
func UpdateUserEssayLike(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	delete := c.GetInt("delete")
	if delete == 0 {
		service.Logger.Error("GetInt delete err", zap.String("err", "get delete"))
		MakeApiResponseErrorDefault(c)
		return
	}

	eid := c.GetInt("eid")
	if eid == 0 {
		service.Logger.Error("geteid err", zap.String("err", "get eid"))
		MakeApiResponseErrorDefault(c)
		return
	}

	//喜欢转不喜欢
	if delete == model.LIKE_NOT_DELETED {
		affectRows, err := service.UpdateUserEssayLikeIsToNot(uid, eid)
		if err != nil || affectRows == 0 {
			service.Logger.Error("UpdateUserEssayLikeIsToNot err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		MakeApiResponseSuccessDefault(c)
		return
	}

	//不喜欢转喜欢
	affectRows, err := service.UpdateUserEssayLikeNotToIs(uid, eid)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateUserEssayLikeNotToIs err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)

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
