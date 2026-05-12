package service

import (
	"bishe/internal/app/knowledge_sharing/dao/kafka"
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
	"encoding/json"
	"time"

	"go.uber.org/zap"
)

// 添加通知
func UserAddNotice(noticeUid int, content string, typei int, createTime time.Time) (err error) {
	return mysql.UserAddNotice(noticeUid, content, typei, createTime)
}

// 添加通知
func UserAddNotices(noticeUids []int, content string, typei int, createTime time.Time) (err error) {
	return mysql.UserAddNotices(noticeUids, content, typei, createTime)
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
	/*



		case model.NOTICE_TYPE_ACCUSATION:
			noticeContent = "用户" + userName + "举报没有违规"
		case model.NOTICE_TYPE_ACCUSATIONED:
			noticeContent = "用户" + userName + "举报有违规"
		case model.NOTICE_TYPE_FEEDBACK:
			noticeContent = "用户" + userName + "反馈了"
	*/
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

	userUrl := GetUrlUserProfile(fanUid)
	content := fanUser.Name + "关注了你" + userUrl

	err = UserAddNotice(msg.FollowUid, content, noticeType, noticeTime)
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

	essayUrl := GetUrlEssayDetail(essay.Id)

	content := "你关注的" + authorUser.Name + "发布了新文章" + essay.Title + "，快去看看吧，地址：" + essayUrl

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

		//给每个粉丝发新文章的消息
		for _, fanUser := range fanUsers {
			err = UserAddNotice(fanUser.Id, content, noticeType, noticeTime)
			if err != nil {
				Logger.Error("UserAddNotice err", zap.Int("fanUser.Id", fanUser.Id), zap.String("content", content), zap.Error(err))
				continue
			}
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

	userUrl := GetUrlUserProfile(likeUid)
	essayUrl := GetUrlEssayDetail(essay.Id)
	content := likeUser.Name + userUrl + "点赞了你的文章" + essay.Title + essayUrl

	err = UserAddNotice(essay.AuthorId, content, noticeType, noticeTime)
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

	userUrl := GetUrlUserProfile(commentUid)
	essayUrl := GetUrlEssayDetail(essay.Id)
	content := commentUser.Name + userUrl + "评论了你的文章" + essay.Title + essayUrl

	err = UserAddNotice(essay.AuthorId, content, noticeType, noticeTime)
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
	circleUrl := GetUrlCircleIndex(circle.Id)
	content := joinUser.Name + userUrl + "加入了你的圈子" + circle.Title + circleUrl

	err = UserAddNotice(circle.CircleOwnerId, content, noticeType, noticeTime)
	if err != nil {
		Logger.Error("UserAddNotice err", zap.Int("circle.CircleOwnerId", circle.CircleOwnerId), zap.String("content", content), zap.Error(err))
		return
	}

}
