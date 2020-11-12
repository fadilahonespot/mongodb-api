package usecase

import (
	"mongodb-api/model"
	"mongodb-api/user")


type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

func CreateUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo}
}

func (e *UserUsecaseImpl) InsertUser(user *model.User) (*model.User, error) {
	return e.userRepo.InsertUser(user)
}

func (e *UserUsecaseImpl) FindAllUser() (*[]model.User, error) {
	return e.userRepo.FindAllUser()
}

func (e *UserUsecaseImpl) FindUser(id string) (*model.User, error) {
	return e.userRepo.FindUser(id)
}

func (e *UserUsecaseImpl) UpdateUser(id string, user *model.User) (*model.User, error) {
	return e.userRepo.UpdateUser(id, user)
}

func (e *UserUsecaseImpl)  DeleteUser(id string) error  {
	return e.userRepo.DeleteUser(id)
}