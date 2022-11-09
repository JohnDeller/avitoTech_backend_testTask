package service

import (
	"github.com/JohnDeller/avitoTech_backend_testTask"
	"github.com/JohnDeller/avitoTech_backend_testTask/pkg/repository"
)

type User interface {
	CreateUser(user avitotech.User) (int, error)
	UpdateUser(user avitotech.User) (int, error)
	GetUserBalance(userId int) (float32, error)
}

type Reserve interface {
	CreateReserve(id int, userId int, orderId int, cost float32) (int, error)
	DeleteReserve(Id int) (int, error)
	GetReserveBalance(userId int) (float32, error)
}

type Order interface {
	CreateOrder(id int, userId int, cost float32) (int, error)
	DeleteOrder(id int) (int, error)
}

type Service struct {
	User
	Reserve
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:    NewUsrService(repos.User),
		Reserve: NewReserveService(repos.Reserve),
		Order:   NewOrdService(repos.Order),
	}
}
