package utils

import (
	"errors"
)

var (
	dbHost = "localhost"
	dbPort = 5432
	//dbName string = os.Getenv("sc_database")
	//DbUser string = os.Getenv("sc_user")
	//DbPassword string = os.Getenv("sc_password")
	dbName string = "meme_database"
	DbUser string = "meme_admin"
	DbPassword string = "12345"
)

var ErrorDb error = errors.New("No connection with database")
var ErrorEmptyField error = errors.New("Fields must be non-empty")
var ErrorInvalidData error = errors.New("Invalid data")
var ErrorSomethingHappenedWrong error = errors.New("Something happened wrong.")
