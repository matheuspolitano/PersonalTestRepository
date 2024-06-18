package db

import (
	"errors"
	"time"

	"github.com/matheuspolitano/PersonalTestRepository/challenge/1-SecureUserAuthenticationService/utils"
)

var NotFound = errors.New("now was found the obejct")

func (store *Store) GetUserId(id int) (*User, error) {
	var user *User
	result := store.db.First(&user, id)
	if result.RowsAffected == 0 {
		return nil, NotFound
	}
	return user, nil
}

func (store *Store) CreateUser(firstName string, lastName string, email string, password string) (*User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}
	var user = &User{
		FirstName:           firstName,
		LastName:            lastName,
		Email:               email,
		HashPassword:        hashedPassword,
		CreatedAt:           time.Now(),
		PasswordLastChanged: time.Now(),
	}

	result := store.db.Create(&user)

	if err = result.Error; err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("row does not insert")
	}

	return user, nil

}
