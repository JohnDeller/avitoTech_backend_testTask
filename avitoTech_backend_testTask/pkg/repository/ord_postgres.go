package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type OrdPostgres struct {
	db *sqlx.DB
}

func NewOrdPostgres(db *sqlx.DB) *OrdPostgres {
	return &OrdPostgres{db: db}
}

func (r *OrdPostgres) CreateOrder(id int, userId int, cost float32) (int, error) {
	//	var id int
	query := fmt.Sprintf("INSERT INTO %s (id,user_id,balance) values ($1,$2,$3)RETURNING id", ordersTable)

	row := r.db.QueryRow(query, id, userId, cost)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	//var user_Id = getU
	return id, nil
}

func (r *OrdPostgres) DeleteOrder(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s  WHERE id = %d",
		reservesTable, id)
	_, err := r.db.Exec(query)

	return id, err
}
