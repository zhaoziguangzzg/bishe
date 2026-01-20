package mysql

import "bishe/internal/app/knowledge_sharing/model"

func UserAddInformation(information *model.Information) (err error) {
	err = DB.Model(&model.Information{}).Create(information).Error
	return
}
