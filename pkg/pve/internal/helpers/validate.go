package helpers

import "github.com/go-playground/validator/v10"

// validate is a global validator instance.
var Validate = validator.New(
	validator.WithRequiredStructEnabled(),
)
