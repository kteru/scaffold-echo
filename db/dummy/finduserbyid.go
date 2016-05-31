package dummy

import (
	"fmt"

	"github.com/kteru/scaffold-echo/db/model"
)

// FindUserByID ...
func (hdl *Handler) FindUserByID(id uint) (*model.User, error) {
	for _, user := range *hdl.users {
		if user.ID == id {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("id not found")
}
