package controller

import (
	"bishe/model"
	"bishe/service"
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

	lockKey := "essay-add-" + strconv.Itoa(cid) + "-" + title

	lockValue, locked, err := service.Lock(c, lockKey, 5*time.Second)
	if err != nil {
		service.Logger.Error("Lock err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if !locked {
		MakeApiResponseError(c, CODE_LOCKED)
		return
	}

	defer service.Unlock(c, lockKey, lockValue)

	uid := service.GetUidFromContext(c)

	nowTime := time.Now()

	// 构造文章
	newEssay := &model.Essay{ //其中包含自动生成的id
		Title:       title,
		Content:     content,
		CircleId:    cid,
		AuthorId:    uid,
		CreateAt:    &nowTime,
		UpdateAt:    &nowTime,
		EssayStatus: model.ESSAY_STATUS_NORMAL,
		IsDeleted:   model.ESSAY_NOT_DELETED,
	}

	err = service.CreateEssay(newEssay)
	if err != nil {
		service.Logger.Error("CreateEssay err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	typei := model.NOTICE_TYPE_ESSAY_ADD
	noticeMsg := &model.NoticeMsg{
		Type:      typei,
		Time:      nowTime.Unix(),
		AuthorUid: newEssay.AuthorId,
		EssayId:   newEssay.Id,
	}

	_, _, err = service.ProduceKafkaNoticeMessage(noticeMsg)
	if err != nil {
		service.Logger.Error("ProduceKafkaNoticeMessage err", zap.Error(err))
		err = nil
	}

	data := map[string]interface{}{
		"eid": newEssay.Id,
	}

	MakeApiResponseSuccess(c, data)
}

// 更新文章信息
func UpdateEssayHandler(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	titleLen := len(title)
	if titleLen == 0 || titleLen > model.ESSAY_MAX_TITLE {
		MakeApiResponseError(c, CODE_ESSAY_TITLE_LEN_INVASLID)
		return
	}

	contentLen := len(content)
	if contentLen == 0 || contentLen > model.ESSAY_MAX_CONTENT {
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

	updateMap := map[string]interface{}{
		"title":   title,
		"content": content,
	}

	//根据eid更新文章
	affectRows, err := service.UpdateEssayByEid(eid, updateMap)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateEssayByEid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	data := map[string]interface{}{
		"essay": essay,
	}

	MakeApiResponseSuccess(c, data)
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
	uid := service.GetUidFromContext(c)

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

// 根据uid获取全部用户文章列表
func GetUserAllEssayByUidHandler(c *gin.Context) {
	uidStr := c.Query("uid")
	if uidStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
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

	author, err := service.GetUserByUserId(essay.AuthorId)
	if err != nil {
		service.Logger.Error("GetUserByUserId", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//获取当前用户是否点赞
	//获取当前用户是否收藏

	data := map[string]interface{}{
		"essay":  essay,
		"author": author,
	}

	MakeApiResponseSuccess(c, data)
}

// 获取圈子全部文章
func GetCircleAllEssayHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

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
		data := map[string]interface{}{
			"userEssays": make([]model.UserEssay, 0),
		}

		MakeApiResponseSuccess(c, data)
		return
	}

	var uids []int
	var eids []int
	for _, v := range essays {
		uids = append(uids, v.AuthorId)
		eids = append(eids, v.Id)
	}

	//根据uids获取userMap
	userMap, err := service.GetUserMapByUids(uids)
	if err != nil {
		service.Logger.Error("GetUserMapByUids", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(userMap) == 0 {
		service.Logger.Error("GetUserMapByUids len(userMap) == 0")
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

	//根据eids获取essayLikeMap
	essayLikeMap, err := service.GetUserEssayLikeMapByEids(uid, eids)
	if err != nil {
		service.Logger.Error("GetUserEssayLikeMapByEids", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//根据eids获取essayCollectMap
	essayCollectMap, err := service.GetUserEssayCollectMapByEids(uid, eids)
	if err != nil {
		service.Logger.Error("GetUserEssayCollectMapByEids", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	// if len(essayLikeMap) == 0 {

	// }

	userEssays := make([]model.UserEssay, 0)

	for _, v := range essays {
		var isLike bool
		var isCollect bool
		vUid := v.AuthorId

		vUser, ok := userMap[vUid]
		if !ok {
			service.Logger.Error("get user err")
			MakeApiResponseErrorDefault(c)
			return
		}

		vLevelScore, ok := levelScoreMap[vUid]
		if !ok {
			vLevelScore = model.LevelScore{Score: 0}
		}

		userEssayLike, ok := essayLikeMap[v.Id]
		if ok && userEssayLike.LikeStatus == model.LIKE_STATUS_NORMAL {
			isLike = true
		}

		userEssayCollect, ok := essayCollectMap[v.Id]
		if ok && userEssayCollect.CollectStatus == model.COLLECT_STATUS_NORMAL {
			isCollect = true
		}

		level := vLevelScore.Score / 1000

		var userEssay model.UserEssay

		userEssay.Author = vUser
		userEssay.Level = level
		userEssay.Essay = v
		userEssay.IsLike = isLike
		userEssay.IsCollect = isCollect
		userEssays = append(userEssays, userEssay)
	}

	data := map[string]interface{}{
		"userEssays": userEssays,
	}

	MakeApiResponseSuccess(c, data)
}

// 更新文章周刊
func UpdateEssayWeeklyHandler(c *gin.Context) {
	eidStr := c.PostForm("eid")
	if eidStr == "" {
		MakeApiResponseErrorParams(c)
		service.Logger.Error("eidStr == nil ")
		return
	}

	eid, err := strconv.Atoi(eidStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	isWeeklyStr := c.PostForm("is_weekly")
	if isWeeklyStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	isWeekly, err := strconv.Atoi(isWeeklyStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	switch isWeekly {
	case model.ESSAY_IS_WEEKLY:
		isWeekly = model.ESSAY_NOT_WEEKLY
	case model.ESSAY_NOT_WEEKLY:
		isWeekly = model.ESSAY_IS_WEEKLY
	default:
		MakeApiResponseErrorParams(c)
		return
	}

	// 更新IsWeekly 添加文章周刊
	affectRows, err := service.UpdateEssayWeekly(eid, isWeekly)
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

// update essay essence
func UpdateEssayEssenceHandler(c *gin.Context) {
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

	isEssenceStr := c.PostForm("is_essence")
	if isEssenceStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	isEssence, err := strconv.Atoi(isEssenceStr)
	if err != nil {
		MakeApiResponseErrorParams(c)
		return
	}

	switch isEssence {
	case model.ESSAY_IS_ESSENCE:
		isEssence = model.ESSAY_NOT_ESSENCE
	case model.ESSAY_NOT_ESSENCE:
		isEssence = model.ESSAY_IS_ESSENCE
	default:
		MakeApiResponseErrorParams(c)
		return
	}

	//根据id获取文章
	essay, err := service.GetEssayByEid(eid)
	if err != nil {
		service.Logger.Error("GetEssayByEid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if essay == nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	switch isEssence {
	case model.ESSAY_IS_ESSENCE:
	case model.ESSAY_NOT_ESSENCE:
	default:
		MakeApiResponseErrorParams(c)
		return
	}

	levelScore, err := service.GetUserLevelScoreByUidCid(essay.AuthorId, essay.CircleId)
	if err != nil {
		service.Logger.Error("GetUserLevelScoreByUidCid err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if levelScore == nil {
		MakeApiResponseErrorDefault(c)
		return
	}

	nowTime := time.Now()

	// Update essay essence
	affectRows, err := service.UpdateEssayEssence(eid, isEssence)
	if err != nil || affectRows == 0 {
		service.Logger.Error("AddEssayEssence err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	typei := model.LEVEL_SCORE_RECORD_TYPE_ESSENCE
	levelScoreRecord, err := service.GetUserLevelScoreRecordByUidCidType(essay.AuthorId, essay.CircleId, typei)
	if err != nil {
		service.Logger.Error("GetUserLevelScoreRecordByUidCidType err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	switch isEssence {
	case model.ESSAY_IS_ESSENCE:
		//加分
		score := 100

		//加精，更新等级分数 增加100
		affectRows, err = service.UpdateLevelScoreByUidCid(essay.AuthorId, essay.CircleId, score)
		if err != nil || affectRows == 0 {
			service.Logger.Error("UpdateLevelScoreByUidCid err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		//根据条件判断是否需要增加等级详情
		if levelScoreRecord == nil {
			//增加等级详情
			err = service.UserAddLevelScoreRecord(essay.AuthorId, essay.CircleId, score, eid, typei, time.Now())
			if err != nil {
				service.Logger.Error("UserAddLevelScoreRecord err", zap.Error(err))
				MakeApiResponseErrorDefault(c)
				return
			}
		}

		if levelScoreRecord.IsDeleted == model.IS_DELETED_YES {
			//更新等级详情
			err = service.UserUpdateLevelScoreRecord(essay.AuthorId, essay.CircleId, typei, model.IS_DELETED_NO)
			if err != nil {
				service.Logger.Error("UserUpdateLevelScoreRecord err", zap.Error(err))
				MakeApiResponseErrorDefault(c)
				return
			}
		}

	case model.ESSAY_NOT_ESSENCE:
		//减分
		score := -100
		//减精，更新等级分数 减少100
		affectRows, err = service.UpdateLevelScoreByUidCid(essay.AuthorId, essay.CircleId, score)
		if err != nil || affectRows == 0 {
			service.Logger.Error("UpdateLevelScoreByUidCid err", zap.Error(err))
			MakeApiResponseErrorDefault(c)
			return
		}

		//根据条件判断是否需要增加等级详情
		if levelScoreRecord.IsDeleted == model.IS_DELETED_NO {
			//删除等级详情
			err = service.UserUpdateLevelScoreRecord(essay.AuthorId, essay.CircleId, typei, model.IS_DELETED_YES)
			if err != nil {
				service.Logger.Error("UserUpdateLevelScoreRecord err", zap.Error(err))
				MakeApiResponseErrorDefault(c)
				return
			}
		}

	default:
		MakeApiResponseErrorParams(c)
		return
	}

	// 给用户发通知
	typei = model.NOTICE_TYPE_ESSENCE

	noticeMsg := &model.NoticeMsg{
		Type:    typei,
		Time:    nowTime.Unix(),
		EssayId: eid,
	}

	_, _, err = service.ProduceKafkaNoticeMessage(noticeMsg)
	if err != nil {
		service.Logger.Error("ProduceKafkaNoticeMessage err", zap.Error(err))
		err = nil
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
	uid := service.GetUidFromContext(c)
	if uid == 0 {
		MakeApiResponseErrorDefault(c)
		return
	}

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
	if titleLen > model.ESSAY_MAX_TITLE || titleLen == 0 {
		MakeApiResponseError(c, CODE_ESSAY_TITLE_LEN_INVASLID)
		return
	}

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pagesize := 10

	//根据title关键词like获取文章
	essays, err := service.GetEssayByLikeTitle(title, cid, page, pagesize)
	if err != nil {
		service.Logger.Error("GetEssayByLikeTitle", zap.Error(err))
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
	var eids []int
	for _, v := range essays {
		uids = append(uids, v.AuthorId)
		eids = append(eids, v.Id)
	}

	//根据uids获取userMap
	userMap, err := service.GetUserMapByUids(uids)
	if err != nil {
		service.Logger.Error("GetUserMapByUids", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(userMap) == 0 {
		service.Logger.Error("GetUserMapByUids len(userMap) == 0")
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

	//根据eids获取essayLikeMap
	essayLikeMap, err := service.GetUserEssayLikeMapByEids(uid, eids)
	if err != nil {
		service.Logger.Error("GetUserEssayLikeMapByEids", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//根据eids获取essayCollectMap
	essayCollectMap, err := service.GetUserEssayCollectMapByEids(uid, eids)
	if err != nil {
		service.Logger.Error("GetUserEssayCollectMapByEids", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	userEssays := make([]model.UserEssay, 0)

	for _, v := range essays {
		isLike := false
		isCollect := false
		userEssayLike, ok := essayLikeMap[v.Id]
		if ok && userEssayLike.LikeStatus == model.LIKE_STATUS_NORMAL {
			isLike = true
		}
		userEssayCollect, ok := essayCollectMap[v.Id]
		if ok && userEssayCollect.CollectStatus == model.COLLECT_STATUS_NORMAL {
			isCollect = true
		}
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

		userEssay.Author = vUser
		userEssay.Level = level
		userEssay.Essay = v
		userEssay.IsLike = isLike
		userEssay.IsCollect = isCollect
		userEssays = append(userEssays, userEssay)
	}

	data := map[string]interface{}{
		"userEssays": userEssays,
	}

	MakeApiResponseSuccess(c, data)

}
