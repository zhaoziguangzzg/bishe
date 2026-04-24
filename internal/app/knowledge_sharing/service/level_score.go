package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
	"time"
)

// 添加等级
func UserAddLevelScore(uid int, cid int, createAt time.Time) (err error) {
	return mysql.UserAddLevelScore(uid, cid, createAt)
}

// 获取用户等级
func GetUserLevelScoreByUidCid(uid int, cid int) (levelScore *model.LevelScore, err error) {
	return mysql.GetUserLevelScoreByUidCid(uid, cid)
}

// 根据uids获取LevelMap
func GetLevelScoreMapByUids(uids []int, cid int) (levelScoreMap map[int]model.LevelScore, err error) {
	return mysql.GetLevelScoreMapByUids(uids, cid)
}

// 更新等级分数 增加
func UpdateLevelScoreByUidCid(uid int, cid int, score int) (int64, error) {
	return mysql.UpdateLevelScoreByUidCid(uid, cid, score)
}
