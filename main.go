package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)
type Movie struct {
	Index int    `json:"index"`
	Name string `json:"name"`
	Url string `json:"url"`
}

func choseMovie(s string) []Movie{

	query := strings.ReplaceAll(s, " ", "+") 
	url := fmt.Sprintf("https://ak.sv/search?q=%s", query)
	c := colly.NewCollector()

	var movies []Movie
	index := 1
	c.OnHTML(".entry-box .entry-image a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		movie := Movie{
			Index: index,
			Name: e.ChildAttr("img", "alt"),
			Url: link,
		}
		index++
		movies = append(movies, movie)
	})	
	
	c.Visit(url)
	// c.OnScraped(func(r *colly.Response) {
	// 	for i := 0; i < len(movies); i++ {
	// 		movies[i].Index = i + 1
	// 	}
	// })

	return movies
}

type Link struct{
	Index int    `json:"index"`
	Quality string `json:"quality"`
	Size string `json:"size"`
	Url string `json:"url"` 
}


func getLink(url string) []Link{

	
	c := colly.NewCollector()

	var links []Link
	index := 1

	c.OnHTML("[id^='tab-']", func(e *colly.HTMLElement) {
		url := e.ChildAttr(".link-download", "href")

		quality := ""
		if e.Attr("id") == "tab-5" {
			quality = "1080p"
			} else if e.Attr("id") == "tab-4"{
				
				quality = "720p"
				} else{

					quality = "480p"
		}

link := Link{
			Index: index,
			Quality: quality,
			Size: e.ChildText(".font-size-14"),
			Url: url,
		}
		index++

		links = append(links, link)
	})	
	
	c.Visit(url)
	return links
}


func main() {
	
	movies := choseMovie("whiplash")

	for _, v := range movies {

		fmt.Printf("\n%v ===> %v\n", v.Index, v.Name)
	}
	var choosenMovie int
	fmt.Print("choose movie: ")

	fmt.Scanln(&choosenMovie)


	links := getLink(movies[choosenMovie  - 1].Url)

for _, v := range links {
		fmt.Printf("\n%v ===> %v ===> %v\n", v.Index, v.Quality, v.Size)
	}


}