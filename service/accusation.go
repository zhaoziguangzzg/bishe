package service

import (
	"bishe/dao/mysql"
	"bishe/model"
)

// create 用户举报
func CreateUserAccusation(newAccusation *model.Accusation) (err error) {
	return mysql.CreateUserAccusation(newAccusation)
}

// 根据uid,eid获取举报文章
func GetUserAccusationEssay(uid int, eid int) (accusation *model.Accusation, err error) {
	return mysql.GetUserAccusationEssay(uid, eid)
}

// 获取全部举报（支持按状态筛选，status为-1时查询所有）
func GetAllAccusationEssay(page int, pagesize int, status int) (accusations []model.Accusation, err error) {
	return mysql.GetAllAccusationEssay(page, pagesize, status)
}

// 获取文章举报内容
func GetAccusationByAid(aid int) (accusation *model.Accusation, err error) {
	return mysql.GetAccusationByAid(aid)
}

// 更新举报信息为无违规
func UpdateAccusationNormalByAid(aid int) (int64, error) {
	return mysql.UpdateAccusationNormalByAid(aid)
}

// 更新举报信息为有违规
func UpdateAccusationViolateByAid(aid int) (int64, error) {
	return mysql.UpdateAccusationViolateByAid(aid)
}
