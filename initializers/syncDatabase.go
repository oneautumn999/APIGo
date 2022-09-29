package initializers

import "btpn/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Photo{})
}
