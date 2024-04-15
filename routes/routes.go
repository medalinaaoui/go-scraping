package routes

import (
	"github.com/gorilla/mux"
	"github.com/medali/go-scraping/controllers"
)



func FetchWorks(router *mux.Router){
	router.HandleFunc("/work/{query}", controllers.SearchWork).Methods("GET")
	router.HandleFunc("/quality", controllers.ChooseQuality).Methods("POST")
	router.HandleFunc("/movie", controllers.GetMovie).Methods("POST")
	// router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	// router.HandleFunc("/book/{id}", controllers.GetSingleBook).Methods("GET")
	// router.HandleFunc("/book/{id}", controllers.UpdateSingleBook).Methods("PUT")
	// router.HandleFunc("/book/{id}", controllers.DeleteSingleBook).Methods("DELETE")
} 