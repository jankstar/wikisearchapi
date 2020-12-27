package wikisearchapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	urlWiki string = "https://%s.wikipedia.org/w/api.php" //Attention: Language as placeholder
)

//SearchWikiResponse is the found set of the search
type SearchWikiResponse struct {
	Batchcomplete string                 `json:"batchcomplete"`
	Warnings      map[string]interface{} `json:"warnings"` //tyWarnings `json: "warnings"`
	Continue      struct {
		Sroffset int    `json:"sroffset"`
		Continue string `json:"continue"`
	} `json:"continue"`
	Query struct {
		Searchinfo struct {
			Ns        int    `json:"ns"`
			Title     string `json:"title"`
			Pageid    int    `json:"pageid"`
			Size      int    `json:"size"`
			Wordcount int    `json:"wordcount"`
			Snippet   string `json:"snippet"`
			Timestamp string `json:"timestamp"`
			Totalhits int    `json:"totalhits"`
		} `json:"searchinfo"`
		Search []struct {
			Ns        int    `json:"ns"`
			Title     string `json:"title"`
			Pageid    int    `json:"pageid"`
			Size      int    `json:"size"`
			Wordcount int    `json:"wordcount"`
			Snippet   string `json:"snippet"`
			Timestamp string `json:"timestamp"`
		} `json:"search"`
	} `json:"query"`
}

//GetSearchWiki searches the query in the WIKI in the language langu
func GetSearchWiki(langu string, query string) (SearchWikiResponse, error) {
	var dst SearchWikiResponse

	if langu == "" {
		langu = "de"
	}
	myURL := fmt.Sprintf(urlWiki, langu[0:2])
	myURL = myURL + "?action=query&format=json&list=search&srlimit=1&srsearch=" + url.QueryEscape(query)
	res, err1 := http.Get(myURL)
	if err1 != nil || res.StatusCode != 200 {
		return dst, err1
	}
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	//dec.DisallowUnknownFields() //Ignore unknown fields

	err2 := dec.Decode(&dst)
	if err2 != nil || len(dst.Query.Search) == 0 {
		return dst, err2
	}

	return dst, nil
}
