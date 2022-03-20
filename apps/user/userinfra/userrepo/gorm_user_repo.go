package userrepo

import (
	"startup-backend/apps/user/userentity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

//NewRepo is the single instance repo that is being created.
// inject dari config
func NewUserRepo(gormDB *gorm.DB) *UserRepository {
	// migrate disini
	gormDB.AutoMigrate(&gormUserModel{})

	return &UserRepository{
		db: gormDB,
	}
}

// AddUser is to get all books data
func (repo *UserRepository) Save(user *userentity.UserEntity) error {
	userRepo := newGormUserModel(user)

	if err := repo.db.Save(userRepo).Error; err != nil {
		return err
	}

	return nil

}

// 10

// admin edit user  version += 1
// user a 10
// user a 10
