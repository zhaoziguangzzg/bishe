package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"

	"gorm.io/gorm"
)

// create圈子
func CreateCircle(newCircle *model.Circle) (err error) {
	err = DB.Model(&model.Circle{}).Create(newCircle).Error
	return
}

// 根据cid获取圈子
func GetCircleByCid(cid int) (circle *model.Circle, err error) {
	circle = new(model.Circle)
	err = DB.Model(&model.Circle{}).Where("id=?", cid).First(&circle).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return circle, nil
}

// get 付费圈子
func GetCircleAllCharge() (circles []model.Circle, err error) {
	err = DB.Model(&model.Circle{}).Where("price>?", 0).Find(&circles).Error
	if err != nil {
		return
	}

	return
}

// get 免费圈子
func GetCricleAllFree() (circles []model.Circle, err error) {
	err = DB.Model(&model.Circle{}).Where("price=?", 0).Find(&circles).Error
	if err != nil {
		return
	}

	return
}

// get all圈子
func GetCircleAllByJoinNum(page int, pagesize int) (circles []model.Circle, err error) {
	offset := (page - 1) * pagesize

	err = DB.Model(&model.Circle{}).Order("join_num DESC").Offset(offset).Limit(pagesize).Find(&circles).Error
	if err != nil {
		return
	}

	return
}

// get 用户创建的圈子
func GetUserCreateCircleByUid(uid int, page int, pagesize int) (circles []model.Circle, err error) {
	offset := (page - 1) * pagesize
	err = DB.Model(&model.Circle{}).Where("circle_own_id=?", uid).Order("join_num DESC").Offset(offset).Limit(pagesize).Find(&circles).Error
	if err != nil {
		return
	}

	return
}

// get 用户加入的圈子
func GetUserJoinCircleListByUid(uid int, page int, pagesize int) (circles []model.Circle, err error) {
	var cids []int
	offset := (page - 1) * pagesize

	err = DB.Model(&model.UserCircleJoin{}).Where("user_id=? and not_join_status=?", uid, model.
		USER_CIRCLE_NOT_NO_JOIN).Order("join_time DESC").Offset(offset).Limit(pagesize).Pluck("circle_id", &cids).Error
	if err != nil {
		return
	}

	err = DB.Where("id IN (?)", cids).Find(&circles).Error
	if err != nil {
		return
	}

	return

}

// 更新圈子join num 增加
func IncrUpdateCircleJoinNumByCid(cid int) (int64, error) {
	result := DB.Model(&model.Circle{}).Where("id=?", cid).UpdateColumn("join_num", gorm.Expr("join_num + ?", 1))
	return result.RowsAffected, result.Error
}

// 更新圈子join num 减少
func DecrrUpdateCircleJoinNumByCid(cid int) (int64, error) {
	result := DB.Model(&model.Circle{}).Where("id=?", cid).UpdateColumn("join_num", gorm.Expr("join_num - ?", 1))
	return result.RowsAffected, result.Error
}
