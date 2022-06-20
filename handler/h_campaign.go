package handler

import (
	"golang-startup-web/campaign"
	"golang-startup-web/helper"
	"golang-startup-web/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	s campaign.Service
}

func NewCampaignHandler(s campaign.Service) *campaignHandler {
	return &campaignHandler{s}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	//tangkap parameter, karena parameter bertipe string maka convert menggunakan
	//strconv.Atoi menjadi integer lalu simpan di variable userID
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.s.GetCampaigns(userID)
	if err != nil {
		response := helper.APIresponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse("List of campaigns", http.StatusOK, "Success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	//handler : mapping id yang di url ke struct input => service, call formatter
	//service : input nya struct input => menangkap id di url => repository
	// repository : get campaign by id,

	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIresponse("Error to get detail campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.s.GetCampaignByID(input)
	if err != nil {
		response := helper.APIresponse("Error to get detail campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse("List of campaigns", http.StatusOK, "Success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIresponse("Failed to create campaign", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.s.CreateCampaign(input)
	if err != nil {
		response := helper.APIresponse("Failed to create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse("Success to create campaign", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)

}
