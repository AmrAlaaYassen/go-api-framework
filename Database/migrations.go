package Database

import (
	"app.com/Models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Models.User{})
}
