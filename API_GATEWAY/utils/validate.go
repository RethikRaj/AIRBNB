package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
var Validate *validator.Validate

func init() {
	fmt.Println("utils package init executed")
	Validate = validator.New(validator.WithRequiredStructEnabled())
}
