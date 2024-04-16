package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/medali/go-scraping/routes"
	// "github.com/medali/go-scraping/internal/sources"
)











func main() {

 r := mux.NewRouter()

    // CORS middleware
    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:3000", "http://localhost:3000/"}),
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
        handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
    )

    // Apply CORS middleware to the router
    http.Handle("/", corsHandler(r))

    // Register routes
    routes.FetchWorks(r)


	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

// start server listen
// with error handling
	log.Fatal(http.ListenAndServe(":" + "8001", handlers.CORS(originsOk, headersOk, methodsOk)(r)))



}