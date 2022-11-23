package io_loginregister

import (
	"STACK-GO/adapter"

	domain "STACK-GO/modules/io_loginregister/domain"
	repo "STACK-GO/modules/io_loginregister/repo"
	jwtToken "STACK-GO/utils/token"
)

type Services struct {
	adapter adapter.RepoAdapter
}

func NewService(port adapter.RepoAdapter) *Services {
	return &Services{
		adapter: port,
	}
}

func (s *Services) Login(username, password string) (string, error) {
	data, err := s.adapter.Auth(username, password, repo.QryFindByUser(), &domain.Login{})
	if err != nil {
		return "", err
	}

	dataLogin := data.(*domain.Login)
	token, err := jwtToken.GenerateToken(dataLogin.UID)

	if err != nil {
		return "", err
	}

	return token, nil
}
