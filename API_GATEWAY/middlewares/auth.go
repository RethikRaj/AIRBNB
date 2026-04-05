package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/RethikRaj/AIRBNB/API_GATEWAY/contextkeys"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/utils"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Step 1 : Check whether Authorization header exists
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			utils.WriteErrorJsonResponse(w, http.StatusBadRequest, "Invalid token", errors.New("Invalid Token"))
			return
		}

		// Step 2 : Check if bearer prefix
		if !strings.HasPrefix(authHeader, "Bearer ") {
			utils.WriteErrorJsonResponse(w, http.StatusBadRequest, "Invalid token format", errors.New("Invalid Token Format"))
			return
		}

		// Step 3 : Extract the token part
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Step 4 : Validate the token
		userID, email, isValid := utils.ValidateJWTToken(token)

		if !isValid {
			utils.WriteErrorJsonResponse(w, http.StatusBadRequest, "Invalid token", errors.New("Invalid token"))
			return
		}

		// Step 5 : Attach userID , email in context
		ctx := context.WithValue(r.Context(), contextkeys.UserID, userID)

		ctx = context.WithValue(ctx, contextkeys.UserEmail, email)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
