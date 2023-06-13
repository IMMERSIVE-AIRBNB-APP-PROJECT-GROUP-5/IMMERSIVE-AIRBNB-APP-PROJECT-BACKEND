package data

import (
	Homestay "github.com/AIRBNBAPP/features/homestay/data"
	Reservation "github.com/AIRBNBAPP/features/reservation/data"
	"gorm.io/gorm"
)

// user
type User struct {
	gorm.Model
	User_name   string                    `gorm:"type:varchar(255)" json:"user_name" form:"user_name"`
	Email       string                    `gorm:"type:varchar(100);unique;not null" json:"email" form:"email"`
	Password    string                    `gorm:"type:longtext;not null" json:"password" form:"password"`
	Phone       string                    `gorm:"type:varchar(100);unique;not null" json:"phone" form:"phone"`
	Role        string                    `gorm:"type:enum('reservant','hosters');not null" json:"role" form:"role"`
	Url         string                    `gorm:"type:longtext" json:"url" form:"url"`
	Homestay    []Homestay.Homestay       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Reservation []Reservation.Reservation `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
