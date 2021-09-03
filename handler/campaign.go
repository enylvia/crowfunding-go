package handler

import (
	"crowdfund-go/campaign"
	"crowdfund-go/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//tangkap parameter di handler
//kemudian handler ke service
//service yang menentukan apakah repository mana yang di call
// repository akses db

// repository dibuat 2 : findall dan findpostbyID

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

// func untuk /api/v1/
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.ApiResponse("Error to get Campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List of Campaigns", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusBadRequest, response)

}
