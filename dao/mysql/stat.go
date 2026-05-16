package mysql

import (
	"bishe/model"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 添加或更新各类型用户数据
func StatInsertUpdate(statUid int, num int, typei int, createTime time.Time) (err error) {
	// 确保初始值不为负数
	initialSum := num
	if initialSum < 0 {
		initialSum = 0
	}

	stat := &model.Stat{
		StatUid:   statUid,
		Sum:       initialSum,
		Type:      typei,
		CreateAt:  &createTime,
		UpdateAt:  &createTime,
		IsDeleted: model.IS_DELETED_NO,
	}

	err = DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "stat_uid"},
			{Name: "type"},
		}, // 冲突检测列
		DoUpdates: clause.Assignments(map[string]interface{}{
			"sum": gorm.Expr("GREATEST(0, sum + ?)", num),
		}),
	}).Create(&stat).Error

	return
}

// // 添加各类型用户数据
// func AddUserStat(statUid int, num int, typei int, createTime time.Time) (err error) {
// 	stat := &model.Stat{
// 		StatUid:   statUid,
// 		Sum:       num,
// 		Type:      typei,
// 		CreateAt:  &createTime,
// 		UpdateAt:  &createTime,
// 		IsDeleted: model.IS_DELETED_NO,
// 	}

// 	err = DB.Model(&model.Stat{}).Create(stat).Error
// 	return

// }

// 获取用户数据列表
func GetUserStatList(uid int) (stats []model.Stat, err error) {

	err = DB.Model(&model.Stat{}).Where("stat_uid=? and is_deleted=?", uid, model.IS_DELETED_NO).
		Order("id DESC").Find(&stats).Error
	if err != nil {
		return
	}

	return
}

// 获取用户数据map
func GetUserStatMapByType(uid int) (userStatMap map[int]int, err error) {

	stats := make([]model.Stat, 0)
	err = DB.Model(&model.Stat{}).Where("stat_uid=? and is_deleted=?", uid, model.IS_DELETED_NO).
		Order("id DESC").Find(&stats).Error
	if err != nil {
		return
	}

	userStatMap = make(map[int]int, 0)
	for _, v := range stats {
		userStatMap[v.Type] = v.Sum
	}

	return
}

// 添加各类型数据详情
func StatDetailsInsert(statUid int, typei int, statStatus int, createTime time.Time) (err error) {
	statDetail := &model.StatDetails{
		StatUid:    statUid,
		Type:       typei,
		StatStatus: statStatus,

		CreateAt:  &createTime,
		UpdateAt:  &createTime,
		IsDeleted: model.IS_DELETED_NO,
	}

	err = DB.Model(&model.StatDetails{}).Create(statDetail).Error

	return
}

// 获取近期各类型数据
func GetStatDetailsByType(uid int, stime time.Time) (results []model.StatDetailsTypeCount, err error) {

	err = DB.Model(&model.StatDetails{}).Select("type, COUNT(*) AS total").
		Where("stat_uid = ? AND is_deleted = ? AND create_at > ?", uid, model.IS_DELETED_NO, stime).
		Group("type").Find(&results).Error
	if err != nil {
		return
	}

	return
}

// 获取近期各类型按日期统计数据
func GetStatDetailsByDateType(uid int, stime time.Time) (results []model.StatDetailsDateCount, err error) {

	err = DB.Model(&model.StatDetails{}).
		Select("type, DATE(create_at) as date, COUNT(*) AS total").
		Where("stat_uid = ? AND is_deleted = ? AND create_at > ?", uid, model.IS_DELETED_NO, stime).
		Group("type, DATE(create_at)").
		Order("date ASC").
		Find(&results).Error
	if err != nil {
		return
	}

	return
}
