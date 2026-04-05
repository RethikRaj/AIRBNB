package utils

import (
	"fmt"
	"time"

	config "github.com/RethikRaj/AIRBNB/API_GATEWAY/config/env"
	"github.com/golang-jwt/jwt/v5"
)

type MyCustomJWTClaims struct {
	UserID int    `json:"user_id"` // Must be exported because the JWT internally calls json.Marshal to convert the struct to JSON and json.Marshal can only access exported fields (fields that start with an uppercase letter)
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func CreateJWTToken(userID int, email string) (string, error) {
	claims := MyCustomJWTClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
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

func ValidateJWTToken(token string) (int, string, bool) {

	jwtToken, err := jwt.ParseWithClaims(
		token,
		&MyCustomJWTClaims{},
		func(t *jwt.Token) (any, error) {
			//  Validate the signing method to prevent algorithm confusion attacks
			// For example, if your server expects tokens to be signed with HMAC (HS256) and an attacker tries to use RSA (RS256) by changing the "alg" field in the token header, this check will prevent that attack by rejecting tokens with unexpected signing methods.
			// jwt.SigningMethodHS256 is of type *jwt.SigningMethodHMAC
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(config.GetStringValue("JWT_SECRET", "fallback_secret")), nil
		},
	)

	if err != nil {
		return 0, "", false
	}

	jwtClaims, ok := jwtToken.Claims.(*MyCustomJWTClaims)

	fmt.Println("JWT Claims : ", jwtClaims)

	if !ok {
		return 0, "", false
	}

	return jwtClaims.UserID, jwtClaims.Email, true
}
