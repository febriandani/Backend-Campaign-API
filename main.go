package main

import (
	"golang-startup-web/handler"
	"golang-startup-web/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	route := gin.Default()

	api := route.Group("/api/v1")
	api.POST("/regusers", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.GET("/email_check", userHandler.CheckEmailAvailable)

	route.Run()
}
