package auth

// This struct contains the standard request body with its validations
// for the sign up endpoint.
type SignUpRequestBody struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	// An email address can contain at most 254 characters
	// more info: https://stackoverflow.com/questions/386294/what-is-the-maximum-length-of-a-valid-email-address#:~:text=%22There%20is%20a%20length%20limit,total%20length%20of%20320%20characters.
	Email string `json:"email" validate:"required" validate:"email" validate:"max=254"`
}

// This struct contains the standard request body with its validations
// for the sign up confirmation endpoint.
type SignUpConfirmationRequestBody struct {
	Token string `json:"token"`
	Email string `json:"email" validate:"required" validate:"email" validate:"max=254"`
}

// This struct contains the standard request body with its validations
// for the sign in endpoint.
type SignInRequestBody struct {
	// An email address can contain at most 254 characters
	// more info: https://stackoverflow.com/questions/386294/what-is-the-maximum-length-of-a-valid-email-address#:~:text=%22There%20is%20a%20length%20limit,total%20length%20of%20320%20characters.
	Email    string `json:"email" validate:"required" validate:"email" validate:"max=254"`
	Password string `json:"password" validate:"required,min=8"`
}
