package database

import (
	"qms/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	// Untuk Local Development
	// dsn := "root:password@tcp(localhost:3306)/namaDatabase"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// untuk Server Production
	dsn := "sql6441433:K7cwfKcqVs@tcp(sql6.freesqldatabase.com:3306)/sql6441433"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = db

	models.InsertBank(db)

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Bank{})
	db.AutoMigrate(&models.SlotBooking{})

	return db
}
