package repository

import (
	"echo_ifElse"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
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

func (a *AccountPostgres) AddAcc(req echo_ifElse.AccountRequest) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (firstName, lastName, email, password, role) values ($1, $2, $3, $4, $5) RETURNING id", usersTable)
	row := a.db.QueryRow(query, req.FirstName, req.LastName, req.Email, req.Password, req.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AccountPostgres) UpdateAcc(userId int, req echo_ifElse.UpdateAccountResponse) (echo_ifElse.AccountResponse, error) {
	var jsonResponse echo_ifElse.AccountResponse

	// Формируем список изменений
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if req.FirstName != nil {
		setValues = append(setValues, fmt.Sprintf("firstname=$%d", argId))
		args = append(args, *req.FirstName)
		argId++
	}
	if req.LastName != nil {
		setValues = append(setValues, fmt.Sprintf("lastname=$%d", argId))
		args = append(args, *req.LastName)
		argId++
	}
	if req.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *req.Email)
		argId++
	}
	if req.Password != nil {
		setValues = append(setValues, fmt.Sprintf("password=$%d", argId))
		args = append(args, *req.Password)
		argId++
	}
	if req.Role != nil {
		setValues = append(setValues, fmt.Sprintf("role=$%d", argId))
		args = append(args, *req.Role)
		argId++
	}

	// Если ничего не обновляется, возвращаем ошибку
	if len(setValues) == 0 {
		return jsonResponse, fmt.Errorf("no fields to update")
	}

	// Генерируем SQL-запрос
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d RETURNING id, firstname, lastname, email, role", usersTable, setQuery, argId)
	args = append(args, userId)

	// Выполняем запрос
	err := a.db.QueryRow(query, args...).Scan(
		&jsonResponse.Id,
		&jsonResponse.FirstName,
		&jsonResponse.LastName,
		&jsonResponse.Email,
		&jsonResponse.Role,
	)
	if err != nil {
		return jsonResponse, fmt.Errorf("failed to update account: %w", err)
	}

	return jsonResponse, nil
}

func (a *AccountPostgres) DeleteAcc(userId int) error {
	query := fmt.Sprintf("delete from %s where id=$1", usersTable)
	_, err := a.db.Exec(query, userId)
	return err
}
