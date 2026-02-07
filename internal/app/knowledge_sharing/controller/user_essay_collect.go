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

	eidStr := c.PostForm("eid")
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

	fidStr := c.PostForm("fid")
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

	//is_deleted 直查询有效的
	collect, err := service.GetUserEssayCollect(uid, eid)
	if err != nil {
		service.Logger.Error("GetUserEssayCollect", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if collect != nil {
		MakeApiResponseError(c, CODE_COLLECT_EXIST)
		return
	}

	createTime := time.Now()

	newUserEssayCollect := &model.UserEssayCollect{ //其中包含自动生成的id
		UserId:     uid,
		EssayId:    eid,
		FavoriteId: fid,
		CreateAt:   &createTime,
		UpdateAt:   &createTime,
	}

	//TODO 去唯一键，可以重复

	err = service.CreateUserEssayCollect(newUserEssayCollect)
	if err != nil {
		service.Logger.Error("CreateUserEssayCollect err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 取消用户收藏
func CancelEssayCollectHandler(c *gin.Context) {
	//删除时，一定知道id
	collectIdStr := c.PostForm("collectId")
	if collectIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	collectId, err := strconv.Atoi(collectIdStr)
	if err != nil {
		service.Logger.Error("Atoi collectIdStr err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//取消收藏
	affectRows, err := service.DeleteCollectById(collectId)
	if err != nil || affectRows == 0 {
		service.Logger.Error("DeleteCollectById err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)

}

// 获取用户文章是否收藏
func GetEssayCollectHandler(c *gin.Context) {
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

	// TODO 仅查询有效的
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
	//TODO 获取收藏夹参数

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
