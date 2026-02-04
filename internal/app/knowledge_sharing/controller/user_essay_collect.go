package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 用户在文章的收藏
func AddUserEssayCollectHandler(c *gin.Context) {
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

	// favorite := c.PostForm("favorite")
	// favoriteLen := len(favorite)
	// if favoriteLen > model.FAVOTRITE_MAX_CONTENT || favoriteLen == 0 {
	// 	MakeApiResponseError(c, CODE_INTERACT_FAVORITE_LEN_INVASLID)
	// 	return
	// }

	fidStr := c.Query("fid")
	if fidStr == "" {
		service.Logger.Error("Getfid err", zap.String("err", "get fid err"))
		MakeApiResponseErrorParams(c)
		return
	}

	fid, err := strconv.Atoi(fidStr)
	if err != nil {
		service.Logger.Error("Atoi fidStr err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//todo 收藏到收藏夹 表

	createTime := time.Now()

	newUserEssayCollect := &model.UserEssayCollect{ //其中包含自动生成的id
		UserId:     uid,
		EssayId:    eid,
		FavoriteId: fid,
		CreateAt:   &createTime,
		UpdateAt:   &createTime,
	}

	err = service.CreateUserEssayCollect(newUserEssayCollect)
	if err != nil {
		service.Logger.Error("CreateUserEssayCollect err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 获取用户文章是否收藏
func GetUserEssayCollectHandler(c *gin.Context) {
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

	collect, err := service.GetUserEssayCollect(uid, eid)
	if err != nil {
		service.Logger.Error("GetUserEssayCollect", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if collect == nil {
		MakeApiResponseError(c, CODE_COLLECT_NOT_EXIST)
		return
	}

	data := map[string]interface{}{
		"collect": collect,
	}

	MakeApiResponseSuccess(c, data)

}

// 获取用户全部收藏
func GetUserAllCollectHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	page := c.GetInt("page")
	if page == 0 {
		page = 1
	}

	pageSize := 10

	//获取用户全部collect
	collects, err := service.GetUserAllCollectByUid(uid, page, pageSize)
	if err != nil {
		service.Logger.Error("GetUserAllCollectByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"collects": collects,
	})
}

// 更新用户收藏
func UpdateUserEssayCollectHandler(c *gin.Context) {
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

	deleteStr := c.Query("delete")
	if deleteStr == "" {
		service.Logger.Error("Getdelete err", zap.String("err", "get delete err"))
		MakeApiResponseErrorParams(c)
		return
	}

	delete, err := strconv.Atoi(deleteStr)
	if err != nil {
		service.Logger.Error("Atoi deleteStr err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
	}

	fidStr := c.Query("fid")
	if fidStr == "" {
		service.Logger.Error("Getfid err", zap.String("err", "get fid err"))
		MakeApiResponseErrorParams(c)
		return
	}

	fid, err := strconv.Atoi(fidStr)
	if err != nil {
		service.Logger.Error("Atoi fidStr err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//取消收藏
	if delete == model.COLLECT_NOT_DELETED {
		affectRows, err := service.UpdateUserEssayCollectIsToNot(uid, eid)
		if err != nil || affectRows == 0 {
			service.Logger.Error("UpdateUserEssayCollectIsToNot err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		MakeApiResponseSuccessDefault(c)
		return
	}

	//添加收藏
	affectRows, err := service.UpdateUserEssayCollectNotToIs(uid, eid, fid)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateUserEssayCollectNotToIs err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)

}
