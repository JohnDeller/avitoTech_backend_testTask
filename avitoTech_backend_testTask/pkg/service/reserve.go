package service

import (
	"github.com/JohnDeller/avitoTech_backend_testTask/pkg/repository"
)

type ReserveService struct {
	repo repository.Reserve
}

func NewReserveService(repo repository.Reserve) *ReserveService {
	return &ReserveService{repo: repo}
}

func (s *ReserveService) CreateReserve(id int, userId int, orderId int, cost float32) (int, error) {
	return s.repo.CreateReserve(id, userId, orderId, cost)
}

func (s *ReserveService) GetReserveBalance(userId int) (float32, error) {
	return s.repo.GetReserveBalance(userId)
}

func (s *ReserveService) DeleteReserve(Id int) (int, error) {
	return s.repo.DeleteReserve(Id)
}
