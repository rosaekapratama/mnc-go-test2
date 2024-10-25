package crypto

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Sub         string `json:"sub"`
	PhoneNumber string `json:"phone_number"`
	jwt.RegisteredClaims
}
