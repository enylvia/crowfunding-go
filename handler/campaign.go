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
	response := helper.ApiResponse("List of Campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusBadRequest, response)

}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	//
	//handler untuk maping id yg di url ke struct input untuk service memanggil formatter
	//service nya struct input untuk menangkap id di url -> memanggil repo
	//repository : get campaign by id

	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.ApiResponse("Failed to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	campaignDetail, err := h.service.GetCampaignByID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Campaign detail", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}

//tangkap parameter dari user ke input struct
//ambil currentuser dari jwt/handler
//panggil service dimana service parameter nya adalah input struct tadi (slug berdasarkan nama campaign)
//pangigl repository untuk save data campaign
