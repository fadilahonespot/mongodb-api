package user

import "mongodb-api/model"

type UserRepo interface {
	InsertUser(user *model.User) (*model.User, error)
	FindAllUser() (*[]model.User, error)
	FindUser(id string) (*model.User, error)
	UpdateUser(id string, user *model.User) (*model.User, error)
	DeleteUser(id string) error 
}