package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"bishe/internal/app/knowledge_sharing/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 用户在文章的评论
func AddUserEssayCommentHandle(c *gin.Context) {
	uid := service.GetUidFromContext(c)

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

	content := c.PostForm("content")
	contentLen := len(content)
	if contentLen > model.COMMENT_MAX_CONTENT || contentLen == 0 {
		MakeApiResponseError(c, CODE_COMMENT_CONTENT_LEN_INVASLID)
		return
	}

	nowTime := time.Now()

	newUserEssayComment := &model.UserEssayComment{ //其中包含自动生成的id
		UserId:        uid,
		EssayId:       eid,
		Content:       content,
		CreateAt:      &nowTime,
		UpdateAt:      &nowTime,
		CommentStatus: model.COMMENT_STATUS_NORMAL,
		IsDeleted:     model.COMMENT_NOT_DELETED,
	}

	err = service.CreateUserEssayComment(newUserEssayComment)
	if err != nil {
		service.Logger.Error("CreateUserEssayComment err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	typei := model.STAT_TYPE_COMMENT

	//添加或更新用户统计数
	err = service.StatInsertUpdate(uid, 1, typei, nowTime)
	if err != nil {
		service.Logger.Error("StatInsertUpdate err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//添加文章评论数据详情
	err = service.StatDetailsInsert(uid, typei, model.STAT_DETAILS_STATUS_INCR, nowTime)
	if err != nil {
		service.Logger.Error("StatDetailsInsert err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//更新文章评论数
	affectRows, err := service.UpdateEssayCommentNum(eid, 1)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateEssayCommentNum err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//通知
	typei = model.NOTICE_TYPE_COMMENT

	noticeMsg := &model.NoticeMsg{
		Type:       typei,
		Time:       nowTime.Unix(),
		CommentUid: uid,
		EssayId:    essay.Id,
	}

	_, _, err = service.ProduceKafkaNoticeMessage(noticeMsg)
	if err != nil {
		service.Logger.Error("ProduceKafkaNoticeMessage err", zap.Error(err))
		err = nil
	}

	MakeApiResponseSuccessDefault(c)
}

// 删除评论
func DeletedCommentByUpdateIsDeletedHandler(c *gin.Context) {
	commentIdStr := c.PostForm("comment_id")
	if commentIdStr == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	commentId, err := strconv.Atoi(commentIdStr)
	if err != nil {
		MakeApiResponseErrorDefault(c)
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

	essay, err := service.GetEssayByEid(eid)
	if err != nil {
		service.Logger.Error("GetEssayByEid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	nowTime := time.Now()

	//更新isDeleted
	affectRows, err := service.UpdateIsDeletedByCommentId(commentId)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateIsDeletedByCommentId err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//更新文章评论数据详情
	err = service.UpdateStatAndStatDetail(essay.AuthorId, model.STAT_TYPE_COMMENT, model.STAT_DETAILS_STATUS_DECR, nowTime)
	if err != nil {
		service.Logger.Error("UpdateStatAndStatDetail err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	//更新文章评论数
	affectRows, err = service.UpdateEssayCommentNum(eid, -1)
	if err != nil || affectRows == 0 {
		service.Logger.Error("UpdateEssayCommentNum err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccessDefault(c)

}

// 获取文章评论列表
func GetEssayAllCommentHandle(c *gin.Context) {
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

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pageSize := 10

	comments, err := service.GetEssayAllComment(eid, page, pageSize)
	if err != nil {
		service.Logger.Error("GetEssayAllComment", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(comments) == 0 {
		data := map[string]interface{}{
			"userComments": make([]model.UserComment, 0),
		}
		MakeApiResponseSuccess(c, data)
		return
	}

	var uids []int
	for _, comment := range comments {
		uids = append(uids, comment.UserId)
	}

	userMap := make(map[int]model.User)
	userMap, err = service.GetUserMapByUids(uids)
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

	userComments := make([]model.UserComment, 0)

	for _, v := range comments {
		vUid := v.UserId

		vUser, ok := userMap[vUid]
		if !ok {
			service.Logger.Error("get user err")
			MakeApiResponseErrorDefault(c)
			return
		}

		var userComment model.UserComment

		userComment.User = vUser
		userComment.Comment = v
		userComments = append(userComments, userComment)
	}

	data := map[string]interface{}{
		"userComments": userComments,
	}

	MakeApiResponseSuccess(c, data)
}

// 获取用户全部评论列表
func GetUserAllCommentHandler(c *gin.Context) {
	uid := service.GetUidFromContext(c)

	pageStr := c.Query("page")
	page := GetPage(pageStr)

	pageSize := 10

	//获取用户全部评论文章
	commentEssays, err := service.GetUserAllCommentIdByUid(uid, page, pageSize)
	if err != nil {
		service.Logger.Error("GetEssayAllCommentByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(commentEssays) == 0 {
		commentEssays = make([]model.CommentEssay, 0)
	}

	data := map[string]interface{}{
		"commentEssays": commentEssays,
	}

	MakeApiResponseSuccess(c, data)
}

// 根据uid获取用户全部评论列表
func GetUserAllCommentByUidHandler(c *gin.Context) {
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

	pageSize := 10

	//获取用户全部评论文章
	commentEssays, err := service.GetUserAllCommentIdByUid(uid, page, pageSize)
	if err != nil {
		service.Logger.Error("GetEssayAllCommentByUid", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	if len(commentEssays) == 0 {
		commentEssays = make([]model.CommentEssay, 0)
	}

	data := map[string]interface{}{
		"commentEssays": commentEssays,
	}

	MakeApiResponseSuccess(c, data)
}
