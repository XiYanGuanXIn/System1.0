package UserService

import (
	"System1.0/Dao/UserDao"
	"System1.0/Dao/UserInfoDao"
	"System1.0/Types/ServiceType/UserType"

	"System1.0/Types/RequestAndResponseType/ErrorCodeAndInfo/ErrNo"
	"System1.0/Types/ServiceType/User"
	"fmt"
)

type UserService struct {
}

func (userService UserService) CheckPassword(username string, password string) (ErrNo.ErrNo, User.User, string) {
	user, err := UserDao.FindUserByUsername(username)
	fmt.Println(user)
	if err != nil {
		fmt.Println("Error happened when check user's password in function checkPassword.")
		return ErrNo.UnknownError, User.User{}, ""
	} else if user.Password != password {
		return ErrNo.WrongPassword, User.User{}, ""
	} else {
		info, _ := UserInfoDao.FindUserInfoByID(user.UserID)
		return ErrNo.OK, user, info
	}
}
func (userService UserService) ChangeUserInfo(username string, password string, displayName string, userInfo string) ErrNo.ErrNo {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	}
	if displayName != "" {
		user.DisplayName = displayName
	}
	if password != "" {
		user.Password = password
	}
	if UserDao.UpdateUser(user) != nil {
		return ErrNo.UnknownError
	}
	if userInfo != "" {
		if UserInfoDao.ChangeUserInfo(user.UserID, userInfo) != nil {
			return ErrNo.UnknownError
		}
	}
	return ErrNo.OK
}
func (userService UserService) RegisterUser(username string, password string) ErrNo.ErrNo {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	}
	if user.Username != "" {
		return ErrNo.UserHasExisted
	}
	user.Username = username
	user.DisplayName = username
	user.UserType = UserType.NormalUser
	user.Money = 0
	user.Password = password
	if err = UserDao.InsertUser(user); err != nil {
		return ErrNo.UnknownError
	}
	return ErrNo.OK
}
func (userService UserService) FindUserInfo(username string) (ErrNo.ErrNo, User.User, string) {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError, user, ""
	}
	if user.Username == "" {
		return ErrNo.UserNotExisted, user, ""
	}
	var info string
	info, err = UserInfoDao.FindUserInfoByID(user.UserID)
	if err != nil {
		return ErrNo.UnknownError, user, info
	}
	return ErrNo.OK, user, info
}
func (userService UserService) AddMoneyService(money float64, username string) ErrNo.ErrNo {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	}
	if user.Username == "" {
		return ErrNo.UserNotExisted
	}
	user.Money = user.Money + money
	err = UserDao.UpdateUser(user)
	if err != nil {
		return ErrNo.UnknownError
	}
	return ErrNo.OK
}
func (userService UserService) SetUserType(username string, targetUser uint, userType UserType.UserType) ErrNo.ErrNo {
	user, err := UserDao.FindUserByUsername(username)
	if err != nil {
		return ErrNo.UnknownError
	}
	if user.Username == "" {
		return ErrNo.UserNotExisted
	}
	if user.UserType < UserType.Admin {
		return ErrNo.PermDenied
	}
	var tar User.User
	tar, err = UserDao.FindUserByID(targetUser);if err != nil {return ErrNo.UnknownError}
	if tar.Username == "" {return ErrNo.UserNotExisted}
	tar.UserType = userType
	err = UserDao.UpdateUser(tar);if err != nil {return ErrNo.UnknownError}
	return ErrNo.OK
}
