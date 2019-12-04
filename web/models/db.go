package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func StartDatabase() (err error) {
	const addr = "postgresql://root@localhost:26257/postgres?sslmode=disable"
	db, err = gorm.Open("postgres", addr)
	if err != nil {
		return
	}
	err = db.AutoMigrate(&User{}, &Blog{}).Error
	admin := User{
		Email:     "skynet0590@gmail.com",
		Password:  "123456",
		FirstName: "Anh",
		LastName:  "Dam Viet",
	}
	admin.Create()
	db.LogMode(true)
	return
}
