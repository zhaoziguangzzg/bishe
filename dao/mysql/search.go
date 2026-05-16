package mysql

import (
	"bishe/model"

	"gorm.io/gorm"
)

// create搜索
func CreateSearch(search *model.Search) (err error) {
	err = DB.Model(&model.Search{}).Create(search).Error
	return
}

// 更新搜索圈子searchnum 增加
func IncrUpdateCircleSearchNum(cid int) (int64, error) {
	result := DB.Model(&model.Circle{}).Where("id=?", cid).UpdateColumn("join_num", gorm.Expr("join_num + ?", 1))
	return result.RowsAffected, result.Error
}
