package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 用户参与圈子
func AddUserCircleJoinHandle(c *gin.Context) {

	cid := c.GetInt("cid")
	if cid == 0 {
		service.Logger.Error("GetInt cid err", zap.String("err", "get cid err"))
		MakeApiResponseErrorDefault(c)
		return
	}

	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	//todo 用户加入圈子之前，判断join_status 是否=1

	joinTime := time.Now()
	newUserCircle := &model.UserCircleJoin{
		UserId:   uid,
		CircleId: cid,
		JoinTime: &joinTime,
		UpdateAt: &joinTime,
	}

	err := service.CreateUserCircleJoin(newUserCircle)
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

	// sendId := 1111

	// createTime := time.Now()

	// notice := &model.Information{
	// 	SendId: sendId,
	// 	// ReceiveAccount: UserAccount,
	// 	Content:  content, //用户加入圈子成功
	// 	CreateAt: &createTime,
	// }

	// err = service.AddUserNotice(notice)
	// if err != nil {
	// 	service.Logger.Error("AddUserNotice", zap.Error(err))
	// 	MakeApiResponseError(c, CODE_SYS_ERROR)
	// }

	MakeApiResponseSuccessDefault(c)
}

// 用户退出圈子
func UserQuitCircleHandler(c *gin.Context) {
	cid := c.GetInt("cid")
	if cid == 0 {
		service.Logger.Error("GetInt cid err", zap.String("err", "get cid err"))
		MakeApiResponseErrorDefault(c)
		return
	}

	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	//退出用户加入圈子
	//更新join_status
	affectRows, err := service.UpdateUserCircleJoinStatusByUidCid(uid, cid)
	if affectRows == 0 || err != nil {
		service.Logger.Error("UpdateUserCircleJoinStatusByUidCid err", zap.Error(err))
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
