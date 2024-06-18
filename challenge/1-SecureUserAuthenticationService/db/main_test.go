package db

import (
	"log"
	"os"
	"testing"

	"github.com/matheuspolitano/PersonalTestRepository/challenge/1-SecureUserAuthenticationService/utils"
)

var testStore *Store

func TestMain(m *testing.M) {
	conf, err := utils.LoadConfig("../.")
	if err != nil {
		log.Fatal(err)
	}

	testStore, err = NewStore(conf)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}
