package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/RethikRaj/AIRBNB/API_GATEWAY/dto"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/utils"
)

func ReadAndValidateCreateUserRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("User middleware called ...")
		// 1. Decode the JSON body(Deserialization)
		var createUserRequestPayload dto.CreateUserRequest
		if err := utils.ReadJsonBody(r, &createUserRequestPayload); err != nil {
			utils.WriteErrorJsonResponse(w, http.StatusBadRequest, "Error decoding json", err)
			return
		}

		// 2. validate
		if err := utils.Validate.Struct(&createUserRequestPayload); err != nil {
			utils.WriteErrorJsonResponse(w, http.StatusBadRequest, "Invalid JSON", err)
			return
		}

		ctx := context.WithValue(r.Context(), "createUserRequestPayload", &createUserRequestPayload)

		// 2. Call next
		next.ServeHTTP(w, r.WithContext(ctx))

		fmt.Println("User middleware ended")
	})
}
