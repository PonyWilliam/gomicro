package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"user/domain/model"
	"user/domain/repository"
)

type IUserDataService interface {
	//标注需要完成的方法
	AddUser(*model.User) (int64,error)
	DeleteUser(int64) error
	UpdateUser(user *model.User,isChangePwd bool)(err error)
	FindUserByName(string)(*model.User,error)
	FindUserByID(int64)(*model.User,error)
	CheckPwd(username string,pwd string)(isOk bool,err error)
}

//创建实例
func NewUserDataService(userRepository repository.IUserRepository)IUserDataService{
	return &UserDataService{UserRepository: userRepository}
}
type UserDataService struct{
	UserRepository repository.IUserRepository
}
//实现所有接口

//加密用户密码
func GeneratePassword(userPassword string)([]byte,error){
	return bcrypt.GenerateFromPassword([]byte(userPassword),bcrypt.DefaultCost)
}
func ValidatePassword(userPassword string,hashed string)(isOk bool,err error){
	if err = bcrypt.CompareHashAndPassword([]byte(hashed),[]byte(userPassword));err!=nil{
		return false,errors.New("密码比对错误")
	}
	return true,nil
}

//插入用户
func (u *UserDataService)AddUser(user *model.User)(UserID int64,err error){
	pwdByte,err := GeneratePassword(user.HashPassword)
	user.HashPassword = string(pwdByte)
	return u.UserRepository.CreateUser(user)
	//创建了user用户，并对密码进行了加密
}

//删除用户
func(u *UserDataService)DeleteUser(UserID int64) error{
	return u.UserRepository.DeleteUserByID(UserID)
}
func(u *UserDataService)UpdateUser(user *model.User,isChangePwd bool) error{
	if isChangePwd{
		pwdByte,err := GeneratePassword(user.HashPassword)
		if err!=nil{
			return err
		}
		user.HashPassword = string(pwdByte)
		//在插入后可以做一个rabbitmq
	}
	return u.UserRepository.UpdateUser(user)
	//用模型做实际更新
}
func(u *UserDataService)FindUserByName(name string)(*model.User,error){
	return u.UserRepository.FindUserByName(name)
}
func(u *UserDataService)FindUserByID(id int64)(*model.User,error){
	return u.UserRepository.FindUserByID(id)
}
func (u *UserDataService)CheckPwd(userName string,pwd string)(isOk bool,err error){
	user,err := u.UserRepository.FindUserByName(userName)
	if err!=nil{
		return false,errors.New("没有此用户")
	}
	return ValidatePassword(pwd,user.HashPassword)//判断是否一样
}