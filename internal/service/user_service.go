package service

import (
	"github.com/klass-lk/test/internal/model"

	"github.com/klass-lk/ginboot/db/inmemory"
)

type UserService interface {
	GetUser(id string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
}

type userService struct {
	userRepo *inmemory.InMemoryRepository[model.User]
}

func NewUserService(userRepo *inmemory.InMemoryRepository[model.User]) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetUser(id string) (model.User, error) {
	return s.userRepo.FindById(id)
}

func (s *userService) CreateUser(user model.User) (model.User, error) {
	err := s.userRepo.Save(user)
	return user, err
}
