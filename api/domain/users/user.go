package users

import (
	"errors"
	"os"
	"strconv"
	"subscribers/domain"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	*domain.Entity
	Name         string `gorm:"size:100; not null"`
	Email        string `gorm:"index;unique;size:100; not null"`
	PasswordHash string `gorm:"not null;size:125"`
}

func (u User) CheckPassword(password string) bool {
	if u.PasswordHash == "" {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

func NewUser(request CreationRequest) (*User, []error) {
	errs := domain.Validate(request)
	if errs != nil {
		return nil, errs
	}

	salt, _ := strconv.Atoi(os.Getenv("sub_salt_hash"))
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), salt)
	if err != nil {
		var errs = []error{errors.New("error to generate password")}
		return nil, errs
	}
	passwordGeneraged := string(bytes)
	return &User{
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: passwordGeneraged,
		Entity:       domain.NewEntity(),
	}, nil
}
