package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

type User struct {
	ID uint `gorm:"primary_key"`
	Name	string
	FirstName	string
	LastName	string
}

func main() {
	const addr = "postgresql://root@localhost:26257/postgres?sslmode=disable"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	//db.CreateTable(Product{})
	err = db.AutoMigrate(&Product{}, &User{}).Error
	fmt.Println(err)
	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})
	db.Create(&User{
		Name:      "skynet",
		FirstName: "Anh",
		LastName:  "Dam Viet",
	})

	// Read
	var user User
	err = db.First(&user, ).Error
	fmt.Printf("User: %+v \n", user)
	var product Product
	err = db.First(&product, "code = ?", "L1212").Error // find product with code l1212
	logInfo(product, err)
	// Update - update product's price to 2000
	err = db.Model(&product).Update("Price", 2000).Error
	logInfo(product, err)
	// Delete - delete product
	// db.Delete(&product)
}

func logInfo(p Product, err error)  {
	if err == nil {
		fmt.Printf("Product info: %+v \n", p)
	}else{
		fmt.Println("Error: ", err)
	}
}