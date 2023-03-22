package services

import (
	"errors"

	"github.com/Zaanter/be-goserver/pkg/entities"
	"github.com/Zaanter/be-goserver/pkg/repository"
)

type UserService interface {
	Validate(user *entities.IUser) error
	Create(user *entities.IUser) (*entities.IUser, error)
	Delete(uid string) error
}

type service struct{}

func NewUserService() UserService {
	return &service{}
}

var (
	repo repository.UserRepository = repository.NewUserRepository()
)

func (*service) Validate(user *entities.IUser) error {
	if user == nil {
		err := errors.New("user is empty")
		return err
	}

	if user.Lastname == "" {
		err := errors.New("user lastname is empty")
		return err
	}

	if user.Firstname == "" {
		err := errors.New("user firstname is empty")
		return err
	}

	if user.Email == "" {
		err := errors.New("user email is empty")
		return err
	}

	if user.Dob == "" {
		err := errors.New("user dob is empty")
		return err
	}

	return nil
}

func (*service) Create(user *entities.IUser) (*entities.IUser, error) {
	return repo.Create(user)
}

func (*service) Delete(uid string) error {
	return repo.Delete(uid)
}
