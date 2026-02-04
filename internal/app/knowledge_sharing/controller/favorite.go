package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 收藏夹
func AddFavoriteHandler(c *gin.Context) { //c
	// 从表单中获取用户信息
	title := c.PostForm("title")

	titleLen := len(title)
	if titleLen > model.FAVORITE_MAX_TITLE || titleLen == 0 {
		MakeApiResponseError(c, CODE_INTERACT_FAVORITE_LEN_INVASLID)
		return
	}

	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	createTime := time.Now()

	// 收藏夹
	newFavorite := &model.Favorite{ //其中包含自动生成的id
		Title:    title,
		UserId:   uid,
		CreateAt: &createTime,
		UpdateAt: &createTime,
	}

	// 插入数据库
	err := service.CreateFavorite(newFavorite)
	if err != nil {
		service.Logger.Error("CreateFavorite err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	// 返回成功响应
	MakeApiResponseSuccessDefault(c)
}

// 获取用户全部的收藏夹
func GetUserAllFavoriteHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	page := c.GetInt("page")
	if page == 0 {
		page = 1
	}

	pagesize := 10

	//获取全部favorite
	favorites, err := service.GetAllFavoriteByUid(uid, page, pagesize)
	if err != nil {
		service.Logger.Error("GetAllFavoriteByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"favorites": favorites,
	})
}

// 根据fid获取收藏夹
func GetFavoriteHandler(c *gin.Context) {
	//获取收藏夹id
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
	}

	//根据fid获取收藏夹
	favorite, err := service.GetFavoriteByFid(fid)
	if err != nil {
		service.Logger.Error("GetFavoriteByFid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if favorite == nil {
		MakeApiResponseError(c, CODE_ESSAY_NOT_EXIST)
		return
	}

	data := map[string]interface{}{
		"favorite": favorite,
	}

	MakeApiResponseSuccess(c, data)
}
