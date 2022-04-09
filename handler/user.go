package handler

import (
	"fmt"
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

func (h *userHandler) Login(c *gin.Context){
		//step
		//user memasukan input berupa email dan password
		//input ditangkap handler 
		//mapping dari input user ke input struct
		//input struct passing ke service
		//didalam service mencari dengan bantuan repository user dengan email x
		//jika ketemu maka perlu mencocokkan password
		//mulai membuat folder dan coding dari pertama - ketiga
		//cara pertama membuat repository terlebih dahulu 
		//kedua membuat service
		//ketiga membuat handler

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

		formatter := user.FormatUser(loggedinUser, "tokentokentoken")

			response := helper.APIResponse("Login Successfully", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) CheckEmailAvailability(c *gin.Context){
	//ada input email dari user
	// input email di mapping ke struc input
	// struct input di passing ke service
	// service akan manggil repository untuk menentukan apakah email sudah ada di database atau belum
	// repo akan melakukan query ke database
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
	//tangkap input dari user
	//simpan gambar di folder "images/"
	//di service kita panggil repo
	//JWT (sementara pakai hardcode, seakan2 user yang login ID = 1)
	//repo ambil data user yg ID = 1
	//repo update data user simpan lokasi file

	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false }
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	//harusnya dapat dari jwt, sementara manual dulu
	userId := 1
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