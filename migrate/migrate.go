package migrate

import (
	"go-crud/initializers"
	"go-crud/models"
)
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	// Auto migrat
	initializers.DB.AutoMigrate(&models.Post{})
}

