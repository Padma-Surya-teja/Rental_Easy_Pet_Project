package models

import (
	gorm "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func SetupDB(db_user *string, db_password *string, db_name *string) *gorm.DB {
	database_link := "user=" + *db_user + " password=" + *db_password + " dbname=" + *db_name + " sslmode=disable"

	db, err := gorm.Open("postgres", database_link)
	if err != nil {
		panic(err.Error())
	}

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
