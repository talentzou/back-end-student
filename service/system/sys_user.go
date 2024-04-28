package system

import (
	"back-end/global"
	"back-end/model/system"
	"back-end/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserService struct{}

// 获取用户列表
func (userService *UserService) GetUserInfoList(offset int, limit int) (list interface{}, total int64, err error) {
	db := global.Global_Db.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return userList, total, err
}

// 删除用户
func (userService *UserService) DeleteUser(id int) (err error) {
	return global.Global_Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&system.SysUser{}).Error; err != nil {
			return err
		}
		// if err := tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error; err != nil {
		// 	return err
		// }
		return nil
	})
}

// 注册用户
func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.Global_Db.Where("user_name = ?", u.UserName).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	err = global.Global_Db.Create(&u).Error
	return u, err
}

// 搜索用户
func (userService *UserService) QueryUser(offset int, limit int, userName string) (interface{}, int64, error) {
	var userList []system.SysUser
	var total int64
	db := global.Global_Db.Limit(limit).Offset(offset).Order("id")
	if userName == "" {
		fmt.Println("搜索参数为空+++++++++++")
		err := db.Find(&userList).Count(&total).Error
		if err != nil {
			// 处理错误
			return nil, 0, err
		}
		return userList, total, nil
	}
	fmt.Println("搜索参数不为空-------------",userName)
	err := db.Where("user_name LIKE ?", "%"+userName+"%").Find(&userList).Count(&total).Error
	if err != nil {
		// 处理错误
		return nil, 0, err
	}
	return userList, total, nil
}
