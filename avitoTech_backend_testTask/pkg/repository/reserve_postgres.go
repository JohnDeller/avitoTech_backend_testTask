package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ReservePostgres struct {
	db *sqlx.DB
}

func NewReservePostgres(db *sqlx.DB) *ReservePostgres {
	return &ReservePostgres{db: db}
}
func (r *ReservePostgres) CreateReserve(id int, userId int, orderId int, cost float32) (int, error) {

	query := fmt.Sprintf("INSERT INTO %s (id,user_id,order_id,balance) values ($1,$2,$3,$4)RETURNING id", reservesTable)

	row := r.db.QueryRow(query, id, userId, orderId, cost)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func (r *ReservePostgres) DeleteReserve(Id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s  WHERE id = %d",
		reservesTable, Id)
	_, err := r.db.Exec(query)

	return Id, err
}

func (r *ReservePostgres) GetReserveBalance(userId int) (float32, error) {
	var balance float32

	query := fmt.Sprintf("SELECT balance FROM %s WHERE user_id = %d",
		reservesTable, userId)
	err := r.db.Get(&balance, query)
	return balance, err
}
