package db

import (
	"fmt"

	"github.com/matheuspolitano/PersonalTestRepository/challenge/1-SecureUserAuthenticationService/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(config *utils.Config) (*Store, error) {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHostname, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Store{
		db: db,
	}, nil
}

func (store *Store) AutoMigration() error {
	return store.db.AutoMigrate(&User{})
}
