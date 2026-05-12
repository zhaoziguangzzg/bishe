package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"
	"time"
)

// 添加通知
func UserAddNotice(noticeUid int, content string, typei int, url string, createTime time.Time) (err error) {
	notice := &model.Notice{
		NoticeUid: noticeUid,
		Content:   content,
		Type:      typei,
		Url:       url,
		CreateAt:  &createTime,
		UpdateAt:  &createTime,
		IsDeleted: model.IS_DELETED_NO,
	}

	err = DB.Model(&model.Notice{}).Create(notice).Error
	return
}

func UserAddNotices(noticeUids []int, content string, typei int, createTime time.Time) (err error) {
	var notices []*model.Notice
	for _, v := range noticeUids {
		noticeUid := v
		notice := &model.Notice{
			NoticeUid: noticeUid,
			Content:   content,
			Type:      typei,
			CreateAt:  &createTime,
			UpdateAt:  &createTime,
			IsDeleted: model.IS_DELETED_NO,
		}
		notices = append(notices, notice)
	}
	err = DB.Model(&model.Notice{}).Create(notices).Error
	return
}

// 获取通知列表
func GetNoticeList(uid int, page int, pageSize int) (notices []model.Notice, err error) {
	offset := (page - 1) * pageSize

	err = DB.Model(&model.Notice{}).Where("notice_uid=? and is_deleted=?", uid, model.IS_DELETED_NO).
		Order("id DESC").Offset(offset).Limit(pageSize).Find(&notices).Error
	if err != nil {
		return
	}

	return
}

// 根据类型获取通知列表
func GetNoticeListByType(uid int, typei int, page int, pageSize int) (notices []model.Notice, err error) {
	offset := (page - 1) * pageSize

	err = DB.Model(&model.Notice{}).Where("notice_uid=? and type=? and is_deleted=?", uid, typei, model.IS_DELETED_NO).
		Order("id DESC").Offset(offset).Limit(pageSize).Find(&notices).Error
	if err != nil {
		return
	}

	return
}
