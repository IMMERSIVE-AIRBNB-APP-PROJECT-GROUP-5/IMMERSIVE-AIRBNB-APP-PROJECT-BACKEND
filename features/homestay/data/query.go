package data

import (
	"errors"

	"github.com/AIRBNBAPP/features/homestay"
	"gorm.io/gorm"
)

type homestayQuery struct {
	db *gorm.DB
}

// GetAllHomestay implements homestay.HomestayDataInterface.
func (repo *homestayQuery) GetAllHomestay(Search string) ([]homestay.Core, error) {
	var results []Homestay
	tx := repo.db
	if Search != "" {
		tx = tx.Where("homestay_name LIKE ?", "%"+Search+"%").
			Or("city_name LIKE ?", "%"+Search+"%").
			Or("address LIKE ?", "%"+Search+"%")
	}
	if err := tx.Preload("Picture").Preload("Review").Find(&results).Error; err != nil {
		return nil, err
	}

	var homestaysCoreAll []homestay.Core
	for _, value := range results {
		var homestayCore homestay.Core
		homestayCore.Homestay_name = value.Homestay_name
		homestayCore.City_name = value.City_name
		homestayCore.Price = value.Price
		homestayCore.Url = value.Picture.Url
		// Hitung rata-rata rating
		var totalRating int
		for _, review := range value.Review {
			totalRating += review.Rating
		}
		if len(value.Review) > 0 {
			averageRating := totalRating / len(value.Review)
			homestayCore.Rating = averageRating
		}

		homestaysCoreAll = append(homestaysCoreAll, homestayCore)
	}

	return homestaysCoreAll, nil
}

// CreateHomestay implements homestay.HomestayDataInterface.
func (repo *homestayQuery) CreateHomestay(id int, userInput homestay.Core) error {
	// Mencari pengguna berdasarkan ID
	var userData Homestay

	userData.UserID = uint(id)
	userData.Picture.Url = userInput.Picture.Url // Mengganti nilai URL di userData dengan URL dari userInput

	// mapping dari struct entities core ke gorm model
	userInputGorm := CoreToModel(userInput)
	userInputGorm.Picture = PictureCoreToModel(userInput.Picture)
	userInputGorm.UserID = userData.UserID

	tx := repo.db.Create(&userInputGorm) // Query insert ke database
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("Gagal menambahkan data homestay")
	}

	// Perbarui nilai FacilityID dan HomestayID di tabel HomestayFacility
	for _, facility := range userInput.Facilities {
		homestayFacility := HomestayFacilityCoreToModel(homestay.HomestayFacility{
			HomestayID: userInputGorm.ID,
			FacilityID: facility.Facility_Id,
		})
		tx := repo.db.Save(&homestayFacility)
		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

func New(db *gorm.DB) homestay.HomestayDataInterface {
	return &homestayQuery{
		db: db,
	}
}
