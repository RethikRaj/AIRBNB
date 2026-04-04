package middlewares

import "net/http"

// This is a middleware just to understand how middlewares work in go

func MyMiddleware(next http.Handler) http.Handler {
	// Normal function
	myFunc := func(w http.ResponseWriter, r *http.Request) {
		// 1. Code to run BEFORE the actual handler (e.g., Auth check)

		next.ServeHTTP(w, r) // Call the next handler in the chain

		// 2. Code to run AFTER the actual handler (e.g., Logging)
	}

	// http.HandlerFunc is not a function that accepts another function
	// http.HandlerFunc is a type that implements http.Handler interface
	// The below is type conversion just like string(65)
	// Here we are doing type conversion : normalFunc -> http.HandlerFunc
	// We need this conversion because plain func doesn't satisfy http.Handler interface
	httpHandlerFunc := http.HandlerFunc(myFunc)

	return httpHandlerFunc
}

// Cleaner and short way to write above
func CustomMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Code to run BEFORE the actual handler (e.g., Auth check)

		next.ServeHTTP(w, r) // Call the next handler in the chain

		// 2. Code to run AFTER the actual handler (e.g., Logging)
	})
}
