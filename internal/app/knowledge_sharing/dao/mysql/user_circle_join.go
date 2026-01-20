package mysql

import "bishe/internal/app/knowledge_sharing/model"

//create 用户加入圈子
func CreateUserCircleJoin(newUserCircle *model.UserCircleJoin) (err error) {
	err = DB.Model(&model.UserCircleJoin{}).Create(newUserCircle).Error
	return
}
