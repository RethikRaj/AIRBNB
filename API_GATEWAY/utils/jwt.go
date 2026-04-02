package utils

import (
	"time"

	config "github.com/RethikRaj/AIRBNB/API_GATEWAY/config/env"
	"github.com/golang-jwt/jwt/v5"
)

type MyCustomJWTClaims struct {
	userID int
	jwt.RegisteredClaims
}

func CreateJWTToken(userID int) (string, error) {
	claims := MyCustomJWTClaims{
		userID,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	// Create a token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token
	signedToken, err := token.SignedString([]byte(config.GetStringValue("JWT_SECRET", "fallback_secret")))

	return signedToken, err
}
