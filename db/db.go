package db

import (
	"github.com/kteru/scaffold-echo/db/model"
)

// Handler ...
type Handler interface {
	AuthUser(username, password string) (*model.User, error)
	FindUserByID(id uint) (*model.User, error)
}
