package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/medali/go-scraping/internal/sources"
	// "github.com/medali/test/pkg/models"
	// "github.com/medali/test/pkg/utils"
)



func SearchWork(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type","application/json")
	params := mux.Vars(req)
	works := sources.ChoseMovie(params["query"])
	json.NewEncoder(res).Encode(works)
}

type LinkForQuality struct {
Url       string    `json:"url"`
}


func ChooseQuality(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type","application/json")
	
	var link LinkForQuality
	_ = json.NewDecoder(req.Body).Decode(&link)
	
	links := sources.ChooseQuality(link.Url)
	fmt.Println(links)
	json.NewEncoder(res).Encode(links)
}

type LinkForMovie struct {
	Url       string    `json:"url"`
}


func GetMovie(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type","application/json")
	var link LinkForMovie
	_ = json.NewDecoder(req.Body).Decode(&link)
	directLink := sources.GetDownloadLinkDirect(link.Url)
	json.NewEncoder(res).Encode(directLink)
}


