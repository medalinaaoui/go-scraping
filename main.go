package main

import (
	"bufio"
	"fmt"
	"os"
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













func getDownloadLink(url string) string{

	c := colly.NewCollector()

	var link string
	c.OnHTML(".download-link", func(e *colly.HTMLElement) {
		link = e.Attr("href")
	})	
	
	c.Visit(url)
	return link
}



func getDownloadLinkDirect(url string) string{

	c := colly.NewCollector()

	var link string
	c.OnHTML(".btn-loader", func(e *colly.HTMLElement) {
		link = e.ChildAttr("a", "href")
	})	
	
	c.Visit(url)
	return link
}













func main() {
	for {
		
		fmt.Println("Enter the work name:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		movie := scanner.Text()
		movies := choseMovie(movie)

		for _, v := range movies {
			fmt.Printf("\n%v ===> %v\n", v.Index, v.Name)
		}
		var choosenMovie int
		fmt.Print("Choose work: ")
		fmt.Scanln(&choosenMovie)

		links := getLink(movies[choosenMovie-1].Url)

		for _, v := range links {
			fmt.Printf("\n%v ===> %v ===> %v \n", v.Index, v.Quality, v.Size)
		}

		var choosenQuality int
		fmt.Print("choose quality: ")
		fmt.Scanln(&choosenQuality)

		downloadLink := getDownloadLinkDirect(getDownloadLink(links[choosenQuality-1].Url))
		fmt.Println(downloadLink)

		// Prompt the user if they want to continue
		var continueOption string
		fmt.Print("Do you want to continue? (yes/no): ")
		fmt.Scanln(&continueOption)
		if continueOption != "yes" {
			break // Exit the loop if the user doesn't want to continue
		}
	}
}