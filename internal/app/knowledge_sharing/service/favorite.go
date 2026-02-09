package service

import (
	"bishe/internal/app/knowledge_sharing/dao/mysql"
	"bishe/internal/app/knowledge_sharing/model"
)

// create收藏夹
func CreateFavorite(newFavorite *model.Favorite) (err error) {
	return mysql.CreateFavorite(newFavorite)
}

// 获取用户全部的收藏夹
func GetAllFavoriteByUid(uid int, page int, pagesize int) (favorites []model.Favorite, err error) {
	return mysql.GetAllFavoriteByUid(uid, page, pagesize)
}

// 根据title获取收藏夹
func GetFavoriteByTitle(title string, uid int) (favorite *model.Favorite, err error) {
	return mysql.GetFavoriteByTitle(title, uid)
}

// 根据fid获取收藏夹
func GetFavoriteByFid(fid int) (favorite *model.Favorite, err error) {
	return mysql.GetFavoriteByFid(fid)
}

// 根据fid更新收藏夹标题
func UpdateFavoriteTitleByFid(fid int, title string) (int64, error) {
	return mysql.UpdateFavoriteTitleByFid(fid, title)
}

// 更新IsDeleted删除收藏夹
func UpdateFavoriteIsDeleted(fid int) (int64, error) {
	return mysql.UpdateFavoriteIsDeleted(fid)
}
