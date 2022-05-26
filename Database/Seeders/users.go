package Seeders

import (
	"app.com/Models"
	"gorm.io/gorm"
)

func seedUsers(DB *gorm.DB) {
	DB.Create(&Models.User{Username: "adminUser", Password: "admin123", Email: "admin@email.com", Group: "admin", Token: "xyz"})
}
