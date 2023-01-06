package models

import (
	"log"

	gorm "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"rental_easy.in/m/pkg/utils"
)

func SetupDB() *gorm.DB {

	db, err := gorm.Open("postgres", utils.GoDotEnvVariable("DATABASE_URL"))

	utils.CheckErr(err)
	log.Println(utils.GoDotEnvVariable("DATABASE_URL"))

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Item{})
	db.AutoMigrate(&Review{})
	db.AutoMigrate(&Booking{})
	db.Debug().Model(&Item{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Debug().Model(&Review{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Debug().Model(&Booking{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Debug().Model(&Review{}).AddForeignKey("item_id", "items(id)", "CASCADE", "CASCADE")
	db.Debug().Model(&Booking{}).AddForeignKey("item_id", "items(id)", "CASCADE", "CASCADE")

	return db
}
