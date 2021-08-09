package main

import (
	"crowdfund-go/user"
	"log"

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
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}

	userInput.Name = "Test Simpan dari service"
	userInput.Email = "contoh@mail.com"
	userInput.Occupation = "pelajar"
	userInput.Password = "12345678"

	userService.RegisterUser(userInput)
}
