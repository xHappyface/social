// connect to core and not to adapater

package database

import (
	"github.com/xHappyface/social/core/user"
	"github.com/xHappyface/social/internal/ports"
)

// define a service interface that requires the same functionality as the repo interface

type UserService interface {
	Create(*user.User) error
	ReadByID(string) (*user.User, error)
	ReadByName(string) (*user.User, error)
	Update(*user.User) error
	DeleteByID(string) error
}

// inject the repository as a depency to the service struct
// the service struct will act as the controller so we never directly interact with the repository

type UserServiceImpl struct {
	userRepo ports.UserRepository
}

// fulfilling the interface by giving the service struct the required functionality
// will allow us to return the struct as the service interface type

func NewUserService(userRepo ports.UserRepository) UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

// have the service struct fulfill the service interface requirements

func (serv *UserServiceImpl) Create(user *user.User) error {
	return serv.userRepo.Create(user)
}

func (serv *UserServiceImpl) ReadByID(userID string) (*user.User, error) {
	return serv.userRepo.ReadByID(userID)
}

func (serv *UserServiceImpl) ReadByName(userName string) (*user.User, error) {
	return serv.userRepo.ReadByName(userName)
}

func (serv *UserServiceImpl) Update(user *user.User) error {
	return serv.userRepo.Update(user)
}

func (serv *UserServiceImpl) DeleteByID(userID string) error {
	return serv.userRepo.DeleteByID(userID)
}
