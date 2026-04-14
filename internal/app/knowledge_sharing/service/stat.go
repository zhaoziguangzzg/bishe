package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
	"time"
)

// 添加或更新各类型用户数据
func StatInsertUpdate(statUid int, num int, typei int, createTime time.Time) (err error) {
	return mysql.StatInsertUpdate(statUid, num, typei, createTime)
}

// 获取用户数据列表
func GetUserStatList(uid int) (stats []model.Stat, err error) {
	return mysql.GetUserStatList(uid)
}

// 获取用户数据map
func GetUserStatMapByType(uid int) (userStatMap map[int]int, err error) {
	return mysql.GetUserStatMapByType(uid)
}

// 添加各类型数据详情
func StatDetailsInsert(statUid int, typei int, createTime time.Time) (err error) {
	return mysql.StatDetailsInsert(statUid, typei, createTime)
}

// 获取近期各类型数据
func GetStatDetailsByType(uid int, stime time.Time) (results []model.StatDetailsTypeCount, err error) {
	return mysql.GetStatDetailsByType(uid, stime)
}

// 更新数据总数和详情
func UpdateStatAndStatDetail(uid int, typei int, createTime time.Time) (err error) {
	num := 1
	//更新数据统计
	err = StatInsertUpdate(uid, num, typei, createTime)
	if err != nil {
		return
	}

	//添加数据详情
	err = StatDetailsInsert(uid, typei, createTime)
	if err != nil {
		return
	}
	return
}
