package users

type CreateNewUser struct {
	Name     string `json:"name" validate:"required,gte=5"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=6"`
}
