package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/medali/go-scraping/internal/sources/arabSeed"
)



func SearchWorkArabSeed(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type","application/json")
	params := mux.Vars(req)
	
	arabSeedWorks := arabSeed.ChoseMovie(params["query"])
	
	json.NewEncoder(res).Encode(arabSeedWorks)


}



// func ChooseQualityWeCinema(res http.ResponseWriter, req *http.Request){
// 	res.Header().Set("Content-Type","application/json")
	
// 	var link LinkForQuality
// 	_ = json.NewDecoder(req.Body).Decode(&link)
	
// 	links := wecinema.ChooseQuality(link.Url)
// 	fmt.Println(links)
// 	json.NewEncoder(res).Encode(links)
// }
