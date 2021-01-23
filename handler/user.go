package handler

import (
	"context"
	"user/domain/model"
	"user/domain/service"
	users "user/proto"
)

//对外暴露的服务都在这里编写。
type User struct{
	UserDataService service.IUserDataService
}
//注册
func(u *User)Register(ctx context.Context,request *users.UserRegisterRequest ,response *users.UserRegisterResponse)error{
	userRegister := &model.User{
		UserName: request.UserName,
		FirstName: request.FirstName,
		HashPassword: request.Pwd,
	}
	_,err := u.UserDataService.AddUser(userRegister)
	if err!=nil{
		return err
	}
	response.Message = "添加成功"
	return nil
}
//登陆
func(u *User)Login(ctx context.Context,request *users.UserLoginRequest,response *users.UserLoginResponse) error{
	isOk,err := u.UserDataService.CheckPwd(request.UserName,request.Pwd)
	if err != nil{
		return err
	}
	response.IsSuccess = isOk
	return nil
}
func(u *User)GetUserInfo(ctx context.Context,request *users.UserInfoRequest,response *users.UserInfoResponse) error{
	userInfo,err := u.UserDataService.FindUserByID(request.UserId)
	if err != nil{
		return err
	}
	response = GetUserForResponse(userInfo)
	return nil
}

//类型转化
func GetUserForResponse(userModel *model.User) *users.UserInfoResponse{
	response := &users.UserInfoResponse{}
	response.UserName = userModel.UserName
	response.FirstName = userModel.FirstName
	response.UserId = userModel.ID
	return response
}