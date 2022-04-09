package main

import (
	"golang-startup-web/handler"
	"golang-startup-web/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//connection database
	dsn := "host=localhost user=postgres password=junior34 dbname=startup port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if(err != nil){
		log.Fatal("DB Connection Error")
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")
	
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("avatars", userHandler.UploadAvatar)

	router.Run()
	//langkah-langkahnya yang harus dibuat sblm form html
	//5input : from user in form html
	//4handler : mapping input from user menjadi -> sebuah struct input
	//3service : melakukan mapping from struct input to struct user
	//2repositor 
	//1db
}