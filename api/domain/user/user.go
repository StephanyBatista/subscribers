package user

import (
	"errors"
	"os"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `gorm:"size:100; not null"`
	Email        string `gorm:"index;unique;size:100; not null"`
	PasswordHash string `gorm:"not null"`
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
		Model:        gorm.Model{CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC()},
	}, nil
}
