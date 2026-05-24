package service

import (
	"bishe/dao/kafka"
	"bishe/dao/mysql"
	"bishe/model"
	"encoding/json"
	"time"

	"go.uber.org/zap"
)

// 添加通知
func UserAddNotice(noticeUid int, content string, typei int, url string, createTime time.Time) (err error) {
	return mysql.UserAddNotice(noticeUid, content, typei, url, createTime)
}

// 添加通知
func UserAddNoticeList(noticeUids []int, content string, typei int, url string, createTime time.Time) (err error) {
	return mysql.UserAddNoticeList(noticeUids, content, typei, url, createTime)
}

// 获取通知列表
func GetNoticeList(uid int, page int, pageSize int) (notices []model.Notice, err error) {
	return mysql.GetNoticeList(uid, page, pageSize)
}

// 根据类型获取通知列表
func GetNoticeListByType(uid int, typei int, page int, pageSize int) (notices []model.Notice, err error) {
	return mysql.GetNoticeListByType(uid, typei, page, pageSize)
}

// 发送notice数据到kafka
func ProduceKafkaNoticeMessage(noticeMsg *model.NoticeMsg) (partition int32, offset int64, err error) {
	return kafka.ProduceKafkaNoticeMessage(noticeMsg)
}

// 处理kafka消息
func ProcessKafkaNotice(msg []byte) (err error) {
	var noticeMsg model.NoticeMsg
	err = json.Unmarshal(msg, &noticeMsg)
	if err != nil {
		Logger.Error("Unmarshal noticeMsg err", zap.Error(err))
		return
	}

	Logger.Info("notice msg", zap.Any("noticeMsg", noticeMsg))

	switch noticeMsg.Type {
	case model.NOTICE_TYPE_FOLLOW:
		AddNoticeFollowAdd(noticeMsg)
	case model.NOTICE_TYPE_ESSAY_ADD:
		AddNoticeEssayAdd(noticeMsg)
	case model.NOTICE_TYPE_LIKE:
		AddNoticeLike(noticeMsg)
	case model.NOTICE_TYPE_COMMENT:
		AddNoticeComment(noticeMsg)
	case model.NOTICE_TYPE_JOIN:
		AddNoticeJoin(noticeMsg)
	case model.NOTICE_TYPE_ACCUSATION:
		AddNoticeAccusation(noticeMsg)
	case model.NOTICE_TYPE_ACCUSATIONED:
		AddNoticeAccusationed(noticeMsg)
	case model.NOTICE_TYPE_FEEDBACK:
		AddNoticeFeedback(noticeMsg)
	default:
	}

	return
}

// 添加关注通知
func AddNoticeFollowAdd(msg model.NoticeMsg) {
	noticeType := model.NOTICE_TYPE_FOLLOW
	noticeTime := time.Unix(msg.Time, 0)

	fanUid := msg.FanUid
	//获取fan的信息
	if fanUid == 0 {
		return
	}

	fanUser, err := GetUserByUserId(fanUid)
	if err != nil {
		Logger.Error("GetUserByUserId err", zap.Int("fanUid", fanUid), zap.Error(err))
		return
	}

	if fanUser == nil {
		Logger.Error("fanUser == nil", zap.Int("fanUid", fanUid))
		return
	}

	url := GetUrlUserProfile(fanUid)
	content := fanUser.Name + "关注了你"

	err = UserAddNotice(msg.FollowUid, content, noticeType, url, noticeTime)
	if err != nil {
		Logger.Error("UserAddNotice err", zap.Int("msg.FollowUid", msg.FollowUid), zap.String("content", content), zap.Error(err))
		return
	}

}

// 添加文章通知
func AddNoticeEssayAdd(msg model.NoticeMsg) {
	pageSize := 100

	noticeType := model.NOTICE_TYPE_ESSAY_ADD
	noticeTime := time.Unix(msg.Time, 0)

	authorUid := msg.AuthorUid
	//获取作者信息
	authorUser, err := GetUserByUserId(authorUid)
	if err != nil {
		Logger.Error("GetUserByUserId err", zap.Int("authorUid", authorUid), zap.Error(err))
		return
	}

	//获取文章标题
	essay, err := GetEssayByEid(msg.EssayId)
	if err != nil {
		Logger.Error("GetEssayByEid err", zap.Int("msg.EssayId", msg.EssayId), zap.Error(err))
		return
	}

	essayUrl := GetUrlEssayDetail(essay.CircleId, essay.Id)

	content := "你关注的" + authorUser.Name + "发布了新文章" + essay.Title + "，快去看看吧"

	//分页获取文章作者的所有粉丝
	for page := 1; ; page++ {
		fanUsers, err := GetUserFanListByUid(msg.AuthorUid, page, pageSize)
		if err != nil {
			Logger.Error("GetUserFanListByUid err", zap.Int("msg.AuthorUid", msg.AuthorUid), zap.Int("page", page), zap.Error(err))
			break
		}

		if len(fanUsers) == 0 {
			//当前页没有粉丝了，停止循环
			break
		}

		var uidList []int
		//给每个粉丝发新文章的消息
		for _, fanUser := range fanUsers {
			uidList = append(uidList, fanUser.Id)
		}

		err = UserAddNoticeList(uidList, content, noticeType, essayUrl, noticeTime)
		if err != nil {
			Logger.Error("UserAddNotices err", zap.Any("uidList", uidList), zap.String("content", content), zap.Error(err))
			continue
		}
	}
}

// 添加点赞通知
func AddNoticeLike(msg model.NoticeMsg) {
	noticeType := model.NOTICE_TYPE_LIKE
	noticeTime := time.Unix(msg.Time, 0)

	likeUid := msg.LikeUid
	//获取likeUid的信息
	if likeUid == 0 {
		return
	}

	likeUser, err := GetUserByUserId(likeUid)
	if err != nil {
		Logger.Error("GetUserByUserId err", zap.Int("likeUid", likeUid), zap.Error(err))
		return
	}

	if likeUser == nil {
		Logger.Error("likeUser == nil", zap.Int("likeUid", likeUid))
		return
	}

	if msg.EssayId == 0 {
		return
	}

	essay, err := GetEssayByEid(msg.EssayId)
	if err != nil {
		Logger.Error("GetEssayByEid err", zap.Int("msg.EssayId", msg.EssayId), zap.Error(err))
		return
	}

	if essay == nil {
		Logger.Error("essay == nil", zap.Int("msg.EssayId", msg.EssayId))
		return
	}

	userUrl := GetUrlUserProfile(likeUser.Id)
	content := likeUser.Name + "点赞了你的文章" + essay.Title

	err = UserAddNotice(essay.AuthorId, content, noticeType, userUrl, noticeTime)
	if err != nil {
		Logger.Error("UserAddNotice err", zap.Int("essay.AuthorId", essay.AuthorId), zap.String("content", content), zap.Error(err))
		return
	}

}

// 添加评论通知
func AddNoticeComment(msg model.NoticeMsg) {
	noticeType := model.NOTICE_TYPE_COMMENT
	noticeTime := time.Unix(msg.Time, 0)

	commentUid := msg.CommentUid
	//获取commentUid的信息
	if commentUid == 0 {
		return
	}

	commentUser, err := GetUserByUserId(commentUid)
	if err != nil {
		Logger.Error("GetUserByUserId err", zap.Int("commentUid", commentUid), zap.Error(err))
		return
	}

	if commentUser == nil {
		Logger.Error("commentUser == nil", zap.Int("commentUid", commentUid))
		return
	}

	if msg.EssayId == 0 {
		return
	}

	essay, err := GetEssayByEid(msg.EssayId)
	if err != nil {
		Logger.Error("GetEssayByEid err", zap.Int("msg.EssayId", msg.EssayId), zap.Error(err))
		return
	}

	if essay == nil {
		Logger.Error("essay == nil", zap.Int("msg.EssayId", msg.EssayId))
		return
	}

	essayUrl := GetUrlEssayDetail(essay.CircleId, essay.Id)
	content := commentUser.Name + "评论了你的文章" + essay.Title

	err = UserAddNotice(essay.AuthorId, content, noticeType, essayUrl, noticeTime)
	if err != nil {
		Logger.Error("UserAddNotice err", zap.Int("essay.AuthorId", essay.AuthorId), zap.String("content", content), zap.Error(err))
		return
	}

}

// 添加用户加入圈子通知
func AddNoticeJoin(msg model.NoticeMsg) {
	noticeType := model.NOTICE_TYPE_JOIN
	noticeTime := time.Unix(msg.Time, 0)

	joinUid := msg.JoinUid
	//获取joinUid的信息
	if joinUid == 0 {
		return
	}

	joinUser, err := GetUserByUserId(joinUid)
	if err != nil {
		Logger.Error("GetUserByUserId err", zap.Int("joinUid", joinUid), zap.Error(err))
		return
	}

	if joinUser == nil {
		Logger.Error("joinUser == nil", zap.Int("joinUid", joinUid))
		return
	}

	if msg.CircleId == 0 {
		return
	}

	circle, err := GetCircleByCid(msg.CircleId)
	if err != nil {
		Logger.Error("GetCircleByCid err", zap.Int("msg.CircleId", msg.CircleId), zap.Error(err))
		return
	}

	if circle == nil {
		Logger.Error("circle == nil", zap.Int("msg.CircleId", msg.CircleId))
		return
	}

	userUrl := GetUrlUserProfile(joinUid)
	content := joinUser.Name + "加入了你的圈子" + circle.Title

	err = UserAddNotice(circle.CircleOwnerId, content, noticeType, userUrl, noticeTime)
	if err != nil {
		Logger.Error("UserAddNotice err", zap.Int("circle.CircleOwnerId", circle.CircleOwnerId), zap.String("content", content), zap.Error(err))
		return
	}

}

// 添加举报通知
func AddNoticeAccusation(msg model.NoticeMsg) {
	noticeType := model.NOTICE_TYPE_ACCUSATION
	noticeTime := time.Unix(msg.Time, 0)

	accusationUid := msg.AccusationUid
	//获取accusationUid的信息
	if accusationUid == 0 {
		return
	}

	accusationUser, err := GetUserByUserId(accusationUid)
	if err != nil {
		Logger.Error("GetUserByUserId err", zap.Int("accusationUid", accusationUid), zap.Error(err))
		return
	}

	if accusationUser == nil {
		Logger.Error("accusationUser == nil", zap.Int("accusationUid", accusationUid))
		return
	}

	accusationId := msg.AccusationId
	//获取accusationId的信息
	if accusationId == 0 {
		return
	}

	accusation, err := GetAccusationByAid(accusationId)
	if err != nil {
		Logger.Error("GetAccusationByAid err", zap.Int("accusationId", accusationId), zap.Error(err))
		return
	}

	if accusation == nil {
		Logger.Error("accusation == nil", zap.Int("accusationId", accusationId))
		return
	}

	if accusation.EssayId == 0 {
		return
	}

	essay, err := GetEssayByEid(accusation.EssayId)
	if err != nil {
		Logger.Error("GetEssayByEid err", zap.Int("accusation.EssayId", accusation.EssayId), zap.Error(err))
		return
	}

	if essay == nil {
		Logger.Error("essay == nil", zap.Int("msg.EssayId", msg.EssayId))
		return
	}

	accusationUrl := GetUrlAccusationDetail(accusationId)

	content := "举报的文章" + essay.Title + "无违规" + accusation.Content

	err = UserAddNotice(accusationUid, content, noticeType, accusationUrl, noticeTime)
	if err != nil {
		Logger.Error("UserAddNotice err", zap.Int("accusationUid", accusationUid), zap.String("content", content), zap.Error(err))
		return
	}
}

// 添加被举报通知
func AddNoticeAccusationed(msg model.NoticeMsg) {
	noticeType := model.NOTICE_TYPE_ACCUSATIONED
	noticeTime := time.Unix(msg.Time, 0)

	accusationId := msg.AccusationId
	//获取accusationId的信息
	if accusationId == 0 {
		return
	}

	accusation, err := GetAccusationByAid(accusationId)
	if err != nil {
		Logger.Error("GetAccusationByAid err", zap.Int("accusationId", accusationId), zap.Error(err))
		return
	}

	if accusation == nil {
		Logger.Error("accusation == nil", zap.Int("accusationId", accusationId))
		return
	}

	if accusation.EssayId == 0 {
		return
	}

	essay, err := GetEssayByEid(accusation.EssayId)
	if err != nil {
		Logger.Error("GetEssayByEid err", zap.Int("accusation.EssayId", accusation.EssayId), zap.Error(err))
		return
	}

	if essay == nil {
		Logger.Error("essay == nil", zap.Int("msg.EssayId", msg.EssayId))
		return
	}

	accusationUrl := GetUrlAccusationDetail(accusationId)

	content := "您的文章" + essay.Title + "存在违规" + accusation.Content

	err = UserAddNotice(essay.AuthorId, content, noticeType, accusationUrl, noticeTime)
	if err != nil {
		Logger.Error("UserAddNotice err", zap.Int("essay.AuthorId", essay.AuthorId), zap.String("content", content), zap.Error(err))
		return
	}
}

// 添加反馈通知
func AddNoticeFeedback(msg model.NoticeMsg) {
	noticeType := model.NOTICE_TYPE_FEEDBACK
	noticeTime := time.Unix(msg.Time, 0)

	feedbackId := msg.FeedbackId
	//获取feedbackId的信息
	if feedbackId == 0 {
		return
	}

	feedback, err := GetFeedbackById(feedbackId)
	if err != nil {
		Logger.Error("GetFeedbackByDid err", zap.Int("feedbackId", feedbackId), zap.Error(err))
		return
	}

	if feedback == nil {
		Logger.Error("feedback == nil", zap.Int("feedbackId", feedbackId))
		return
	}

	feedbackUrl := GetUrlFeedbackDetail(feedback.Id)
	content := feedback.Reply + feedback.Content

	err = UserAddNotice(feedback.UserId, content, noticeType, feedbackUrl, noticeTime)
	if err != nil {
		Logger.Error("UserAddNotice err", zap.Int("feedback.UserId", feedback.UserId), zap.String("content", content), zap.Error(err))
		return
	}
}

// 加精
func AddNoticeEssence(msg model.NoticeMsg) {
	noticeType := model.NOTICE_TYPE_ESSENCE
	noticeTime := time.Unix(msg.Time, 0)

	eid := msg.EssayId

	if eid == 0 {
		return
	}

	essay, err := GetEssayByEid(eid)
	if err != nil {
		Logger.Error("GetEssayByEid err", zap.Int("eid", eid), zap.Error(err))
		return
	}

	if essay == nil {
		Logger.Error("essay == nil", zap.Int("eid", eid))
		return
	}

	url := GetUrlEssayDetail(essay.CircleId, eid)
	content := "你的文章" + essay.Title + "被加精"

	err = UserAddNotice(essay.AuthorId, content, noticeType, url, noticeTime)
	if err != nil {
		Logger.Error("UserAddNotice err", zap.Int("essay.AuthorId", essay.AuthorId), zap.String("content", content), zap.Error(err))
		return
	}
}

func GetNoticeTypeNameList() (typeNameList []model.NoticeTypeName) {
	typeNameList = []model.NoticeTypeName{
		{model.NOTICE_TYPE_OTHER, "全部"},
		{model.NOTICE_TYPE_FOLLOW, "关注"},
		{model.NOTICE_TYPE_LIKE, "点赞"},
		{model.NOTICE_TYPE_COMMENT, "评论"},
		{model.NOTICE_TYPE_ACCUSATION, "举报"},
		{model.NOTICE_TYPE_ACCUSATIONED, "被举报"},
		{model.NOTICE_TYPE_FEEDBACK, "反馈"},
		{model.NOTICE_TYPE_ESSAY_ADD, "关注者发布文章"},
		{model.NOTICE_TYPE_ESSENCE, "被加精"},
		{model.NOTICE_TYPE_JOIN, "加入圈子"},
	}

	return typeNameList
}
