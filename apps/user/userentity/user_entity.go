package userentity

import "errors"

type UserID string

type UserEntity struct {
	ID       UserID
	FullName string
}

func (user *UserEntity) Validate() error {

	if user.ID == "" {
		return errors.New("id kosong bangsat")
	}

	return nil
}
