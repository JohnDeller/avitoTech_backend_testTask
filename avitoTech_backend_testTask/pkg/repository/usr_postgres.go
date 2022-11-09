package repository

import (
	"fmt"
	"github.com/JohnDeller/avitoTech_backend_testTask"
	"github.com/jmoiron/sqlx"
)

type UsrPostgres struct {
	db *sqlx.DB
}

func NewUsrPostgres(db *sqlx.DB) *UsrPostgres {
	return &UsrPostgres{db: db}
}

func (r *UsrPostgres) CreateUser(user avitotech.User) (int, error) {

	if user.Id == 0 {
		return 0, nil
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (id,balance) values ($1,$2)RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Id, user.Balance)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (r *UsrPostgres) UpdateUser(user avitotech.User) (int, error) {

	query := fmt.Sprintf("UPDATE %s SET balance = %f WHERE id = %d", usersTable, user.Balance, user.Id)

	_, err := r.db.Exec(query)
	return user.Id, err
}

func (r *UsrPostgres) GetUserBalance(userId int) (float32, error) {
	var balance float32

	query := fmt.Sprintf("SELECT balance FROM %s WHERE id = %d",
		usersTable, userId)
	err := r.db.Get(&balance, query)
	return balance, err

}
