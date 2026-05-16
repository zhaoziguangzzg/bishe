package mysql

import (
	"bishe/model"

	"gorm.io/gorm"
)

// create圈子
func CreateFavorite(newFavorite *model.Favorite) (err error) {
	err = DB.Model(&model.Favorite{}).Create(newFavorite).Error
	return
}

// 获取用户全部的收藏夹
func GetAllFavoriteByUid(uid int, page int, pagesize int) (favorites []model.Favorite, err error) {
	offset := (page - 1) * pagesize

	err = DB.Model(&model.Favorite{}).Where("user_id=? and is_deleted=?", uid, model.FAVORITE_NOT_DELETED).
		Order("id DESC").Offset(offset).Limit(pagesize).Find(&favorites).Error
	if err != nil {
		return
	}

	return
}

// 根据fid获取收藏夹
func GetFavoriteByFid(fid int) (favorite *model.Favorite, err error) {
	favorite = new(model.Favorite)
	err = DB.Model(&model.Favorite{}).Where("id=? and is_deleted=?", fid, model.FAVORITE_NOT_DELETED).First(&favorite).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return favorite, nil
}

// 根据title获取收藏夹
func GetFavoriteByTitle(title string, uid int) (favorite *model.Favorite, err error) {
	favorite = new(model.Favorite)
	err = DB.Model(&model.Favorite{}).Where("title=? and user_id=? and is_deleted=?", title, uid, model.FAVORITE_NOT_DELETED).
		First(&favorite).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return favorite, nil
}

// 根据fid更新收藏夹标题
func UpdateFavoriteTitleByFid(fid int, title string) (int64, error) {
	result := DB.Model(&model.Favorite{}).Where("id=?", fid).Update("title", title)
	return result.RowsAffected, result.Error
}

// 更新IsDeleted删除收藏夹
func UpdateFavoriteIsDeleted(fid int) (int64, error) {
	result := DB.Model(&model.Favorite{}).Where("id=?", fid).Update("is_deleted", model.FAVORITE_IS_DELETED)
	return result.RowsAffected, result.Error
}
