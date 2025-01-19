package repository

import (
	"echo_ifElse"
	"github.com/jmoiron/sqlx"
)

type Account interface {
	Registration(req echo_ifElse.AccountRequest) (int, error)
	GetAcc(id int) (echo_ifElse.AccountResponse, error)
	AddAcc(req echo_ifElse.AccountRequest) (int, error)
	UpdateAcc(userId int, req echo_ifElse.UpdateAccountResponse) (echo_ifElse.AccountResponse, error)
	DeleteAcc(userId int) error
}

type Repository struct {
	Account
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Account: NewAccountPostgres(db),
	}
}
