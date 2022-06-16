package entity

import (
	"time"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        ID
	Name      string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Links     []Link
}

func NewUser(email string, password string, name string) (*User, error) {
	// TODO: password validation
	password_hash, err := generatePassword(password)

	if err != nil {
		return nil, err
	}

	user := User{
		ID:        NewID(),
		Email:     email,
		Password:  password_hash,
		Name:      name,
		CreatedAt: time.Now(),
	}

	err = user.Validate()

	if err != nil {
		return nil, ErrInvalidEntity
	}

	return &user, nil
}

func (u *User) ValidatePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err != nil {
		return err
	}
	return nil
}

func (user *User) Validate() error {
	validate := validator.New()

	return validate.Struct(user)
}

func generatePassword(raw_password string) (string, error) {
	password_hash, err := bcrypt.GenerateFromPassword([]byte(raw_password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(password_hash), nil
}
