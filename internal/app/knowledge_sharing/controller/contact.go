package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 添加联系人
func AddUserContactHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	// model.Contact

	receiveIdStr := c.Query("receive_id")
	if receiveIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	receiveId, err := strconv.Atoi(receiveIdStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	//查询用户的联系人
	contact, err := service.GetUserContact(uid, receiveId)
	if err != nil {
		service.Logger.Error("GetUserContact", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//仅 不存在，存在状态为删除 两种
	if contact == nil {
		createTime := time.Now()

		newContact := &model.Contact{ //其中包含自动生成的id
			SendId:        uid,
			ReceiveId:     receiveId,
			CreateAt:      &createTime,
			UpdateAt:      &createTime,
			ContactStatus: model.CONTACT_STATUS_NORMAL,
		}

		err = service.CreateUserContact(newContact)
		if err != nil {
			service.Logger.Error("CreateUserContact err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		MakeApiResponseSuccessDefault(c)
		return
	}
}
