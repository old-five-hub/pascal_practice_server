package account_service

import "pascal_practice_server/models"

type Account struct {
	Username string
	Password string
}

func (a *Account) Login() (bool, error) {
	return models.Login(a.Username, a.Password)
}
