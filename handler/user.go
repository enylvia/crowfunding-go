package handler

import (
	"crowdfund-go/helper"
	"crowdfund-go/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)
		//gin.H adalah map yang key nya string namun valuenya interface
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to create account", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse("Failed to create account", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser, "tokentokentokentokentoken")
	response := helper.ApiResponse("Account has been created", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	// mapping kedlaam struct RegisterUserInput
	// kemudian struct diatas di passing sebagai parameter service

}

func (h *userHandler) Login(c *gin.Context) {

	// user input (email&password)
	// input ditangkap handler
	// mapping dari input user ke struct
	// input struct di passing ke service
	// didalam service mencari dengan bantuan repository user dengan email x
	// mencocokan password

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		//gin.H adalah map yang key nya string namun valuenya interface
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Failed to Login", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Failed to Login", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokencode")
	response := helper.ApiResponse("Successfuly Loggedin", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
