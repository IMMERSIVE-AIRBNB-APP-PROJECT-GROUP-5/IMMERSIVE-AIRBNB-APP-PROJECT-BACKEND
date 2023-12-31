package database

import (
	"fmt"

	"github.com/AIRBNBAPP/app/config"
	homestaydata "github.com/AIRBNBAPP/features/homestay/data"
	reservationdata "github.com/AIRBNBAPP/features/reservation/data"
	reviewdata "github.com/AIRBNBAPP/features/review/data"
	userdata "github.com/AIRBNBAPP/features/user/data"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(cfg *config.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

func InitialMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&userdata.User{},
		&homestaydata.Homestay{},
		&homestaydata.Picture{},
		&homestaydata.HomestayFacility{},
		&homestaydata.Facility{},
		&reservationdata.Reservation{},
		&reviewdata.Review{},
	)
	if err != nil {
		return err
	}
	return nil
}
