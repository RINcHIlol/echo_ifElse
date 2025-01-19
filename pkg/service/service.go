package service

import (
	"echo_ifElse"
	"echo_ifElse/pkg/repository"
)

type Account interface {
	Registration(req echo_ifElse.AccountRequest) (int, error)
	GetAcc(id int) (echo_ifElse.AccountResponse, error)
	AddAcc(req echo_ifElse.AccountRequest) (int, error)
	UpdateAcc(userId int, req echo_ifElse.UpdateAccountResponse) (echo_ifElse.AccountResponse, error)
	DeleteAcc(userId int) error
}

type Service struct {
	Account
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		NewAccountService(repos.Account),
	}
}
