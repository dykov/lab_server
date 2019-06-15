package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var DB *sql.DB

func InitDatabase() {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, DbUser, DbPassword, dbName)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Infof("Database: failed to connect database: %s", dbInfo)
		db = nil
	} else {
		log.Infof("Database: connected with database: %s", dbInfo)
	}

	DB = db

}

func GetDatabase() *sql.DB {
	if DB == nil {
		InitDatabase()
	}
	return DB
}

func CheckErr(w http.ResponseWriter, status int, err error) {
	if err != nil {
		w.WriteHeader(status)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
	}
}