package main

import (
	"../domains"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)


var router *mux.Router

func GetRouter(){

	router = mux.NewRouter()

	router.HandleFunc("/memes/new", domains.GetMemes).Methods("GET")
	router.HandleFunc("/memes/top", domains.GetTopMemes).Methods("GET")
	router.HandleFunc("/like", domains.LikeMeme).Methods("POST")


	//router.HandleFunc("/memes/time/{time_from}/{time_to}", Test).Methods("GET")

	router.HandleFunc("/registration", domains.AddUser).Methods("POST")

	router.HandleFunc("/test", Test ).Methods("GET")

	/* Templates */

	router.HandleFunc( "/settings/{id}" , domains.Settings )
	router.HandleFunc("/saved/{id:[0-9]+}", domains.GetSavedMemesByUserId).Methods("GET")


}

func Test( w http.ResponseWriter , r *http.Request ) {

	type Data struct {
		UserName string
	}

	fmt.Println("qwe")

	var data Data
	data.UserName = "QWERTY"

	tmpl, _ := template.ParseFiles("../frontend/settings/settings.html")

	tmpl.Execute(w, data)

}
