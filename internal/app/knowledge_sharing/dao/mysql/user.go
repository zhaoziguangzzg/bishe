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

// 获取用户密码
func GetPasswordFromUser(account int) (password string, err error) {
	err = DB.Model(&model.User{}).Where("account=?", account).Select("password").Scan(&password).Error
	return
}

// 更新
func UpdateFromUser(user *model.User, password string, email string, age int, phone int) (result *gorm.DB) {
	// if password == "" {
	// 	email = user.Password
	// }
	// if email == "" {
	// 	email = user.Email
	// }

	// if age == 0 {
	// 	age = user.Age
	// }

	// if phone == 0 {
	// 	phone = user.Phone
	// }
	result = DB.Model(&model.User{}).Where("id=?", user.Id).Updates(model.User{Password: password, Email: email, Age: age, Phone: phone})
	return
}

// 根据account获取用户
func GetUserByAccount(account int) (user *model.User, err error) {
	err = DB.Model(&model.User{}).Where("account=?", account).Find(user).Error

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
	err = DB.Model(&model.User{}).Where("id=?", UserId).Find(user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { //没查到数据返回空
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}
