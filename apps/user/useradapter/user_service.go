package useradapter

import (
	"startup-backend/apps/user/userapp"
	"startup-backend/apps/user/userinfra/userrepo"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserService struct {
	// app service
	App        *fiber.App
	UserAppSvc *userapp.UserAppSvc
	// query
	// UserAppQuerySvc

}

func NewUserService(gormDB *gorm.DB, app *fiber.App) *UserService {
	// userRepo := userrepo.NewUserRepo(config.GormDB)
	userRepo := userrepo.NewUserRepo(gormDB)

	return &UserService{
		App:        app,
		UserAppSvc: userapp.NewUserAppSvc(userRepo),
	}
}

func (service *UserService) StartUserService() *UserService {
	initUserRestRouter(service)
	return service
}
