package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const URLComic = "https://xkcd.com"
const Query = "/info.0.json"

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func GetComic(num int) (*Comic, error) {

	resp, err := http.Get(URLComic + "/" + strconv.Itoa(num) + Query)
	if err != nil {
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result Comic
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil

}
func CountUrl() int { // return count of Urls
	req, _ := http.NewRequest("GET", URLComic+Query, nil)
	client := &http.Client{}

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	var comic Comic
	json.Unmarshal(body, &comic)
	return comic.Num
}
func GetAllComic(count int) []*Comic {
	var lstComic []*Comic
	for i := 1; i <= count; i++ {
		comic, _ := GetComic(i)
		lstComic = append(lstComic, comic)
	}
	return lstComic
}

// func main() {
// 	list := GetAllComic(5)

// 	var index int
// 	fmt.Println("Nhap index")
// 	fmt.Scan(&index)
// 	SearchOffline(index, list)

// }
func SearchOffline(index int, lst []*Comic) {
	for i := 0; i < len(lst); i++ {
		if index == lst[i].Num {
			fmt.Printf("URL: %s/%s%s\n", URLComic, strconv.Itoa(index), Query)
			fmt.Printf("Transcript: %s\n", lst[i].Transcript)
			return
		}
	}
	fmt.Println("Not found")

}
