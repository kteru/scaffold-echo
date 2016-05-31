package dummy

import (
	"github.com/kteru/scaffold-echo/db/model"
)

// Handler ...
type Handler struct {
	users *[]model.User
}

// NewHandler ...
func NewHandler() (*Handler, error) {
	hdl := &Handler{
		users: &[]model.User{
			{
				ID:       1,
				Username: "taro",
				Password: "1234",
			},
			{
				ID:       2,
				Username: "hanako",
				Password: "1234",
			},
		},
	}

	return hdl, nil
}
