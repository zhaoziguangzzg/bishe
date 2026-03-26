package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"

	"gorm.io/gorm"
)

// create 用户联系人
func CreateUserContact(newContect *model.Contact) (err error) {
	err = DB.Model(&model.Contact{}).Create(newContect).Error
	return
}

// 根据uid,receiveId获取联系人
func GetUserContact(uid int, receiveId int) (contact *model.Contact, err error) {
	contact = new(model.Contact)
	err = DB.Model(&model.Contact{}).
		Where("send_id=? and contact_id=?", uid, receiveId).First(&contact).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return contact, nil
}
