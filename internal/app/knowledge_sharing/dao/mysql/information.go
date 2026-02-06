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
	err = DB.Model(&model.Information{}).Where("receive_name=? and is_deleted=?", uname, model.
		INFORMATION_NOT_DELETED).First(&information).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return information, nil
}

// 获取用户接收消息
func GetReceiveInformationByUid(uid int, sendId int, page int, pageSize int) (informations []model.Information, err error) {
	offset := (page - 1) * pageSize

	err = DB.Model(&model.Information{}).Where("send_id=? and receive_id=? and is_deleted=?", sendId, uid,
		model.INFORMATION_NOT_DELETED).Order("id DESC").Offset(offset).Limit(pageSize).Find(&informations).Error
	if err != nil {
		return
	}

	return
}

// 获取用户发送消息
func GetSendInformationByUid(uid int, receiveId int, page int, pageSize int) (informations []model.Information, err error) {
	offset := (page - 1) * pageSize

	err = DB.Model(&model.Information{}).Where("send_id=? and receive_id=? and is_deleted=?", uid, receiveId,
		model.INFORMATION_NOT_DELETED).Order("id DESC").Offset(offset).Limit(pageSize).Find(&informations).Error
	if err != nil {
		return
	}

	return
}

// 获取用户消息
func GetUserAllInformation(uid int, page int, pageSize int) (informations []model.Information, err error) {
	offset := (page - 1) * pageSize
	err = DB.Model(&model.Information{}).Where("send_id=? or receive_id=? and isdeleted=?", uid, uid, model.INFORMATION_NOT_DELETED).Order("id DESC").
		Offset(offset).Limit(pageSize).Find(&informations).Error
	if err != nil {
		return
	}
	return
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
