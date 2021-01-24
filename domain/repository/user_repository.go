package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"user/domain/model"
)

type IUserRepository interface {
	I() error

	//根据用户名称查找用户信息
	FindUserByName(string) (*model.User,error)
	//根据用户id查找
	FindUserByID(int64) (*model.User,error)
	//创建用户
	CreateUser(*model.User) (int64,error)//返回id
	//根据用户ID删除信息
	DeleteUserByID(int64) error
	//更新用户信息
	UpdateUser(*model.User) error
	//查找所有用户
	FindAll()([]model.User,error)
}
func NewUserRepository(db *gorm.DB) IUserRepository{
	return &UserRepository{db}
}
type UserRepository struct{
	mysqlDb *gorm.DB
}
//初始化表
func(u *UserRepository) I() error{
	if u.mysqlDb.HasTable(&model.User{})==true{
		return nil
	}
	return u.mysqlDb.CreateTable(&model.User{}).Error
}
//根据名称查找用户
func(u *UserRepository)FindUserByName(name string)(user *model.User,err error){
	user = &model.User{}
	return user,u.mysqlDb.Where("user_name = ?",name).Find(user).Error
}

func(u *UserRepository)FindUserByID(id int64)(user *model.User,err error){
	user = &model.User{}
	return user,u.mysqlDb.First(user,id).Error
}
//创建用户
func (u *UserRepository)CreateUser(user *model.User) (int64,error){
	return user.ID,u.mysqlDb.Model(user).Create(&user).Error
}
//根据ID删除用户
func (u *UserRepository)DeleteUserByID(UserID int64) error{
	return u.mysqlDb.Where("id = ?",UserID).Delete(&model.User{}).Error
}
func (u *UserRepository)UpdateUser(user *model.User) error{
	return u.mysqlDb.Model(user).Update(&user).Error
}
//查找所有用户
func (u *UserRepository)FindAll()(userAll []model.User,err error){
	return userAll,u.mysqlDb.Find(&userAll).Error
}