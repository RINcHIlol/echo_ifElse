package repository

import (
	"echo_ifElse"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AccountPostgres struct {
	db *sqlx.DB
}

func NewAccountPostgres(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{db: db}
}

func (a *AccountPostgres) Registration(req echo_ifElse.AccountRequest) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (firstName, lastName, email, password, role) values ($1, $2, $3, $4, $5) RETURNING id", usersTable)
	row := a.db.QueryRow(query, req.FirstName, req.LastName, req.Email, req.Password, "USER")
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AccountPostgres) GetAcc(id int) (echo_ifElse.AccountResponse, error) {
	var response echo_ifElse.AccountResponse
	query := fmt.Sprintf("Select id, firstname, lastname, email, role from %s where id=$1", usersTable)
	row := a.db.QueryRow(query, id)
	if err := row.Scan(&response.Id, &response.FirstName, &response.LastName, &response.Email, &response.Role); err != nil {
		return response, err
	}
	return response, nil
}
