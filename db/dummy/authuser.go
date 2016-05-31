package dummy

import (
	"fmt"

	"github.com/kteru/scaffold-echo/db/model"
)

// AuthUser ...
func (hdl *Handler) AuthUser(username, password string) (*model.User, error) {
	for _, user := range *hdl.users {
		if user.Username == username && user.Password == password {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("username or password is incorrect")
}
