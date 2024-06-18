package db

import (
	"errors"
	"time"

	"github.com/matheuspolitano/PersonalTestRepository/challenge/1-SecureUserAuthenticationService/utils"
)

func (store *Store) GetUserId(id int) (*User, error) {
	var user *User
	result := store.db.First(&user, id)
	if result.RowsAffected == 0 {
		return nil, NotFound
	}
	return user, nil
}

func (store *Store) DeleteById(id int) error {
	result := store.db.Delete(&User{}, id)
	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return errors.New("row does not delete")
	}

	return nil

}

func (store *Store) ChangePassword(id int, newPassword string) error {
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	result := store.db.Model(&User{
		UserID: id,
	}).Updates(User{
		HashPassword: hashedPassword,
	})

	if err = result.Error; err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return errors.New("row does not change password")
	}

	return nil
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
