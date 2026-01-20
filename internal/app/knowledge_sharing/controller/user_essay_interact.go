package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddUserEssayLikeHandle(c *gin.Context) {
	uidStr := c.PostForm("user_id")
	circleIdStr := c.PostForm("circle_id")
	essayIdStr := c.PostForm("essay_id")
	favorite := c.PostForm("favorite")

	if uidStr == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	if circleIdStr == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	if essayIdStr == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	circleId, err := strconv.Atoi(circleIdStr)
	if err != nil {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	essayId, err := strconv.Atoi(essayIdStr)
	if err != nil {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	newUserEssayInterAct := &model.UserEssayInteract{ //其中包含自动生成的id
		UserId:        uid,
		CircleId:      circleId,
		EssayId:       essayId,
		LikeStatus:    model.UserEssayInteractNotJoin,
		CollectStatus: model.UserEssayInteractNotJoin,
		Favorite:      favorite,
	}

	err = service.CreateUserEssayInteract(newUserEssayInterAct)
	if err != nil {
		service.Logger.Error("CreateUserEssayInteract err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccess(c, newUserEssayInterAct)
}

// func GetUserEssayInteract(c *gin.Context) {

// }

func UpdateUserEssayInteract(c *gin.Context) {

	uidStr := c.PostForm("user_id")
	circleIdStr := c.PostForm("circle_id")
	essayIdStr := c.PostForm("essay_id")
	favorite := c.PostForm("favorite")

	if uidStr == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	if circleIdStr == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	if essayIdStr == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	circleId, err := strconv.Atoi(circleIdStr)
	if err != nil {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	essayId, err := strconv.Atoi(essayIdStr)
	if err != nil {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	userEssayInteract, err := service.GetUserEssayInteract(uid, circleId, essayId)
	if err != nil {
		service.Logger.Error("GetUserEssayInteract err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	result := service.UpdateUserEssayInteractLike(userEssayInteract, uid, circleId, essayId)
	if result.Error != nil {
		service.Logger.Error("UpdateUserEssayInteractLike err", zap.Error(result.Error))
		MakeApiResponseErrorDefault(c)
		return
	}

	result = service.UpdateUserEssayInteractCollect(userEssayInteract, uid, circleId, essayId, favorite)
	if result.Error != nil {
		service.Logger.Error("UpdateUserEssayInteractCollect err", zap.Error(result.Error))
		MakeApiResponseErrorDefault(c)
		return
	}

}
