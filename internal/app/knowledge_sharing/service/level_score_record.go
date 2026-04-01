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

// 获取用户在圈子全部等级详情
func GetUserOfCircleLevelRecordList(uid int, cid int) (levelScoreRecords []model.LevelScoreRecord, err error) {
	return mysql.GetUserOfCircleLevelRecordList(uid, cid)
}
