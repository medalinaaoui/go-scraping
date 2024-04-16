package app

import (
	"strings"
)


func GetWorkType(url string) string {
    if strings.Contains(url, "/series/") {
        return "Serie"
    } else if strings.Contains(url, "/movie/") {
        return "Movie"
    }
    return ""
}
func GetWorkTypeForArabSeed(title string) string {
    if strings.Contains(title, "الحلقة") {
        return "Episode"
    } else {
        return "Movie"
    }
}