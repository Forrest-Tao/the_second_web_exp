package service

import (
	"experimen_2/modle"
	"experimen_2/pkg/utiles"
	"experimen_2/serializer"
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

func (UserService *UserService) Register() serializer.Response {
	var user modle.User
	//修复bug 遗忘名字的赋值
	user.UserName = UserService.UserName
	var count int
	modle.DB.Model(&modle.User{}).Where("user_name=?", UserService.UserName).
		First(&user).Count(&count)
	if count >= 1 {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "该用户已存在,无需再注册",
			Error:  "",
		}
	} else { //数据库中没有用户的信息，则为他注册账号，密码加密
		if err := user.SetPassWord(UserService.Password); err != nil {
			return serializer.Response{
				Status: 400,
				Data:   nil,
				Msg:    err.Error(),
				Error:  err.Error(),
			}
		}
	}

	//创建用户
	if err := modle.DB.Model(&modle.User{}).Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Data:   nil,
			Msg:    "数据库操作错误",
			Error:  "",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   user,
		Msg:    "用户注册成功",
		Error:  "",
	}
}

func (UserService *UserService) Login() serializer.Response {
	var user modle.User
	//var count int
	//查询用户是否存在
	//报错 找不到或者是查询没有反应 user_name 应该改为user_name=?
	if err := modle.DB.Where("user_name=?", UserService.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println(UserService.UserName, err)
			return serializer.Response{
				Status: 400,
				Data:   nil,
				Msg:    "用户不存在，请先注册",
				Error:  "",
			}
		}
	}
	//检验密码是否正确
	if user.CheckPassWord(UserService.Password) == false {
		fmt.Println("用户密码： ", UserService.Password)
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "密码错误",
			Error:  "",
		}
	}
	//当用户存在且密码正确时 发一个token
	token, err := utiles.GenerateToken(user.ID, user.UserName, user.PassWordDigest)
	//如果签证错误
	if err != nil {
		return serializer.Response{
			Status: 500,
			Data:   nil,
			Msg:    "Token签发错误",
			Error:  "",
		}
	}
	//如果用户存在 密码正确 签证成功 则返回200
	return serializer.Response{
		Status: 200,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
		Msg:   "登录成功",
		Error: "",
	}
}
