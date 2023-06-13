package data

import (
	Homestay "github.com/AIRBNBAPP/features/homestay/data"
	Reservation "github.com/AIRBNBAPP/features/reservation/data"
	"github.com/AIRBNBAPP/features/user"
	"gorm.io/gorm"
)

type UserStatus string

const (
	Reservant UserStatus = "reservant"
	Hosters   UserStatus = "hosters"
)

// user
type User struct {
	gorm.Model
	User_name   string                    `gorm:"type:varchar(255)" json:"user_name" form:"user_name"`
	Email       string                    `gorm:"type:varchar(100);unique;not null" json:"email" form:"email"`
	Password    string                    `gorm:"type:longtext;not null" json:"password" form:"password"`
	Phone       string                    `gorm:"type:varchar(100);unique;not null" json:"phone" form:"phone"`
	Status      UserStatus                `gorm:"type:enum('reservant','hosters');not null" json:"status" form:"status"`
	Url         string                    `gorm:"type:longtext" json:"url" form:"url"`
	Homestay    []Homestay.Homestay       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Reservation []Reservation.Reservation `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// mapping dari core ke gorm
func CoreToModel(dataCore user.Core) User {
	return User{
		Model:     gorm.Model{},
		User_name: dataCore.User_name,
		Email:     dataCore.Email,
		Password:  dataCore.Password,
		Phone:     dataCore.Phone,
		Status:    UserStatus(Reservant),
	}
}
