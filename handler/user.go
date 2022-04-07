package handler

import (
	"golang-startup-web/helper"
	"golang-startup-web/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler{
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context){
	//tangkap input dari user
	// mapping input dari user ke struct RegisterUserInput
	// struct diatas kita passing ke parameter service

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

	//token, err := h.jwtService.GenerateToken()

	formatter := user.FormatUser(newUser, "tokentokentoken")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusOK, response)
}