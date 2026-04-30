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

	uid := c.GetInt("uid")

	// 用户创建收藏夹之前，判断Favorite
	favorite, err := service.GetFavoriteByTitle(title, uid)
	if err != nil {
		service.Logger.Error("GetFavoriteByTitle err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if favorite != nil {
		MakeApiResponseError(c, CODE_FAVORITE_EXIST)
		return
	}

	createTime := time.Now()

	// 收藏夹
	newFavorite := &model.Favorite{ //其中包含自动生成的id
		Title:          title,
		UserId:         uid,
		CreateAt:       &createTime,
		UpdateAt:       &createTime,
		FavoriteStatus: model.FAVORITE_STATUS_NORMAL,
		IsDeleted:      model.FAVORITE_NOT_DELETED,
	}

	err = service.CreateFavorite(newFavorite)
	if err != nil {
		service.Logger.Error("CreateFavorite err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 更新收藏夹名
func UpdateFavoriteTitleHandler(c *gin.Context) {
	title := c.PostForm("title")

	titleLen := len(title)
	if titleLen == 0 || titleLen > 100 {
		MakeApiResponseError(c, CODE_ESSAY_TITLE_LEN_INVASLID)
		return
	}

	fidStr := c.PostForm("fid")
	if fidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	fid, err := strconv.Atoi(fidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	//根据fid获取收藏夹
	favorite, err := service.GetFavoriteByFid(fid)
	if err != nil {
		service.Logger.Error("GetFavoriteByFid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if favorite == nil {
		MakeApiResponseError(c, CODE_FAVORITE_NOT_EXIST)
		return
	}

	affectRows, err := service.UpdateFavoriteTitleByFid(fid, title)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateFavoriteTitleByFid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 删除收藏夹
func DeletedFavoriteByUpdateIsDeletedHandler(c *gin.Context) {
	//更新字段
	fidStr := c.PostForm("fid")
	if fidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	fid, err := strconv.Atoi(fidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
	}

	affectRows, err := service.UpdateFavoriteIsDeleted(fid)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateFavoriteIsDeleted err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 获取用户全部的收藏夹
func GetUserAllFavoriteHandler(c *gin.Context) {
	uid := c.GetInt("uid")

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	//获取全部favorite
	favorites, err := service.GetAllFavoriteByUid(uid, page, pagesize)
	if err != nil {
		service.Logger.Error("GetAllFavoriteByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if favorites == nil {
		favorites = make([]model.Favorite, 0)
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"favorites": favorites,
	})
}

// 根据uid获取用户全部的收藏夹
func GetUserAllFavoriteByUidHandler(c *gin.Context) {
	uidStr := c.Query("uid")
	if uidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	//获取全部favorite
	favorites, err := service.GetAllFavoriteByUid(uid, page, pagesize)
	if err != nil {
		service.Logger.Error("GetAllFavoriteByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if favorites == nil {
		favorites = make([]model.Favorite, 0)
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
		MakeApiResponseErrorParams(c)
		return
	}

	fid, err := strconv.Atoi(fidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
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
