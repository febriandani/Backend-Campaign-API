package main

import (
	"fmt"
	"golang-startup-web/auth"
	"golang-startup-web/campaign"
	"golang-startup-web/handler"
	"golang-startup-web/helper"
	"golang-startup-web/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	campaignRepository := campaign.NewRepository(db)

	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	authService := auth.NewService()

	campaigns, _ := campaignService.FindCampaigns(9)
		fmt.Println(len(campaigns))

	
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")
	
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/campaigns")
	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context){
				authHeader := c.GetHeader("Authorization")

				if !strings.Contains(authHeader, "Bearer"){
						response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
						c.AbortWithStatusJSON(http.StatusUnauthorized, response)
						return
				}

				tokenString := ""
				arrayToken := strings.Split(authHeader, " ")
				if len(arrayToken) == 2{
					tokenString = arrayToken[1]
				}

				token, err := authService.ValidateToken(tokenString)
					if err != nil {
							response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
						c.AbortWithStatusJSON(http.StatusUnauthorized, response)
						return
					}

					claim, ok := token.Claims.(jwt.MapClaims)
					if !ok || !token.Valid {
							response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
						c.AbortWithStatusJSON(http.StatusUnauthorized, response)
						return
					}

					userID := claim["user_id"].(float64)

					user, err := userService.GetUserByID(int(userID))
					if err != nil {
							response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
						c.AbortWithStatusJSON(http.StatusUnauthorized, response)
						return
					}

					c.Set("currentUser", user)
				}
}


