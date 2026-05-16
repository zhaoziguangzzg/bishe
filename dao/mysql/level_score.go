package mysql

import (
	"bishe/model"
	"time"

	"gorm.io/gorm"
)

// 添加等级分数
func UserAddLevelScore(uid int, cid int, createAt time.Time) (err error) {
	levelScore := &model.LevelScore{
		Uid:       uid,
		Cid:       cid,
		CreateAt:  &createAt,
		UpdateAt:  &createAt,
		IsDeleted: model.IS_DELETED_NO,
	}
	err = DB.Model(&model.LevelScore{}).Create(levelScore).Error
	return
}

// 获取用户等级分数
func GetUserLevelScoreByUidCid(uid int, cid int) (levelScore *model.LevelScore, err error) {
	levelScore = new(model.LevelScore)

	err = DB.Model(&model.LevelScore{}).Where("user_id=? and circle_id=?", uid, cid).First(&levelScore).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return levelScore, nil
}

// 根据uids获取LevelScoreMap
func GetLevelScoreMapByUids(uids []int, cid int) (levelScoreMap map[int]model.LevelScore, err error) {
	levelScores := make([]model.LevelScore, 0)
	err = DB.Model(&model.LevelScore{}).Where("cid=? and uid IN (?)", cid, uids).Find(&levelScores).Error
	if err != nil {
		return
	}

	levelScoreMap = make(map[int]model.LevelScore, 0)
	for _, v := range levelScores {
		levelScoreMap[v.Uid] = v
	}

	return
}

// 更新等级分数
func UpdateLevelScoreByUidCid(uid int, cid int, score int) (int64, error) {
	now := time.Now()

	// 先尝试更新
	result := DB.Model(&model.LevelScore{}).
		Where("uid=? and cid=?", uid, cid).
		UpdateColumn("score", gorm.Expr("score + ?", score))

	// 如果没有更新任何记录，说明记录不存在，需要创建
	if result.RowsAffected == 0 {
		levelScore := &model.LevelScore{
			Uid:       uid,
			Cid:       cid,
			Score:     score,
			CreateAt:  &now,
			UpdateAt:  &now,
			IsDeleted: model.IS_DELETED_NO,
		}
		err := DB.Model(&model.LevelScore{}).Create(levelScore).Error
		if err != nil {
			return 0, err
		}
		return 1, nil
	}

	return result.RowsAffected, result.Error
}
