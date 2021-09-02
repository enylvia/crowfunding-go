package main

import (
	"crowdfund-go/auth"
	"crowdfund-go/campaign"
	"crowdfund-go/handler"
	"crowdfund-go/helper"
	"crowdfund-go/user"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/crowdfund?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())

	}
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)

	campaigns, err := campaignRepository.FindAll()

	fmt.Println("debug")
	fmt.Println("debug")
	fmt.Println("debug")
	fmt.Println(len(campaigns))

	for _, v := range campaigns {
		fmt.Println(v.Name)
		if len(v.CampaignImages) > 0 {
			fmt.Println(v.CampaignImages[0].FileName)
		}
	}

	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	// Route Golang
	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil nilai header Authorization: Bearer (Generate token)
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Unauthorize", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		//split token "bearer tokentoken"
		tokenString := ""
		// dari header authorization kita ambil token nya saja
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		// validasi token yang didapat
		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.ApiResponse("Unauthorize", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.ApiResponse("Unauthorize", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		// ambiil user_id
		userID := int(claim["user_id"].(float64))
		// ambil user dari db berdasarkan user_id melalui service
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.ApiResponse("Unauthorize", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		// kalo usernya ada kita set context isinya user

		c.Set("currentUser", user)

	}
}
