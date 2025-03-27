package dto

type SignUpInput struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	Password2 string `json:"password2" validate:"eqfield=Password"`
	Role      string `json:"role" validate:"required,role"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
