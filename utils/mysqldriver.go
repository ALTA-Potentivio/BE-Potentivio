package utils

import (
	"fmt"
	"potentivio-app/configs"
	"potentivio-app/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}

	InitialMigration(db)
	return db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&entities.Catagory{})
	db.AutoMigrate(&entities.Genre{})
	db.AutoMigrate(&entities.Artist{})
	db.AutoMigrate(&entities.Cafe{})
	db.AutoMigrate(&entities.ImageCafe{})
	db.AutoMigrate(&entities.VideoArtist{})
	db.AutoMigrate(&entities.Hire{})
	db.AutoMigrate(&entities.Notification{})
	db.AutoMigrate(&entities.Rating{})
}
