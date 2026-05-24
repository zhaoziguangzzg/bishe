package mysql

import (
	"bishe/model"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	err = DB.Model(&model.LevelScore{}).Clauses(clause.OnConflict{DoNothing: true}).Create(levelScore).Error
	return
}

// 获取用户等级分数
func GetUserLevelScoreByUidCid(uid int, cid int) (levelScore *model.LevelScore, err error) {
	levelScore = new(model.LevelScore)

	err = DB.Model(&model.LevelScore{}).
		Where("uid=? and cid=?", uid, cid).
		First(&levelScore).Error
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

func UpdateLevelScoreByUidCid(uid int, cid int, score int, now time.Time) (int64, error) {
	levelScore := &model.LevelScore{
		Uid:       uid,
		Cid:       cid,
		Score:     score,
		CreateAt:  &now,
		UpdateAt:  &now,
		IsDeleted: model.IS_DELETED_NO,
	}

	result := DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "uid"}, {Name: "cid"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"score":     gorm.Expr("level_score.score + ?", score),
			"update_at": &now,
		}),
	}).Create(levelScore)

	return result.RowsAffected, result.Error
}
