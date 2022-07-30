package user

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

func NewUser(name string, email string, password string) (*User, error) {
	salt, _ := strconv.Atoi(os.Getenv("sub_salt_hash"))
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	if err != nil {
		return nil, errors.New("error to generate the password hash")
	}
	passwordGeneraged := string(bytes)
	return &User{
		Name:         name,
		Email:        email,
		PasswordHash: passwordGeneraged,
		Entity:       domain.NewEntity(),
	}, nil
}
