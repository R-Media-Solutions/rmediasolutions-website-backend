package models

import (
	"golang.org/x/crypto/bcrypt"
)

type AdmUser struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(300)" json:"name"`
	Username string `gorm:"type:varchar(300)" json:"username"`
	Email    string `gorm:"type:varchar(100)" json:"email"`
	Password string `gorm:"type:varchar(100)" json:"password"`
}

func (user *AdmUser) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *AdmUser) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
