package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 用户在文章的收藏
func AddUserEssayCollectHandle(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	eid := c.GetInt("eid")
	if eid == 0 {
		service.Logger.Error("geteid err", zap.String("err", "get eid"))
		MakeApiResponseErrorDefault(c)
		return
	}

	favorite := c.PostForm("favorite")
	favoriteLen := len(favorite)
	if favoriteLen > model.FAVOTRITE_MAX_CONTENT || favoriteLen == 0 {
		MakeApiResponseError(c, CODE_INTERACT_FAVORITE_LEN_INVASLID)
		return
	}

	createTime := time.Now()

	newUserEssayCollect := &model.UserEssayCollect{ //其中包含自动生成的id
		UserId:   uid,
		EssayId:  eid,
		Favorite: favorite,
		CreateAt: &createTime,
		UpdateAt: &createTime,
	}

	err := service.CreateUserEssayCollect(newUserEssayCollect)
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

	eid := c.GetInt("eid")
	if eid == 0 {
		service.Logger.Error("geteid err", zap.String("err", "get eid"))
		MakeApiResponseErrorDefault(c)
		return
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
		"like": collect,
	}

	MakeApiResponseSuccess(c, data)

}
