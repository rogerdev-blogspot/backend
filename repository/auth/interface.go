package auth

import "templateLoginRegister/entities"

type Auth interface {
	Login(email, password string) (entities.User, error)
	LoginGoogle(email string) (entities.User, error)
}
