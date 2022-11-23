package io_loginregister

type Services interface {
	Login(username, password string) (string, error)
}
