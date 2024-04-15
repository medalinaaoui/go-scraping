package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/medali/go-scraping/routes"
	// "github.com/medali/go-scraping/internal/sources"
)











func main() {





	// for {



	// 	var choosenSource string
	// 	fmt.Print("Choose source: ")
	// 	fmt.Scanln(&choosenSource)

	// 	if choosenSource == "Akwam" {
	// 		sources.Akwam()
	// 	} else {
	// 	fmt.Print("source doesn't exists")
			
	// 	}




		
	
	// 	// Prompt the user if they want to continue
	// 	var continueOption string
	// 	fmt.Print("Do you want to continue? (yes/no): ")
	// 	fmt.Scanln(&continueOption)
	// 	if continueOption != "yes" {
	// 		break 
	// 	}
	// }


	r := mux.NewRouter()
	routes.FetchWorks(r)
	http.Handle("/", r)

	if err := http.ListenAndServe(":8001", r); err == nil {
		log.Fatal(err)
	}
}