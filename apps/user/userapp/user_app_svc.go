package userapp

import (
	"errors"
	"log"
	"startup-backend/apps/user/userentity"
	"startup-backend/apps/user/userinfra/userrepo"
	"startup-backend/pkg/shared"
)

type UserAppSvc struct {
	userRepo *userrepo.UserRepository
}

func NewUserAppSvc(userrepo *userrepo.UserRepository) *UserAppSvc {
	return &UserAppSvc{
		userRepo: userrepo,
	}
}


func (service *UserAppSvc) AddNewUser(fullname string) error {

	err := service.userRepo.Save(&userentity.UserEntity{FullName: fullname})

	if err != nil {
		if errors.Is(err, &shared.SystemError{}) {
			return err
		}
		log.Println(err)
		return err
	}

	return nil

}
