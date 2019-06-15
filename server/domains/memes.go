package domains

import (
	"../utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type Meme struct {
	Id uint `json:"id"`
	Image string `json:"image"`
	Tags string `json:"tags"`
	Owner uint `json:"owner"`
	Likes uint `json:"likes"`
	Date string `json:"date"`
}

type Memes struct {
	Meme []Meme `json:"meme"`
}

func GetMemes( w http.ResponseWriter , r *http.Request )  {

	db := utils.GetDatabase()
	if db == nil {
		utils.CheckErr(w, http.StatusServiceUnavailable, utils.ErrorDb)
	}

	queryStr := `select * from memes order by date;`
	rows, err := db.Query(queryStr)
	utils.CheckErr(w, http.StatusBadRequest, err)

	var memes Memes
	memes.Meme = make( []Meme , 0 )

	for rows.Next() {

		meme := new( Meme )

		err = rows.Scan(
			&meme.Id ,
			&meme.Image ,
			&meme.Tags ,
			&meme.Owner ,
			&meme.Likes ,
			&meme.Date ,
		)

		memes.Meme = append( memes.Meme , *meme )

	}

	outputJson, _ := json.Marshal(memes)
	fmt.Fprintln(w, string(outputJson))

}

func GetTopMemes( w http.ResponseWriter , r *http.Request )  {

	db := utils.GetDatabase()
	if db == nil {
		utils.CheckErr(w, http.StatusServiceUnavailable, utils.ErrorDb)
	}

	queryStr := `select * from memes order by likes desc;`
	rows, err := db.Query(queryStr)
	utils.CheckErr(w, http.StatusBadRequest, err)

	var memes Memes
	memes.Meme = make( []Meme , 0 )

	for rows.Next() {

		meme := new( Meme )

		err = rows.Scan(
			&meme.Id ,
			&meme.Image ,
			&meme.Tags ,
			&meme.Owner ,
			&meme.Likes ,
			&meme.Date ,
		)

		memes.Meme = append( memes.Meme , *meme )

	}

	outputJson, _ := json.Marshal(memes)
	fmt.Fprintln(w, string(outputJson))

}

func LikeMeme( w http.ResponseWriter , r *http.Request ){

	db := utils.GetDatabase()
	if db == nil {
		utils.CheckErr(w, http.StatusServiceUnavailable, utils.ErrorDb)
	}

	var savedMemes SavedMemes
	json.NewDecoder(r.Body).Decode(&savedMemes)

	sqlQuery := fmt.Sprintf(
		"call like_meme( %d , %d ) ;" ,
		savedMemes.SavedMeme[0].User ,
		savedMemes.SavedMeme[0].Meme ,
		)

	fmt.Println( sqlQuery )
	_, err := db.Exec( sqlQuery )
	if err != nil {
		utils.CheckErr(w, http.StatusServiceUnavailable, utils.ErrorSomethingHappenedWrong)
	}


}
