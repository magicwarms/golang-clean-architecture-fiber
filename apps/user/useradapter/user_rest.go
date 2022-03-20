package useradapter

func initUserRestRouter(service *UserService) {
	customerGroup := service.App.Group("/costumer")

	v1 := customerGroup.Group("/v1.0.0/user")

	v1.Post("/", addNewUser(service.UserAppSvc))

}
