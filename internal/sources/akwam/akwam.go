package akwam

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/medali/go-scraping/internal/app"
)


type Work struct {
	Source string `json:"source"` 

	Index int    `json:"index"`
	Type string    `json:"type"`
	Name string `json:"name"`
	Url string `json:"url"`
}

func ChoseMovie(s string) []Work{

	query := strings.ReplaceAll(s, " ", "+") 
	url := fmt.Sprintf("https://ak.sv/search?q=%s", query)
	c := colly.NewCollector()

	var movies []Work
	index := 1
	c.OnHTML(".entry-box .entry-image a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		var workType string = app.GetWorkType(link)
		

		movie := Work{
			Source: "akwam",
			Index: index,
			Type: workType,
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




func ChooseQuality(url string) []Link{

	
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


type Episode struct{
	Index int    `json:"index"`
	Eposide string `json:"eposide"`
	Url string `json:"url"` 
}


func ChooseEpisode(url string) []Episode{

	
	c := colly.NewCollector()

	var episodes []Episode
	index := 1

	c.OnHTML("h2 a[href*='ak.sv/episode/']", func(e *colly.HTMLElement) {
		url := e.Attr("href")

episode := Episode{
			Index: index,
			Eposide: e.Text,
			Url: url,
		}
		index++

		episodes = append(episodes, episode)
	})	
	
	c.Visit(url)
	return episodes
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

	type directLink struct{
		Url string `json:"url"`
	}

func GetDownloadLinkDirect(url string) directLink{

	mainUrl := getDownloadLink(url)

	c := colly.NewCollector()

	var link directLink
	c.OnHTML(".btn-loader", func(e *colly.HTMLElement) {
		link.Url = e.ChildAttr("a", "href")
	})	
	
	c.Visit(mainUrl)
	return link
}








func Akwam () {
		fmt.Println("Enter the work name:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		movie := scanner.Text()
		movies := ChoseMovie(movie)

		for _, v := range movies {
			fmt.Printf("\n%v ===> %v ===> %v\n", v.Index, v.Name, v.Type)
		}




		var choosenMovie int
		fmt.Print("Choose work: ")
		fmt.Scanln(&choosenMovie)


		if movies[choosenMovie-1].Type == "Movie"{
			
					links := ChooseQuality(movies[choosenMovie-1].Url)
			
					for _, v := range links {
						fmt.Printf("\n%v ===> %v ===> %v \n", v.Index, v.Quality, v.Size)
					}


					

		var choosenQuality int
		fmt.Print("choose quality: ")
		fmt.Scanln(&choosenQuality)

		downloadLink := GetDownloadLinkDirect(getDownloadLink(links[choosenQuality-1].Url))
		fmt.Println(downloadLink)



		} else{
			episodes := ChooseEpisode(movies[choosenMovie-1].Url)
			
					for _, v := range episodes {
						fmt.Printf("\n%v ===> %v\n", v.Index, v.Eposide)
					}
			
		var choosenEpisode int
		fmt.Print("Choose Episode: ")
		fmt.Scanln(&choosenEpisode)


					links := ChooseQuality(episodes[choosenEpisode-1].Url)
			
					for _, v := range links {
						fmt.Printf("\n%v ===> %v ===> %v \n", v.Index, v.Quality, v.Size)
					}

		
		var choosenQuality int
		fmt.Print("choose quality: ")
		fmt.Scanln(&choosenQuality)

		downloadLink := GetDownloadLinkDirect(getDownloadLink(links[choosenQuality-1].Url))
		fmt.Println(downloadLink)


		}
}