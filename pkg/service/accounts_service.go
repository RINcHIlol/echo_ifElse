package service

import (
	"echo_ifElse"
	"echo_ifElse/pkg/repository"
	"fmt"
)

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

func (a *AccountService) Registration(req echo_ifElse.AccountRequest) (int, error) {
	fmt.Println("service")
	return a.repo.Registration(req)
}

func (a *AccountService) GetAcc(id int) (echo_ifElse.AccountResponse, error) {
	return a.repo.GetAcc(id)
}
