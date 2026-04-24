package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
	"time"
)

// 增加等级分数详情
func UserAddLevelScoreRecord(uid int, cid int, score int, relateId int, typei int, createAt time.Time) (err error) {
	return mysql.UserAddLevelScoreRecord(uid, cid, score, relateId, typei, createAt)
}

// 更新等级分数详情
func UserUpdateLevelScoreRecord(uid int, cid int, typei int, isDeleted int) (err error) {
	return mysql.UserUpdateLevelScoreRecord(uid, cid, typei, isDeleted)
}

// 获取用户等级分数详情
func GetUserLevelScoreRecordByUidCidType(uid int, cid int, typei int) (levelScoreRecord *model.LevelScoreRecord, err error) {
	return mysql.GetUserLevelScoreRecordByUidCidType(uid, cid, typei)
}

// 获取用户在圈子全部等级详情
func GetUserOfCircleLevelRecordList(uid int, cid int) (levelScoreRecords []model.LevelScoreRecord, err error) {
	return mysql.GetUserOfCircleLevelRecordList(uid, cid)
}
