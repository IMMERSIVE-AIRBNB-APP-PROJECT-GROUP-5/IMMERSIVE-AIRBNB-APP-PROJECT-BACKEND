package handler

type UserRequest struct {
	Homestay_name string `form:"homestay_name"`
	Price         int    `form:"price"`
	Total_guest   int    `form:"total_guest"`
	Description   string `form:"description"`
	City_name     string `form:"city_name"`
	Address       string `form:"address"`
	Facility_Id   []uint `form:"facility"`
}
