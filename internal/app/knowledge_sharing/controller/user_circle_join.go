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

	cidStr := c.Query("cid")
	if cidStr == "" {
		service.Logger.Error("Getcid err", zap.String("err", "get cid err"))
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		service.Logger.Error("Atoi cidStr err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
	}

	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
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
		if join.JoinStatus == model.USER_CIRCLE_JOIN_JOIN_STATUS_JOIN {
			//TODO 提示已加入
			MakeApiResponseErrorDefault(c)
			return
		}

		//用户加入圈子
		//更新join_status
		affectRows, err := service.UpdateUserCircleJoinStatusByJid(join.Id, model.USER_CIRCLE_JOIN_JOIN_STATUS_JOIN)
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
		joinTime := time.Now()
		newUserCircle := &model.UserCircleJoin{
			UserId:     uid,
			CircleId:   cid,
			JoinTime:   &joinTime,
			UpdateAt:   &joinTime,
			JoinStatus: model.USER_CIRCLE_JOIN_JOIN_STATUS_JOIN,
		}

		err = service.CreateUserCircleJoin(newUserCircle)
		if err != nil {
			service.Logger.Error("CreateUserCircleJoin err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		affectRows, err := service.IncrUpdateCircleJoinNumByCid(cid)
		if affectRows == 0 || err != nil {
			service.Logger.Error("IncrUpdateCircleJoinNumByCid err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		MakeApiResponseSuccessDefault(c)
		return
	}

}

// 用户退出圈子
func QuitCircleHandler(c *gin.Context) {
	cidStr := c.Query("cid")
	if cidStr == "" {
		service.Logger.Error("Getcid err", zap.String("err", "get cid err"))
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		service.Logger.Error("Atoi cidStr err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
	}

	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	//TODO 获取用户是否加入圈子
	// 用户加入圈子之前，判断join_status 是否=1
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
	//更新join_status
	affectRows, err := service.UpdateUserCircleJoinStatusByJid(join.Id, model.USER_CIRCLE_JOIN_JOIN_STATUS_NO_JOIN)
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
