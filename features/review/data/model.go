package model

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ReservationID uint   `json:"reservation_id" form:"reservation_id"`
	HomestayID    uint   `json:"homestay_id" form:"homestay_id"`
	Rating        int    `gorm:"type:enum('1','2','3','4','5');not null" json:"rating" form:"rating"`
	Comment       string `gorm:"type:longtext" json:"comment" form:"comment"`
}
