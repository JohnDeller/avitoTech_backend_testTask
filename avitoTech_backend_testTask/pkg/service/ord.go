package service

import "github.com/JohnDeller/avitoTech_backend_testTask/pkg/repository"

type OrdService struct {
	repo repository.Order
}

func NewOrdService(repo repository.Order) *OrdService {
	return &OrdService{repo: repo}
}

func (s *OrdService) CreateOrder(id int, userId int, cost float32) (int, error) {
	return s.repo.CreateOrder(id, userId, cost)
}

func (s *OrdService) DeleteOrder(id int) (int, error) {
	return s.repo.DeleteOrder(id)
}
