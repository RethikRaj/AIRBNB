package middlewares

import (
	"errors"
	"net/http"
	"time"

	"github.com/RethikRaj/AIRBNB/API_GATEWAY/utils"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(rate.Every(1*time.Minute), 10) // 5 requests per minute

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !limiter.Allow() {
			utils.WriteErrorJsonResponse(w, http.StatusTooManyRequests, "Too many requests", errors.New("Too Many Requests"))
			return
		}

		next.ServeHTTP(w, r)
	})

}
