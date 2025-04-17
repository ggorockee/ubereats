package dto

type CreateAccountIn struct {
	Email     string `json:"email" validate:"email" mapstructure:"email"`
	Password  string `json:"password" validate:"eqfield=Password2,required" mapstructure:"password"`
	Password2 string `json:"password2" validate:"required" mapstructure:"password2"`
}
