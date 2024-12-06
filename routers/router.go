package routers

import (
	"go-project/controllers"
	"go-project/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetupRouter initializes
func SetupRouter() *gin.Engine {
	// MySQL connection
	dsn := "root:Admin@123@tcp(localhost:3306)/indore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	// Initialize the database
	models.InitializeDB(db)
	r := gin.Default()

	// Create a UserController instance
	userController := &controllers.UserController{DB: db}

	r.POST("/users", userController.CreateUser)       // Create a new user
	r.GET("/users", userController.GetAllUsers)       // Get all users
	r.GET("/users/:id", userController.GetUser)       // Get user by ID
	r.PUT("/users/:id", userController.UpdateUser)    // Update user by ID
	r.DELETE("/users/:id", userController.DeleteUser) // Delete user by ID

	return r
}
