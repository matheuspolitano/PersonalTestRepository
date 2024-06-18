package main

import (
	"log"

	"github.com/matheuspolitano/PersonalTestRepository/challenge/1-SecureUserAuthenticationService/db"
	"github.com/matheuspolitano/PersonalTestRepository/challenge/1-SecureUserAuthenticationService/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	store, err := db.NewStore(config)
	if err != nil {
		log.Fatal(err)
	}

	err = store.AutoMigration()
	if err != nil {
		log.Fatal(err)
	}
}
