package contextkeys

type contextKey string

const (
	CreateUserPayload contextKey = "createUserRequestPayload"
	SignInUserPayload contextKey = "signInUserRequestPayload"
	UserID            contextKey = "userID"
	UserEmail         contextKey = "userEmail"
)
