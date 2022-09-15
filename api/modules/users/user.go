package users

import (
	"errors"
	"net/mail"
	"os"
	"strconv"
	"time"

	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           string
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

func (u User) CheckPassword(password string) bool {
	if u.PasswordHash == "" {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

func (u *User) ChangePassword(oldPassword, newPassword string) []error {

	var errs []error
	if oldPassword == "" {
		errs = append(errs, errors.New("'OldPassword' is required"))
	}
	if newPassword == "" {
		errs = append(errs, errors.New("'NewPassword' is required"))
	}
	if len(errs) > 0 {
		return errs
	}

	if !u.CheckPassword(oldPassword) {
		return []error{errors.New("old password invalid")}
	}

	salt, _ := strconv.Atoi(os.Getenv("sub_salt_hash"))
	newBytes, _ := bcrypt.GenerateFromPassword([]byte(newPassword), salt)
	newPasswordGeneraged := string(newBytes)
	u.PasswordHash = newPasswordGeneraged
	return nil
}

func NewUser(name, email, password string) (User, []error) {

	errs := validate(name, email, password)
	if len(errs) > 0 {
		return User{}, errs
	}

	salt, _ := strconv.Atoi(os.Getenv("sub_salt_hash"))
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	if err != nil {
		return User{}, []error{errors.New("error to generate password")}
	}
	passwordHash := string(bytes)
	return User{
		Id:           xid.New().String(),
		CreatedAt:    time.Now().UTC(),
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
	}, nil
}

func validate(name string, email string, password string) []error {
	var errs []error
	if name == "" {
		errs = append(errs, errors.New("'Name' is required"))
	} else if len(name) < 4 {
		errs = append(errs, errors.New("'Name' invalid size, min 4"))
	}
	if email == "" {
		errs = append(errs, errors.New("'Email' is required"))
	} else if _, err := mail.ParseAddress(email); err != nil {
		errs = append(errs, errors.New("'Email' invalid"))
	}
	if password == "" {
		errs = append(errs, errors.New("'Password' is required"))
	} else if len(password) < 6 {
		errs = append(errs, errors.New("'Password' invalid size, min 6"))
	}
	return errs
}
