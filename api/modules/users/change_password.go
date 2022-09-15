package users

type ChangePassword struct {
	OldPassword string `json:"oldpassword" validate:"required"`
	NewPassword string `json:"newpassword" validate:"required"`
}
