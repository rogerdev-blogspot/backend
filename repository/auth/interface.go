package auth

import "backend/entities"

type Auth interface {
	Login(email, password string) (entities.User, error)
	LoginGoogle(email string) (entities.User, error)
}
