package users

import (
	"errors"
	"os"
	"strconv"
	"subscribers/domain"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	domain.Entity
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

func (u *User) ChangePassword(oldPassword, newPassword string) error {
	if !u.CheckPassword(oldPassword) {
		return errors.New("old password invalid")
	}

	salt, _ := strconv.Atoi(os.Getenv("sub_salt_hash"))
	newBytes, _ := bcrypt.GenerateFromPassword([]byte(newPassword), salt)
	newPasswordGeneraged := string(newBytes)
	u.PasswordHash = newPasswordGeneraged
	return nil
}

func NewUser(name, email, password string) (*User, error) {
	salt, _ := strconv.Atoi(os.Getenv("sub_salt_hash"))
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	if err != nil {
		return nil, errors.New("error to generate password")
	}
	passwordGeneraged := string(bytes)
	return &User{
		Name:         name,
		Email:        email,
		PasswordHash: passwordGeneraged,
		Entity:       domain.NewEntity(),
	}, nil
}
