package Seeders

import "gorm.io/gorm"

func Seed(DB *gorm.DB) {
	seedUsers(DB)
}
