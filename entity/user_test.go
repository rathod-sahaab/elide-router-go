package entity_test

import (
	"elide-router/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	email := "email@email.com"
	password := "password"
	name := "name"

	user, err := entity.NewUser(email, password, name)
	assert.Nil(t, err)

	assert.Equal(t, email, user.Email)
	assert.NotNil(t, user.ID)

	assert.NotEqual(t, password, user.Password)
}

func TestUserValidate(t *testing.T) {
	t.Run("Email validation", func(t *testing.T) {
		t.Run("Valid email", func(t *testing.T) {

			email := "email@email.com"
			password := "password"
			name := "name"

			user, err := entity.NewUser(email, password, name)
			assert.Nil(t, err)

			assert.Equal(t, email, user.Email)
			assert.NotNil(t, user.ID)

			assert.NotEqual(t, password, user.Password)
		})

		t.Run("Invalid email", func(t *testing.T) {
			email := "afed"
			password := "password"
			name := "name"

			user, err := entity.NewUser(email, password, name)
			assert.NotNil(t, err)

			assert.Equal(t, err, entity.ErrInvalidEntity)

			assert.Nil(t, user)
		})

		t.Run("Invalid email", func(t *testing.T) {
			email := "email.com"
			password := "password"
			name := "name"

			user, err := entity.NewUser(email, password, name)
			assert.NotNil(t, err)

			assert.Equal(t, err, entity.ErrInvalidEntity)

			assert.Nil(t, user)
		})

		t.Run("Empty email", func(t *testing.T) {
			email := ""
			password := "password"
			name := "name"

			user, err := entity.NewUser(email, password, name)
			assert.NotNil(t, err)

			assert.Equal(t, err, entity.ErrInvalidEntity)

			assert.Nil(t, user)
		})
	})
}
