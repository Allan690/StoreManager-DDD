package entity

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID ID
	Email string
	Password string
	FirstName string
	LastName string
	CreatedAt time.Time
	UpdatedAt time.Time
}


// NewUser creates a new user
func NewUser(email, password, firstName, lastName string) (*User, error) {
	user := &User{
		ID:        NewID(),
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	pwd, err := generatePassword(password)
	if err != nil {
		return nil, err
	}
	user.Password = pwd
	err = user.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return user, nil
}

// validates the user struct
func (u User) Validate() error {
	err := validation.ValidateStruct(&u,
		// last name cannot be empty and must be between 3 and 50 chars long
		validation.Field(&u.LastName, validation.Required, validation.Length(3, 50)),
		// first name cannot be empty and must be between 3 and 50 chars long
		validation.Field(&u.FirstName, validation.Required, validation.Length(3, 50)),
		// email cannot be empty and should be a valid email
		validation.Field(&u.Email, validation.Required, is.Email),
		)
	return err
}

// ValidatePassword validates user password
func (u *User) ValidatePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err
}

func generatePassword(password string)(string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), err
}
