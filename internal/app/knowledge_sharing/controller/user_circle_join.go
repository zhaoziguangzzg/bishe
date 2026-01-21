package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 用户参与圈子
func AddUserCircleJoinHandle(c *gin.Context) {
	uidStr := c.PostForm("user_id")
	circleIdStr := c.PostForm("circle_id")

	if uidStr == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	if circleIdStr == "" {
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

	joinTime := time.Now()

	newUserCircle := &model.UserCircleJoin{
		UserId:   uid,
		CircleId: circleId,
		JoinTime: &joinTime,
		Status:   model.UserCircleJoinNormal,
	}

	err = service.CreateUserCircleJoin(newUserCircle)
	if err != nil {
		service.Logger.Error("CreateUserCircleJoin err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccess(c, newUserCircle)
}
