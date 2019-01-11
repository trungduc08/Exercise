package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func Say() {
	fmt.Println("ahihi")
}
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	// q := url.QueryEscape(strings.Join(terms, "windows"))
	resp, err := http.Get(IssuesURL + "?q=windows")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil

}

// func main() {

// 	now := time.Now()
// 	year := now.Year()
// 	month := now.Month()
// 	result, err := SearchIssues(os.Args[1:])
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	categories := map[int][]string{}
// 	j := 0
// 	for i := 0; i < len(result.Items); i++ {
// 		if result.Items[i].CreatedAt.Year()-year > 0 {
// 			categories[j] = []string{
// 				result.Items[i].User.Login,
// 				"more than a year old",
// 			}
// 		} else {
// 			if (result).Items[i].CreatedAt.Month()-month > 0 {
// 				categories[j] = []string{
// 					result.Items[i].User.Login,
// 					"less than a year old",
// 				}
// 			} else {
// 				categories[j] = []string{
// 					result.Items[i].User.Login,
// 					"less than a month old",
// 				}
// 			}
// 		}
// 		j++
// 	}
// 	for i := 0; i < len(categories); i++ {
// 		fmt.Printf("%v\t%9.9s:%s\n", i, categories[i][0], categories[i][1])
// 	}

// }
