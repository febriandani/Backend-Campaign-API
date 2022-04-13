package handler

import (
	"fmt"
	"golang-startup-web/auth"
	"golang-startup-web/helper"
	"golang-startup-web/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler{
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context){
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil{
		errors := helper.FormatterValidationsErr(err)
		errorMessage := gin.H{"errors": errors}

			response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	
	if err != nil{
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil{
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context){
		var input user.LoginInput

		err := c.ShouldBindJSON(&input)
		if err != nil {
			errors := helper.FormatterValidationsErr(err)
			errorMessage := gin.H{"errors": errors}

				response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
		}

		loggedinUser, err := h.userService.Login(input)

		if err != nil {
				errorMessage := gin.H{"errors": err.Error()}
				
					response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
				
		}

		token, err := h.authService.GenerateToken(loggedinUser.ID)
		if err != nil{
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
		formatter := user.FormatUser(loggedinUser, token)

			response := helper.APIResponse("Login Successfully", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) CheckEmailAvailability(c *gin.Context){
	var input user.CheckEmailInput

		err := c.ShouldBindJSON(&input)
		if err != nil {
			errors := helper.FormatterValidationsErr(err)
			errorMessage := gin.H{"errors": errors}

				response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
		}

		isEmailAvailable, err := h.userService.IsEmailAvailable(input)
		if err != nil {
			errorMessage := gin.H{"errors": "Server Error"}
				response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
		}

		data := gin.H{"is_available" : isEmailAvailable}

		var metaMessage string

		if isEmailAvailable{
			metaMessage = "Email is available"
		} else{
			metaMessage = "Email has been registered"
		}

		response := helper.APIResponse(metaMessage, http.StatusOK, "Success", data)
		c.JSON(http.StatusOK, response)		

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false }
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID
	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false }
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}


	_, err = h.userService.SaveAvatar(userId, path)
	if err != nil {
		data := gin.H{"is_uploaded": false }
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true }
		response := helper.APIResponse("Avatar successfully uploaded", http.StatusOK, "success", data)

		c.JSON(http.StatusOK, response)
		
}