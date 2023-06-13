package data

import (
	Reservation "github.com/AIRBNBAPP/features/reservation/data"
	Review "github.com/AIRBNBAPP/features/review/data"
	"gorm.io/gorm"
)

type Homestay struct {
	gorm.Model
	Homestay_name string                    `gorm:"type:varchar(255);not null" json:"homestay_name" form:"homestay_name"`
	UserID        uint                      `json:"user_id" form:"user_id"`
	City_name     string                    `json:"city_name" form:"city_name"`
	Address       string                    `gorm:"type:longtext;not null" json:"address" form:"address"`
	Total_guest   int                       `gorm:"not null" json:"total_guest" form:"total_guest"`
	Price         int                       `gorm:"not null" json:"price" form:"price"`
	Description   string                    `gorm:"type:longtext;not null" json:"description" form:"description"`
	Reservation   []Reservation.Reservation `gorm:"foreignKey:HomestayID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Review        []Review.Review           `gorm:"foreignKey:HomestayID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Facility      []*Facility               `gorm:"many2many:detail;" json:"facility" form:"facility"`
	Picture       Picture                   `gorm:"foreignKey:HomestayID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Picture struct {
	gorm.Model
	HomestayID uint   `json:"homestay_id" form:"homestay_id"`
	Url        string `gorm:"type:longtext" json:"url" form:"url"`
}

type Facility struct {
	gorm.Model
	Facility_name string      `gorm:"type:varchar(255);not null" json:"facility_name" form:"facility_name"`
	Homestay      []*Homestay `gorm:"many2many:detail;" json:"homestay" form:"homestay"`
}
