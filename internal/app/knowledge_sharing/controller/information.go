package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 消息通知
func CreateInformationHandle(c *gin.Context) {
	receiveAccountStr := c.PostForm("receiveAccont")
	content := c.PostForm("content")

	if receiveAccountStr == "" {
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	receiveAccount, err := strconv.Atoi(receiveAccountStr)
	if err != nil {
		service.Logger.Error("receiveAccountStrAtoi err", zap.Error(err))
		MakeApiResponseError(c, CODE_PARAMS_ERROR)
		return
	}

	createTime := time.Now()

	//消息结构体
	information := &model.Information{
		// SendId:         UserId,
		ReceiveAccount: receiveAccount,
		Content:        content,
		CreateAt:       &createTime,
	}

	err = service.UserAddInformation(information)
	if err != nil {
		service.Logger.Error("userAddInformation err", zap.Error(err))
		MakeApiResponseError(c, CODE_SYS_ERROR)
		return
	}
}
