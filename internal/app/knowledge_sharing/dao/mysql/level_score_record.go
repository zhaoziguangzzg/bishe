package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"
	"time"

	"gorm.io/gorm"
)

// 增加等级分数详情
func UserAddLevelScoreRecord(uid int, cid int, score int, relateId int, typei int, createAt time.Time) (err error) {
	levelScoreRecord := &model.LevelScoreRecord{
		Uid:       uid,
		Cid:       cid,
		Score:     score,
		RelateId:  relateId,
		Typei:     typei,
		CreateAt:  &createAt,
		UpdateAt:  &createAt,
		IsDeleted: model.IS_DELETED_NO,
	}

	err = DB.Model(&model.LevelScoreRecord{}).Create(levelScoreRecord).Error
	return
}

// 更新等级分数详情
func UserUpdateLevelScoreRecord(uid int, cid int, typei int, isDeleted int) (err error) {
	err = DB.Model(&model.LevelScoreRecord{}).Where("uid=? and cid=? and type=?", uid, cid, typei).
		Update("is_deleted", isDeleted).Error
	return
}

// 获取用户等级分数详情
func GetUserLevelScoreRecordByUidCidType(uid int, cid int, typei int) (levelScoreRecord *model.LevelScoreRecord, err error) {
	levelScoreRecord = new(model.LevelScoreRecord)

	err = DB.Model(&model.LevelScoreRecord{}).Where("uid=? and cid=? and type=?", uid, cid, typei).
		First(&levelScoreRecord).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return levelScoreRecord, nil
}

// 获取用户在圈子全部等级详情
func GetUserOfCircleLevelRecordList(uid int, cid int) (levelScoreRecords []model.LevelScoreRecord, err error) {

	err = DB.Model(&model.LevelScoreRecord{}).Where("uid=? and cid=? and is_deleted=?", uid, cid, model.IS_DELETED_NO).
		Order("id DESC").Find(&levelScoreRecords).Error
	if err != nil {
		return
	}

	return
}
