package service

import (
	avitotech "github.com/JohnDeller/avitoTech_backend_testTask"
	"github.com/JohnDeller/avitoTech_backend_testTask/pkg/repository"
)

type UsrService struct {
	repo repository.User
}

func NewUsrService(repo repository.User) *UsrService {
	return &UsrService{repo: repo}
}

func (s *UsrService) CreateUser(user avitotech.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *UsrService) UpdateUser(user avitotech.User) (int, error) {
	return s.repo.UpdateUser(user)
}

func (s *UsrService) GetUserBalance(userId int) (float32, error) {
	return s.repo.GetUserBalance(userId)
}
