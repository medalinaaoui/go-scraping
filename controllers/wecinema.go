package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/medali/go-scraping/internal/sources/wecinema"
)



func SearchWorkWeCinema(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type","application/json")
	params := mux.Vars(req)
	
	weCinemaWorks := wecinema.ChoseMovie(params["query"])
	
	json.NewEncoder(res).Encode(weCinemaWorks)


}



func ChooseQualityWeCinema(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type","application/json")
	
	var link LinkForQuality
	_ = json.NewDecoder(req.Body).Decode(&link)
	
	links := wecinema.ChooseQuality(link.Url)
	fmt.Println(links)
	json.NewEncoder(res).Encode(links)
}
