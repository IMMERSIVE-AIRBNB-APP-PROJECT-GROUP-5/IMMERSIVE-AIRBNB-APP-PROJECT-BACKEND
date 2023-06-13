package data

import (
	"errors"

	"github.com/AIRBNBAPP/app/middlewares"
	"github.com/AIRBNBAPP/features/user"
	"github.com/AIRBNBAPP/helper"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

// GetUserById implements user.UserDataInterface.
func (repo *userQuery) GetUserById(id int) ([]user.Core, error) {
	var userData []User
	tx := repo.db.First(&userData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping dari struct gorm model ke struct entities core
	var usersCoreAll []user.Core
	for _, value := range userData {
		var userCore = user.Core{
			Id:         value.ID,
			User_name:  value.User_name,
			Email:      value.Email,
			Phone:      value.Phone,
			Password:   value.Password,
			Status:     user.UserStatus(value.Status),
			Created_at: value.CreatedAt,
			Updated_at: value.UpdatedAt,
		}
		usersCoreAll = append(usersCoreAll, userCore)
	}
	return usersCoreAll, nil
}

// Login implements user.UserDataInterface.
func (repo *userQuery) Login(email string, password string) (user.Core, string, error) {
	var userData User

	// Mencocokkan data inputan email dengan email di database
	tx := repo.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		return user.Core{}, "", tx.Error
	}
	if tx.RowsAffected == 0 {
		return user.Core{}, "", errors.New("Login gagal, email anda salah")
	}
	// Mencocokkan data inputan password dengan password yang telah di hashing di database
	checkPassword := helper.CheckPasswordHash(userData.Password, password)
	if !checkPassword {
		return user.Core{}, "", errors.New("Login gagal, password anda salah")
	}

	token, errToken := middlewares.CreateToken(int(userData.ID))
	if errToken != nil {
		return user.Core{}, "", errToken
	}

	dataCore := user.Core{
		Id:         userData.ID,
		User_name:  userData.User_name,
		Email:      userData.Email,
		Password:   userData.Password,
		Phone:      userData.Phone,
		Status:     user.UserStatus(userData.Status),
		Created_at: userData.CreatedAt,
		Updated_at: userData.UpdatedAt,
	}

	return dataCore, token, nil
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
