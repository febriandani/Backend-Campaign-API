package handler

import (
	"golang-startup-web/campaign"
	"golang-startup-web/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// tangkap parameter di handler
// handler ke service
// service menentukan apakah repository mana yang di call
// repository : FindAll, FindByUserID
// db

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("List of campaigns campaigns", http.StatusOK, "Success", campaigns )
		c.JSON(http.StatusOK, response)
		return
	}


}