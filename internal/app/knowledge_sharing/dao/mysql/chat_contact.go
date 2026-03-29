package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 添加或更新联系人
func ChatContactInsertUpdate(chatContact *model.ChatContact) (err error) {

	unionUid := model.MakeChatContactUnionUid(chatContact.SendUid, chatContact.ReceiveUid)
	chatContact.UnionUid = unionUid

	err = DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "union_uid"}},          // 冲突检测列
		DoUpdates: clause.AssignmentColumns([]string{"content"}), // 更新字段
	}).Create(&chatContact).Error

	return
}

// 获取联系人列表
func GetChatContactList(uid int, page int, pageSize int) (chatContacts []model.ChatContact, err error) {
	offset := (page - 1) * pageSize

	err = DB.Model(&model.ChatContact{}).Where("send_uid=? or receive_uid=? and is_deleted=?", uid, uid, model.IS_DELETED_NO).
		Order("id DESC").Offset(offset).Limit(pageSize).Find(&chatContacts).Error
	if err != nil {
		return
	}

	return
}

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

// 根据id获取联系人
func GetUserContactById(id int) (contact *model.Contact, err error) {
	contact = new(model.Contact)
	err = DB.Model(&model.Contact{}).Where("id=?", id).First(&contact).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return contact, nil
}

// 获取用户全部联系人
func GetUserAllContact(uid int, page int, pagesize int) (contacts []model.Contact, err error) {
	offset := (page - 1) * pagesize
	err = DB.Model(&model.Contact{}).Where("send_id=? and is_deleted=?", uid, model.CONTACT_NOT_DELETED).
		Order("id DESC").Offset(offset).Limit(pagesize).Find(&contacts).Error
	if err != nil {
		return
	}

	return
}

// 删除联系人
func DeleteUserContactByReceiveId(uid int, receiveId int) (int64, error) {
	result := DB.Model(&model.Contact{}).Where("send_id and receive_id=?", uid, receiveId).
		Update("is_deleted", model.CONTACT_IS_DELETED)
	return result.RowsAffected, result.Error
}
