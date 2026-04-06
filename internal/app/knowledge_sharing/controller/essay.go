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

	if len(essays) == 0 {
		essays = make([]model.Essay, 0)
		data := map[string]interface{}{
			"essays": essays,
		}

		MakeApiResponseSuccess(c, data)
		return
	}

	var uids []int
	for _, v := range essays {
		uids = append(uids, v.AuthorId)
	}

	//根据uids获取userMap
	userMap, err := service.GetUsersByUidMap(uids)
	if err != nil {
		service.Logger.Error("GetUsersByUids", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(userMap) == 0 {
		service.Logger.Error("GetUsersByUidMap len(userMap) == 0")
		MakeApiResponseErrorDefault(c)
		return
	}

	//根据uids获取levelScoreMap
	levelScoreMap, err := service.GetLevelScoreMapByUids(uids, cid)
	if err != nil {
		service.Logger.Error("GetLevelScoreMapByUids", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(levelScoreMap) == 0 {
		service.Logger.Error("GetLevelScoreMapByUids len(levelScoreMap) == 0")
		MakeApiResponseErrorDefault(c)
		return
	}

	userEssays := make([]model.UserEssay, 0)

	for _, v := range essays {
		vUid := v.AuthorId

		vUser, ok := userMap[vUid]
		if !ok {
			service.Logger.Error("get user err")
			MakeApiResponseErrorDefault(c)
			return
		}

		vLevelScore, ok := levelScoreMap[vUid]
		if !ok {
			service.Logger.Error("get levelScore err")
			MakeApiResponseErrorDefault(c)
			return
		}

		level := vLevelScore.Score / 1000

		var userEssay model.UserEssay

		userEssay.Uid = vUid
		userEssay.Name = vUser.Name
		userEssay.Level = level
		userEssay.Essay = v
		userEssays = append(userEssays, userEssay)
	}

	data := map[string]interface{}{
		"userEssays": userEssays,
	}

	MakeApiResponseSuccess(c, data)
}

// 将文章添加周刊
func AddEssayWeeklyHandler(c *gin.Context) {
	eidStr := c.PostForm("eid")
	if eidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	// 更新IsWeekly 添加文章周刊
	affectRows, err := service.AddEssayWeekly(eid)
	if err != nil || affectRows == 0 {
		service.Logger.Error("AddEssayWeekly err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 获取文章周刊
func GetEssayWeeklylistHandler(c *gin.Context) {
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

	essays, err := service.GetEssayWeeklyList(cid, page, pagesize)
	if err != nil {
		service.Logger.Error("GetEssayWeeklyList err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(essays) == 0 {
		essays = make([]model.Essay, 0)
	}

	data := map[string][]model.Essay{
		"essays": essays,
	}

	MakeApiResponseSuccess(c, data)

}

// 将文章添加精粹
func AddEssayEssenceHandler(c *gin.Context) {
	eidStr := c.PostForm("eid")
	if eidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	// 将文章添加精粹
	affectRows, err := service.AddEssayEssence(eid)
	if err != nil || affectRows == 0 {
		service.Logger.Error("AddEssayEssence err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)
}

// 获取文章精粹
func GetEssayEssonceHandler(c *gin.Context) {
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

	essays, err := service.GetEssayEssenceList(cid, page, pagesize)
	if err != nil {
		service.Logger.Error("GetEssayEssenceList err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(essays) == 0 {
		essays = make([]model.Essay, 0)
	}

	data := map[string][]model.Essay{
		"essays": essays,
	}

	MakeApiResponseSuccess(c, data)
}

// 获取文章
func GetEssayByTitleHandler(c *gin.Context) {
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

	title := c.Query("title")
	if title == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	titleLen := len(title)
	if titleLen > model.CIRCLE_MAX_TITLE || titleLen == 0 {
		MakeApiResponseError(c, CODE_CIRCLE_TITLE_LEN_INVASLID)
		return
	}

	//根据title获取圈子文章
	essay, err := service.GetEssayByTitle(title, cid)
	if err != nil {
		service.Logger.Error("GetEssayByTitle", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if essay == nil {
		MakeApiResponseError(c, CODE_ESSAY_NOT_EXIST)
		return
	}

	//根据id获取用户
	user, err := service.GetUserByUserId(essay.AuthorId)
	if err != nil {
		service.Logger.Error("GetUserByUserId", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if user == nil {
		MakeApiResponseError(c, CODE_SYS_ERROR)
		return
	}

	data := map[string]interface{}{
		"essay": essay,
		"user":  user,
	}

	MakeApiResponseSuccess(c, data)

}
