package model

type UserRegister struct {
	Username string `json:"username" validate:"required,min=6,max=20"`
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"full_name" validate:"required"`
	Password string `json:"password" validate:"required,min=12"`
}

type UserLogin struct {
	Username string `json:"username" validate:"required,min=6,max=20"`
	Password string `json:"password" validate:"required,min=12"`
}
