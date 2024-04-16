package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/medali/go-scraping/internal/sources/akwam"
	"github.com/medali/go-scraping/internal/sources/wecinema"
)


func SearchWork(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type","application/json")
	params := mux.Vars(req)
	akwamWorks := akwam.ChoseMovie(params["query"])
	weCinemaWorks := wecinema.ChoseMovie(params["query"])
	
	
// Convert weCinemaWorks to akwam.Work type
var works []akwam.Work
		for _, w := range weCinemaWorks {
   		 works = append(works, akwam.Work(w))
	}

	works = append(works, akwamWorks...)
	json.NewEncoder(res).Encode(works)


}





func SearchWorkAkwam(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type","application/json")
	params := mux.Vars(req)
	akwamWorks := akwam.ChoseMovie(params["query"])
	json.NewEncoder(res).Encode(akwamWorks)


}










type LinkForQuality struct {
Url       string    `json:"url"`
}


func ChooseQuality(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type","application/json")
	
	var link LinkForQuality
	_ = json.NewDecoder(req.Body).Decode(&link)
	
	links := akwam.ChooseQuality(link.Url)
	fmt.Println(links)
	json.NewEncoder(res).Encode(links)
}



type LinkForEpisode struct {
Url       string    `json:"url"`
}


func ChooseEpisode(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type","application/json")
	
	var link LinkForEpisode
	_ = json.NewDecoder(req.Body).Decode(&link)
	
	links := akwam.ChooseEpisode(link.Url)
	fmt.Println(links)
	json.NewEncoder(res).Encode(links)
}







type LinkForMovie struct {
	Url       string    `json:"url"`
}


func GetLink(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type","application/json")
	var link LinkForMovie
	_ = json.NewDecoder(req.Body).Decode(&link)
	directLink := akwam.GetDownloadLinkDirect(link.Url)
	json.NewEncoder(res).Encode(directLink)
}


