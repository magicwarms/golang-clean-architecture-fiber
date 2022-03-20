package userinterface

import "startup-backend/apps/user/userentity"

type UserRepositoryInterface interface {
	Save(user *userentity.UserEntity) error
}
