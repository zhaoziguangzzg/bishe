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

	cidStr := c.PostForm("cid")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	uid := service.GetUidFromContext(c)

	circle, err := service.GetCircleByCid(cid)
	if err != nil {
		service.Logger.Error("GetCircleByCid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if circle == nil {
		MakeApiResponseError(c, CODE_CIRCLE_NOT_EXIST)
		return
	}

	// 用户加入圈子之前，判断join_status 是否=1
	join, err := service.GetUserCircleJoinByUidCid(uid, cid)
	if err != nil {
		service.Logger.Error("GetUserCircleJoinByUidCid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if join != nil {
		//join非空,加入过
		if join.NotJoinStatus == model.USER_CIRCLE_JOIN_STATUS_JOIN {
			// 提示已加入
			MakeApiResponseError(c, CODE_USER_JOIN_CIRCLE)
			return
		}

		//用户加入圈子
		//更新join_status
		affectRows, err := service.UpdateUserCircleJoinStatusByJid(join.Id, model.USER_CIRCLE_JOIN_STATUS_JOIN)
		if affectRows == 0 || err != nil {
			service.Logger.Error("UpdateUserCircleJoinStatusByJid err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		//更新circle join_num+1
		affectRows, err = service.IncrUpdateCircleJoinNumByCid(cid)
		if affectRows == 0 || err != nil {
			service.Logger.Error("IncrUpdateCircleJoinNumByCid err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		MakeApiResponseSuccessDefault(c)
		return
	} else {
		//空，未加入过
		nowTime := time.Now()
		endTime := nowTime.AddDate(1, 0, 0)
		joinId, err := service.CreateUserJoinCircleAndUpdateJoinNum(uid, cid, nowTime, endTime)
		if err != nil {
			service.Logger.Error("CreateUserJoinCircleAndUpdateJoinNum err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		if joinId == 0 {
			MakeApiResponseErrorDefault(c)
			return
		}

		//用户加入圈子，发送通知
		typei := model.NOTICE_TYPE_JOIN
		noticeMsg := &model.NoticeMsg{
			Type:     typei,
			Time:     nowTime.Unix(),
			JoinUid:  uid,
			CircleId: cid,
		}

		_, _, err = service.ProduceKafkaNoticeMessage(noticeMsg)
		if err != nil {
			service.Logger.Error("ProduceKafkaNoticeMessage err", zap.Error(err))
			err = nil
		}

		err = service.UserAddLevelScore(uid, cid, nowTime)
		if err != nil {
			service.Logger.Error("UserAddLevelScore err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		MakeApiResponseSuccessDefault(c)
		return
	}

}

// 用户退出圈子
func QuitCircleHandler(c *gin.Context) {
	cidStr := c.PostForm("cid")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	uid := service.GetUidFromContext(c)

	// 获取用户是否加入圈子
	// 用户加入圈子之前，判断not_join_status 是否=0
	join, err := service.GetUserCircleJoinByUidCid(uid, cid)
	if err != nil {
		service.Logger.Error("GetUserCircleJoinByUidCid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if join == nil {
		MakeApiResponseError(c, CODE_USER_NOT_JOIN_CIRCLE)
		return
	}

	//退出用户加入圈子
	//更新not_join_status
	affectRows, err := service.UpdateUserCircleJoinStatusByJid(join.Id, model.USER_CIRCLE_JOIN_STATUS_NOT_JOIN)
	if affectRows == 0 || err != nil {
		service.Logger.Error("UpdateUserCircleJoinStatusByJid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//更新circle join_num-1
	affectRows, err = service.DecrrUpdateCircleJoinNumByCid(cid)
	if affectRows == 0 || err != nil {
		service.Logger.Error("DecrrUpdateCircleJoinNumByCid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}
