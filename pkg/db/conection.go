package db

import (
	config "GlassGalore/pkg/config"

	domain "GlassGalore/pkg/domain"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{SkipDefaultTransaction: true})
	if dbErr != nil {
		return nil, dbErr
	}
	db.AutoMigrate(&domain.Users{})
	db.AutoMigrate(&domain.Admin{})
	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.Inventories{})
	db.AutoMigrate(&domain.Address{})
	db.AutoMigrate(domain.Cart{})
	db.AutoMigrate(domain.LineItems{})

	CheckAndCreateAdmin(db)

	return db, dbErr
}

func CheckAndCreateAdmin(db *gorm.DB) {
	var count int64
	db.Model(&domain.Admin{}).Count(&count)
	if count == 0 {
		password := "glassgalore"
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return
		}

		admin := domain.Admin{
			ID:       1,
			Name:     "glassgalore",
			Username: "glassgalore@gmail.com",
			Password: string(hashedPassword),
		}
		db.Create(&admin)
	}
}
