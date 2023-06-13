package model

import (
	"time"

	Review "github.com/AIRBNBAPP/features/review/data"
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	UserID      uint          `json:"user_id" form:"user_id"`
	HomestayID  uint          `json:"homestay_id" form:"homestay_id"`
	Check_in    time.Time     `gorm:"type:datetime;not null" json:"check_in" form:"check_in"`
	Check_out   time.Time     `gorm:"type:datetime" json:"check_out" form:"check_out"`
	Total_night int           `json:"total_night" form:"total_night"`
	Total_price int           `json:"total_price" form:"total_price"`
	Review      Review.Review `gorm:"foreignKey:ReservationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
