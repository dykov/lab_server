package domains

import (
	"../utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

type SavedMeme struct {
	User uint `json:"user,omitempty"`
	Meme uint `json:"meme"`
}

type SavedMemes struct {
	SavedMeme []SavedMeme `json:"saved_meme"`
}

func GetSavedMemesByUserId( w http.ResponseWriter , r *http.Request )  {

	db := utils.GetDatabase()
	if db == nil {
		utils.CheckErr(w, http.StatusServiceUnavailable, utils.ErrorDb)
	}

	vars := mux.Vars(r)
	userId , err :=  strconv.ParseUint( vars["id"] , 10 , 64 )
	if err != nil {
		utils.CheckErr(w, http.StatusBadRequest, utils.ErrorSomethingHappenedWrong)
	}

	queryStmt, err := db.Prepare("select u.login , m.tags , m.likes " +
										"from saved_memes sm " +
										"join users u on sm.users = u.id " +
										"join memes m on sm.memes = m.id " +
										"where u.id = $1 ;" )
	if err != nil {
		utils.CheckErr(w, http.StatusServiceUnavailable, utils.ErrorSomethingHappenedWrong)
	}

	rows, err := queryStmt.Query(userId)

	defer rows.Close()
	utils.CheckErr(w, http.StatusBadRequest, err)

	type Saved struct {
		Login string
		Tags string
		Likes uint
	}

	saved := new( Saved )

	for rows.Next() {
		err = rows.Scan(
			&saved.Login ,
			&saved.Tags ,
			&saved.Likes ,
		)
	}

	outputJson, _ := json.Marshal(saved)
	fmt.Println(string(outputJson))

	var tmpl = template.Must(template.ParseGlob("../frontend/saved_memes/*"))
	tmpl.ExecuteTemplate(w, "saved_memes.html", saved)

}