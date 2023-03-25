package main

import (
	"TopTips/config"
	"TopTips/controller"
	"TopTips/repository"
	"TopTips/service"

	"github.com/gin-gonic/gin"
)

var dbPool, err = config.NewDBPool(config.DataBaseConfig{
	Username: "postgres",
	Password: "123456",
	Hostname: "localhost",
	Port:     "5432",
	DBName:   "postgres",
})

var (
	authDB         = repository.NewDatabase(dbPool)
	jwtService     = service.NewJWTService()
	authService    = service.NewAuthService(authDB)
	authController = controller.NewAuthController(authService, jwtService)
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Top Tips now developing"})
	})

	auth := r.Group("/authorization")
	{
		auth.POST("/sign-in", authController.Login)
		auth.POST("/sign-up", authController.Register)
		auth.POST("/confirm", authController.ConfirmSMS)
		auth.POST("/forgot-password", authController.Forgot)
		//auth.POST("/sign-up-by-link")
		//auth.POST("/confirm-by-link")
	}

	r.Run(":8080")
}
