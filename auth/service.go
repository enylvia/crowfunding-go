package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("crowfunding")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	//generate token

	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signToken, err
	}
	return signToken, nil

}

//melakukan validasi token apakah token ini valid atau tidak
