package db

import "errors"

var NotFound = errors.New("now was found the obejct")

func (store *Store) GetUserId(id int) (*User, error) {
	var user *User
	result := store.db.First(&user, id)
	if result.RowsAffected == 0 {
		return nil, NotFound
	}
	return user, nil
}
