package service

import (
	"github.com/with-insomnia/Forum-Golang/internal/model"
	"github.com/with-insomnia/Forum-Golang/internal/repository"
)

type UserService interface {
	GetUserByToken(token string) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
}

type userService struct {
	repository repository.UserQuery
}

func NewUserService(dao repository.DAO) UserService {
	return &userService{
		dao.NewUserQuery(),
	}
}

func (u *userService) GetUserByToken(token string) (model.User, error) {
	userID, err := u.repository.GetUserIdByToken(token)
	if err != nil {
		return model.User{}, err
	}
	user, err := u.repository.GetUserByUserId(userID + 1)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userService) GetUserByEmail(email string) (model.User, error) {
	return model.User{}, nil
}
