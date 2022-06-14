package handler

import (
	"golang-startup-web/campaign"
	"golang-startup-web/helper"
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
