package data

import (
	"errors"

	"github.com/AIRBNBAPP/features/user"
	"github.com/AIRBNBAPP/helper"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

// CreateUser implements user.UserDataInterface.
func (repo *userQuery) CreateUser(userInput user.Core) error {
	hashedPassword, errHash := helper.HashPassword(userInput.Password)
	if errHash != nil {
		return errors.New("error hash password")
	}
	userInput.Password = hashedPassword
	// mapping dari struct entities core ke gorm model
	userInputGorm := CoreToModel(userInput)
	tx := repo.db.Create(&userInputGorm) //Ini query insert ke database
	if tx.Error != nil {
		return errors.New("Email atau phone telah terdaftar")
	}
	if tx.RowsAffected == 0 {
		return errors.New("Insert Failed, row affected = 0")
	}
	return nil
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}
