package main

import (
	"fmt"

	"github.com/qkveri/validator"
)

// MyStruct ..
type MyStruct struct {
	String string `validate:"is-awesome"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {
	validate = validator.New()
	validate.RegisterValidation("is-awesome", ValidateMyVal)

	s := MyStruct{String: "awesome"}

	err := validate.Struct(s)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}

	s.String = "not awesome"
	err = validate.Struct(s)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}
}

// ValidateMyVal implements validator.Func
func ValidateMyVal(fl validator.FieldLevel) error {
	if fl.Field().String() != "awesome" {
		return fmt.Errorf("not awesome")
	}

	return nil
}