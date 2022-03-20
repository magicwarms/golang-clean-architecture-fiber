package userrepo

import (
	"startup-backend/apps/user/userentity"
)

type gormUserModel struct {
	// gorm.Model
	ID       string `gorm:"default:uuid_generate_v4();primaryKey"`
	Fullname string `gorm:"index"`
	// Version
}

func newGormUserModel(user *userentity.UserEntity) *gormUserModel {
	return &gormUserModel{
		ID:       string(user.ID),
		Fullname: user.FullName,
	}
}

func (gormUserModel) TableName() string {
	return "user"
}

// entity
// type UserID string

// func (userID UserID) Validate() error {

// 	if userID == "" {
// 		return shared.NewSystemError("")
// 	}

// 	return nil

// 	// var err error

// 	// if userID == "" {
// 	// 	err = shared.NewSystemError("")
// 	// 	// return err
// 	// }

// 	// if errors.Is(err, &shared.SystemError{}) {
// 	// 	return err
// 	// }

// 	// return err
// }
