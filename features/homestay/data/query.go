package data

import (
	"errors"

	"github.com/AIRBNBAPP/features/homestay"
	"gorm.io/gorm"
)

type homestayQuery struct {
	db *gorm.DB
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
