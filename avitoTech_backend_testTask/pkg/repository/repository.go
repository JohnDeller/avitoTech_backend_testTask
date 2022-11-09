package repository

import (
	"github.com/JohnDeller/avitoTech_backend_testTask"
	"github.com/jmoiron/sqlx"
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

type Repository struct {
	User
	Reserve
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewUsrPostgres(db),
		Reserve: NewReservePostgres(db),
		Order:   NewOrdPostgres(db),
	}
}
