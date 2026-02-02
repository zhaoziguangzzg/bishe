package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"

	"gorm.io/gorm"
)

// 用户创建消息
func UserAddInformation(information *model.Information) (err error) {
	err = DB.Model(&model.Information{}).Create(information).Error
	return
}

// 获取用户消息
func GetInformationByUname(uname string) (information *model.Information, err error) {
	information = new(model.Information)
	err = DB.Model(&model.Information{}).Where("receive_name=? and is_deleted=?", uname, model.ESSAY_NOT_DELETED).First(&information).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return information, nil
}

// 通知
func AddUserNotice(notice *model.Information) (err error) {
	err = DB.Model(&model.Information{}).Create(notice).Error
	return
}

// 更新IsDeleted删除information
func UpdateInformationIsDeleted(iid int) (int64, error) {
	result := DB.Model(&model.Information{}).Where("id=?", iid).Update("is_deleted", model.INFORMATION_IS_DELETED)
	return result.RowsAffected, result.Error
}
