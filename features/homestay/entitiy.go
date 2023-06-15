package homestay

import (
	"time"
)

type Core struct {
	Homestay_Id   uint
	userID        uint
	City_name     string `json:"city_name" form:"city_name"`
	Homestay_name string `json:"homestay_name" form:"homestay_name"`
	Address       string `json:"address" form:"address"`
	Total_guest   int    `json:"total_guest" form:"total_guest"`
	Price         int    `json:"price" form:"price"`
	Description   string `json:"description" form:"description"`
	Picture       PictureCore
	Facilities    []FacilityCore
	Review        ReviewCore
	Rating        int    `json:"rating" form:"rating"`
	Url           string `json:"url" form:"url"`
	Created_at    time.Time
	Updated_at    time.Time
	Deleted_at    time.Time
}

type PictureCore struct {
	Picture_Id uint
	HomestayID uint
	Url        string `json:"url" form:"url"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}

type FacilityCore struct {
	Facility_Id   uint
	Facility_name string `json:"facility" form:"facility"`
	Created_at    time.Time
	Updated_at    time.Time
	Deleted_at    time.Time
}

type HomestayFacility struct {
	HomestayID uint
	FacilityID uint
	Created_at time.Time
	Deleted_at time.Time
}

type ReviewCore struct {
	ReservationID uint   `json:"reservation_id" form:"reservation_id"`
	HomestayID    uint   `json:"homestay_id" form:"homestay_id"`
	Rating        int    `json:"rating" form:"rating"`
	Comment       string `json:"comment" form:"comment"`
}

type HomestayDataInterface interface {
	CreateHomestay(id int, userInput Core) error
	GetAllHomestay(Search string) ([]Core, error)
	// CreateHomestayFacility(Homestay_Id uint, Facility_Id int) ([]FacilityCore, error)
	// GetHomestayByUserId(Homestay_Id uint, userID int) error
	// InsertHomestayPicture([]PictureCore) (Homestay_Id int)
	// GetHomestayId(Homestay_Id string) ([]Core, error)
	// UpdateUserById(Homestay_Id string, userInput Core) error
	// DeleteHomestayById(Homestay_Id string) error
}

type HomestayServiceInterface interface {
	CreateHomestay(id int, userInput Core) error
	GetAllHomestay(Search string) ([]Core, error)
}

type Register struct {
	Homestay_name string `json:"homestay_name" form:"homestay_name"`
	Price         int    `json:"price" form:"price"`
	Total_guest   int    `json:"total_guest" form:"total_guest"`
	Description   string `json:"description" form:"description"`
	City_name     string `json:"city_name" form:"city_name"`
	Address       string `json:"address" form:"address"`
	Url           string `json:"url" form:"url"`
	Facility_Id   []uint `json:"facility" form:"facility"`
}
