package service

import (
	"errors"

	"github.com/AIRBNBAPP/features/homestay"
	"github.com/AIRBNBAPP/features/user"

	"github.com/go-playground/validator/v10"
)

type HomestayService struct {
	HomestayData homestay.HomestayDataInterface
	validate     *validator.Validate
}

// GetAllHomestay implements homestay.HomestayServiceInterface.
func (service *HomestayService) GetAllHomestay(Search string) ([]homestay.Core, error) {
	data, err := service.HomestayData.GetAllHomestay(Search)
	return data, err
}

// CreateHomestay implements homestay.HomestayServiceInterface.
func (service *HomestayService) CreateHomestay(id int, userInput homestay.Core) error {
	var users user.Core
	// Cek apakah pengguna hosters atau bukan
	if users.Status != user.Hosters {
		return errors.New("Gagal membuat homestay, silahkan lakukan validasi hosters")
	}

	// Membuat variabel untuk menyimpan facility_id
	var facilityIDs []uint
	// Mengambil facility_id dari userInput.Facilities
	for _, facility := range userInput.Facilities {
		facilityIDs = append(facilityIDs, facility.Facility_Id)
	}
	// Mengatur ulang validator
	validate := validator.New()
	Register := homestay.Register{
		Homestay_name: userInput.Homestay_name,
		Price:         userInput.Price,
		Total_guest:   userInput.Total_guest,
		Description:   userInput.Description,
		City_name:     userInput.City_name,
		Address:       userInput.Address,
		Facility_Id:   facilityIDs, // Menggunakan facilityIDs yang sudah disiapkan
		Url:           userInput.Picture.Url,
	}

	errValidate := validate.Struct(Register)
	if errValidate != nil {
		return errValidate
	}

	// Cek apakah pengguna mengirimkan data kosong untuk semua bidang
	if userInput.Homestay_name == "" || userInput.Price == 0 || userInput.Total_guest == 0 || userInput.Description == "" || userInput.City_name == "" || userInput.Address == "" || userInput.Picture.Url == "" {
		return errors.New("Semua data harus diisi")
	}

	// Cek apakah pengguna mengirimkan data kosong untuk facility
	for _, f := range userInput.Facilities {
		if f.Facility_Id == 0 {
			return errors.New("Semua data harus diisi")
		}
	}

	errRegister := service.HomestayData.CreateHomestay(id, userInput)
	return errRegister
}

func New(repo homestay.HomestayDataInterface) homestay.HomestayServiceInterface {
	return &HomestayService{
		HomestayData: repo,
		validate:     validator.New(),
	}
}
