package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 文章
func AddEssayHandler(c *gin.Context) {
	// 从表单中获取用户信息
	title := c.PostForm("title")
	content := c.PostForm("content")

	titleLen := len(title)
	if titleLen > model.ESSAY_MAX_TITLE || titleLen == 0 {
		MakeApiResponseError(c, CODE_ESSAY_TITLE_LEN_INVASLID)
		return
	}

	contentLen := len(content)
	if contentLen > model.ESSAY_MAX_CONTENT || contentLen == 0 {
		MakeApiResponseError(c, CODE_ESSAY_CONTENT_LEN_INVASLID)
		return
	}

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

	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	createTime := time.Now()

	// 构造文章
	newEssay := &model.Essay{ //其中包含自动生成的id
		Title:       title,
		Content:     content,
		CircleId:    cid,
		AuthorId:    uid,
		CreateAt:    &createTime,
		UpdateAt:    &createTime,
		EssayStatus: model.ESSAY_STATUS_NORMAL,
		IsDeleted:   model.ESSAY_NOT_DELETED,
	}

	err = service.CreateEssay(newEssay)
	if err != nil {
		service.Logger.Error("CreateEssay err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 更新文章信息
func UpdateEssayHandler(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	titleLen := len(title)
	if titleLen == 0 || titleLen > 100 {
		MakeApiResponseError(c, CODE_ESSAY_TITLE_LEN_INVASLID)
		return
	}

	contentLen := len(content)
	if contentLen == 0 || contentLen > 200 {
		MakeApiResponseError(c, CODE_ESSAY_CONTENT_LEN_INVASLID)
		return
	}

	eidStr := c.PostForm("eid")
	if eidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	//根据eid获取文章
	essay, err := service.GetEssayByEid(eid)
	if err != nil {
		service.Logger.Error("GetEssayByEid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if essay == nil {
		MakeApiResponseError(c, CODE_ESSAY_NOT_EXIST)
		return
	}

	//根据eid更新文章
	affectRows, err := service.UpdateEssayByEid(eid, title, content)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateEssayByEid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 删除发布的文章
func DeletedEssayByUpdateIsDeletedHandler(c *gin.Context) {
	//更新字段
	eidStr := c.PostForm("eid")
	if eidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	//更新删除字段来删除文章
	affectRows, err := service.UpdateEssayIsDeleted(eid)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateEssayIsDeleted err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 获取全部用户文章列表
func GetUserAllEssayHandler(c *gin.Context) {
	uid, _ := service.GetUserFromCookie(c)
	if uid == 0 {
		MakeApiResponseError(c, CODE_USER_NOT_LOGIN)
		return
	}

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	//获取全部essay
	essays, err := service.GetAllEssayByUid(uid, page, pagesize)
	if err != nil {
		service.Logger.Error("GetAllEssayByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if essays == nil {
		essays = make([]model.Essay, 0)
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"essays": essays,
	})
}

// 获取某文章
func GetEssayHandler(c *gin.Context) {
	//获取文章id
	eidStr := c.Query("eid")
	if eidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	//根据eid获取文章
	essay, err := service.GetEssayByEid(eid)
	if err != nil {
		service.Logger.Error("GetEssayByEid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if essay == nil {
		MakeApiResponseError(c, CODE_ESSAY_NOT_EXIST)
		return
	}

	data := map[string]interface{}{
		"essay": essay,
	}

	MakeApiResponseSuccess(c, data)
}

// 获取圈子全部文章
func GetCircleAllEssayHandler(c *gin.Context) {
	cidStr := c.Query("cid")
	if cidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	//获取圈子中全部essay
	essays, err := service.GetAllEssayByCid(cid, page, pagesize)
	if err != nil {
		service.Logger.Error("GetAllEssayByCid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if essays == nil {
		essays = make([]model.Essay, 0)
	}

	MakeApiResponseSuccess(c, map[string]interface{}{
		"essays": essays,
	})
}
