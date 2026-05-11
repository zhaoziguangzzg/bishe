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
	/*
		case model.NOTICE_TYPE_LIKE:
				//noticeContent = "又有用户" + userName + "点赞啦"
			case model.NOTICE_TYPE_COMMENT:
				//noticeContent = "又有用户" + userName + "评论啦"

			case model.NOTICE_TYPE_JOIN:
				noticeContent = "又有用户" + userName + "加入圈子啦"
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
