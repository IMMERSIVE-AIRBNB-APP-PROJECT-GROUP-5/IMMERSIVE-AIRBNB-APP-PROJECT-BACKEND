package data

import (
	"time"

	"github.com/AIRBNBAPP/features/homestay"
	"gorm.io/gorm"
)

type Homestay struct {
	gorm.Model
	Homestay_name string        `gorm:"type:varchar(255);not null" json:"homestay_name" form:"homestay_name"`
	UserID        uint          `json:"user_id" form:"user_id"`
	City_name     string        `json:"city_name" form:"city_name"`
	Address       string        `gorm:"type:longtext;not null" json:"address" form:"address"`
	Total_guest   int           `gorm:"not null" json:"total_guest" form:"total_guest"`
	Price         int           `gorm:"not null" json:"price" form:"price"`
	Description   string        `gorm:"type:longtext;not null" json:"description" form:"description"`
	Reservation   []Reservation `gorm:"foreignKey:HomestayID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Review        []Review      `gorm:"foreignKey:HomestayID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Facilities    []Facility    `gorm:"many2many:homestay_facilities;"`
	Picture       Picture       `gorm:"foreignKey:HomestayID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Facility struct {
	gorm.Model
	Facility_name string `gorm:"type:varchar(255);not null" json:"facility_name" form:"facility_name"`
}

type HomestayFacility struct {
	gorm.Model
	HomestayID uint `gorm:"foreignKey:HomestayRefer"`
	FacilityID uint `gorm:"foreignKey:FacilityRefer"`
}

type Reservation struct {
	gorm.Model
	UserID      uint      `json:"user_id" form:"user_id"`
	HomestayID  uint      `json:"homestay_id" form:"homestay_id"`
	Check_in    time.Time `gorm:"type:datetime;not null" json:"check_in" form:"check_in"`
	Check_out   time.Time `gorm:"type:datetime" json:"check_out" form:"check_out"`
	Total_night int       `json:"total_night" form:"total_night"`
	Total_price int       `json:"total_price" form:"total_price"`
	Review      Review    `gorm:"foreignKey:ReservationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Review struct {
	gorm.Model
	ReservationID uint   `json:"reservation_id" form:"reservation_id"`
	HomestayID    uint   `json:"homestay_id" form:"homestay_id"`
	Rating        int    `gorm:"type:enum('1','2','3','4','5');not null" json:"rating" form:"rating"`
	Comment       string `gorm:"type:longtext" json:"comment" form:"comment"`
}

type Picture struct {
	gorm.Model
	HomestayID uint   `json:"homestay_id" form:"homestay_id"`
	Url        string `gorm:"type:longtext" json:"url" form:"url"`
}

// mapping dari core ke gorm
func CoreToModel(dataCore homestay.Core) Homestay {
	return Homestay{
		Homestay_name: dataCore.Homestay_name,
		Price:         dataCore.Price,
		Total_guest:   dataCore.Total_guest,
		Description:   dataCore.Description,
		City_name:     dataCore.City_name,
		Address:       dataCore.Address,
	}
}

// mapping dari core ke gorm model Picture
func PictureCoreToModel(dataCore homestay.PictureCore) Picture {
	return Picture{
		HomestayID: dataCore.HomestayID,
		Url:        dataCore.Url,
	}
}

// mapping dari core ke gorm model Facility
func FacilityCoreToModel(dataCore homestay.FacilityCore) Facility {
	return Facility{
		Facility_name: dataCore.Facility_name,
	}
}

// mapping dari core ke gorm model HomestayFacility
func HomestayFacilityCoreToModel(dataCore homestay.HomestayFacility) HomestayFacility {
	return HomestayFacility{
		HomestayID: dataCore.HomestayID,
		FacilityID: dataCore.FacilityID,
	}
}

type Get_HomeStay struct {
	Homestay_name string
	City_name     string
	Address       string
	Price         int
	Description   string
	Rating        float64
	Url           string
}
