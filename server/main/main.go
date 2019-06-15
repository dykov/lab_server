package main

import (
	"../utils"
	"log"
	"net/http"
	"time"
)

func main(){

	utils.InitDatabase()

	GetRouter()

	server := &http.Server{
		Handler: router,
		Addr:    "localhost:7070", // 0.0.0.0:7070

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())

}
