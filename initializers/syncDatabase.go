package initializers

import "github.com/SMarsaDewo/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
