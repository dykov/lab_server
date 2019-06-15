package domains

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"../utils"
	"strconv"
)

type User struct {
	Id uint `json:"id"`
	Login string `json:"login"`
	Password string `json:"password"`
	Photo string `json:"photo"`
}

type Users struct {
	User []User `json:"user"`
}

func AddUser( w http.ResponseWriter , r *http.Request)  {

	db := utils.GetDatabase()
	if db == nil {
		utils.CheckErr(w, http.StatusServiceUnavailable, utils.ErrorDb)
	}

	var users Users
	json.NewDecoder(r.Body).Decode(&users)

	queryStr := fmt.Sprintf(
		"call add_user('%s','%s','%s');" ,
		users.User[0].Login,
		users.User[0].Password,
		users.User[0].Photo ,
		)

	fmt.Println(queryStr)

	if db.QueryRow( queryStr ) == nil {
		utils.CheckErr(w , http.StatusBadRequest , utils.ErrorInvalidData)
	}

}

func Settings( w http.ResponseWriter , r *http.Request )  {

	db := utils.GetDatabase()
	if db == nil {
		utils.CheckErr(w, http.StatusServiceUnavailable, utils.ErrorDb)
	}

	vars := mux.Vars(r)
	userId , err :=  strconv.ParseUint( vars["id"] , 10 , 64 )
	if err != nil {
		utils.CheckErr(w, http.StatusBadRequest, utils.ErrorSomethingHappenedWrong)
	}

	queryStmt, err := db.Prepare("select login from users where id = $1 ;")
	if err != nil {
		utils.CheckErr(w, http.StatusServiceUnavailable, utils.ErrorSomethingHappenedWrong)
	}

	rows, err := queryStmt.Query(userId)
	defer rows.Close()
	utils.CheckErr(w, http.StatusBadRequest, err)

	var users Users
	for rows.Next() {

		user := new( User )

		err = rows.Scan(
			&user.Login ,
		)

		users.User = append( users.User , *user )

	}

	fmt.Println( users.User[0] )

	tmpl, _ := template.ParseFiles("../frontend/settings/settings.html")
	tmpl.Execute(w, users.User[0])

}

