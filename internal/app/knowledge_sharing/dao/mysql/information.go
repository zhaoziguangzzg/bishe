package mysql

import "bishe/internal/app/knowledge_sharing/model"

//用户创建消息
func UserAddInformation(information *model.Information) (err error) {
	err = DB.Model(&model.Information{}).Create(information).Error
	return
}

//通知
func AddUserNotice(notice *model.Information) (err error) {
	err = DB.Model(&model.Information{}).Create(notice).Error
	return
}
