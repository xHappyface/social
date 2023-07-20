// connect to core and not to adapater

package application

import "social/internal/core/user"

// define a service interface that requires the same functionality as the repo interface

type UserService interface {
	Create(*user.User) error
	ReadByID(string) (*user.User, error)
	Update(*user.User) error
	DeleteByID(string) error
}

// inject the repository as a depency to the service struct
// the service struct will act as the controller so we never directly interact with the repository

type UserServiceImpl struct {
	userRepo user.UserRepository
}

// fulfilling the interface by giving the service struct the required functionality
// will allow us to return the struct as the service interface type

func NewUserService(userRepo user.UserRepository) UserService {
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

func (serv *UserServiceImpl) Update(user *user.User) error {
	return serv.userRepo.Update(user)
}

func (serv *UserServiceImpl) DeleteByID(userID string) error {
	return serv.userRepo.DeleteByID(userID)
}
