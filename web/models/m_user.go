package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type (
	User struct {
		gorm.Model
		Email	string		`gorm:"type:varchar(100);unique_index"`
		Password	string 	`json:"-"`
		FirstName	string
		LastName	string
	}
	UserFilter struct {

	}
)

func (u *User) Create() error {
	err := u.generatePasswordHash(u.Password)
	if err != nil {
		return err
	}
	return db.Save(u).Error
}

func (u *User) generatePasswordHash(pw string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), 10)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) Update() error {
	u.Password = ""
	return db.Model(&User{}).Updates(u).Error
}

func (u *User) Save() error {
	return db.Save(u).Error
}

func (u *User) GetForLogin() (msg string, err error) {
	var user User
	err = db.Where("email = ?", u.Email).First(&user).Error
	if err != nil {
		msg = "Không tìm thấy người dùng"
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		msg = "Mật khẩu không hợp lệ"
		return
	}
	*u = user
	return
}