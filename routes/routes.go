package routes

import (
	"github.com/gorilla/mux"
	"github.com/medali/go-scraping/controllers"
)



func FetchWorks(router *mux.Router){
	router.HandleFunc("/work/{query}", controllers.SearchWork).Methods("GET")
	router.HandleFunc("/episode", controllers.ChooseEpisode).Methods("POST")
	router.HandleFunc("/quality", controllers.ChooseQuality).Methods("POST")
	router.HandleFunc("/link", controllers.GetLink).Methods("POST")
} 