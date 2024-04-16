package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/medali/go-scraping/internal/sources/wecinema"
)

func ChooseQualityWeCinema(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type","application/json")
	
	var link LinkForQuality
	_ = json.NewDecoder(req.Body).Decode(&link)
	
	links := wecinema.ChooseQuality(link.Url)
	fmt.Println(links)
	json.NewEncoder(res).Encode(links)
}
