package mysql

import (
	"bishe/internal/app/knowledge_sharing/model"

	"gorm.io/gorm"
)

// create用户
func CreateUser(newUser *model.User) (err error) {
	err = DB.Model(&model.User{}).Create(newUser).Error
	return
}

// 更新
func UpdateUserByUid(uid int, updateMap map[string]interface{}) (int64, error) {

	// user := model.User{
	// 	Name:  name,
	// 	Email: email,
	// 	Age:   age,
	// 	Phone: phone,
	// }
	result := DB.Model(&model.User{}).Where("id=?", uid).Updates(updateMap)
	return result.RowsAffected, result.Error
}

// 更新用户密码
func UpdateUserPasswordByUid(uid int, password string) (int64, error) {

	result := DB.Model(&model.User{}).Where("id=?", uid).Update("password", password)
	return result.RowsAffected, result.Error
}

// 根据name获取用户
func GetUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	err = DB.Model(&model.User{}).Where("name=? and is_deleted=?", name, model.USER_NOT_DELETED).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

// 根据uid获取用户
func GetUserByUserId(UserId int) (user *model.User, err error) {
	user = new(model.User)
	err = DB.Model(&model.User{}).Where("id=? and is_deleted=?", UserId, model.USER_NOT_DELETED).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

// 根据uids获取userMap
func GetUserMapByUids(uids []int) (userMap map[int]model.User, err error) {
	users := make([]model.User, 0)
	err = DB.Model(&model.User{}).Where("id IN (?)", uids).Find(&users).Error
	if err != nil {
		return
	}

	userMap = make(map[int]model.User, 0)
	for _, v := range users {
		userMap[v.Id] = v
	}

	return
}
