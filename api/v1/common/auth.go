package common

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

var validate = validator.New()

var decoder = schema.NewDecoder()

func handleParseForm(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "invalid form")
		return err
	}

	return nil
}

// This function will attempt to serialize contents of the parsed HTTP form
// into our desired struct.
func handleDecodeBody[T any](w http.ResponseWriter, r *http.Request, repr *T) error {
	if err := decoder.Decode(repr, r.Form); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "decode error")
		return err
	}

	return nil
}

// This function will take care of validating the HTTP form by using
// a predefined schema struct.
func handleValidateBody[T any](w http.ResponseWriter, r *http.Request, repr T) error {
	if err := validate.Struct(repr); err != nil {
		err := err.(validator.ValidationErrors)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "invalid form: "+err.Error())
		return err
	}

	return nil
}

func baseValidationHandler[T any](w http.ResponseWriter, r *http.Request, repr *T) error {
	if err := handleParseForm(w, r); err != nil {
		return err
	}

	if err := handleDecodeBody(w, r, repr); err != nil {
		return err
	}

	if err := handleValidateBody(w, r, repr); err != nil {
		return err
	}

	return nil
}

// This struct contains the standard request body with its validations
// for the sign up endpoint.
type SignUpRequestBody struct {
	// An email address can contain at most 254 characters
	// more info: https://stackoverflow.com/questions/386294/what-is-the-maximum-length-of-a-valid-email-address#:~:text=%22There%20is%20a%20length%20limit,total%20length%20of%20320%20characters.
	Email    string `json:"email" validate:"required,email,max=254"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8,alphanum"`
}

func (b *SignUpRequestBody) Validate(w http.ResponseWriter, r *http.Request) error {
	err := baseValidationHandler(w, r, b)
	return err
}

// This struct contains the standard request body with its validations
// for the sign in endpoint.
type SignInRequestBody struct {
	// An email address can contain at most 254 characters
	// more info: https://stackoverflow.com/questions/386294/what-is-the-maximum-length-of-a-valid-email-address#:~:text=%22There%20is%20a%20length%20limit,total%20length%20of%20320%20characters.
	Email    string `json:"email" validate:"required,email,max=254"`
	Password string `json:"password" validate:"required,min=8"`
}

func (b *SignInRequestBody) Validate(w http.ResponseWriter, r *http.Request) error {
	err := baseValidationHandler(w, r, b)
	return err
}
