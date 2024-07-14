package service

import (
	"awesomeProject/internal/entity"
	"awesomeProject/internal/repository"
)

type Service interface {
	CreateUser(entity.User) error
	GetUser(id string) (entity.User, error)
	UpdateUser(entity.User) (entity.User, error)
	DeleteUser(id string) error
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateUser(user entity.User) error {
	return s.repo.CreateUser(user)
}
func (s *service) GetUser(id string) (entity.User, error) {
	return s.repo.GetUser(id)
}
func (s *service) UpdateUser(user entity.User) (entity.User, error) {
	return s.repo.UpdateUser(user)
}
func (s *service) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}
