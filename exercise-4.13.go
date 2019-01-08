package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Movie struct {
	Title  string `json:"title"`
	Poster string `json:poster`
}

const URLMOVIE = "http://www.omdbapi.com/?i=tt3896198&apikey=77cccf2b"

func SearchMovie() (*Movie, error) {
	resp, err := http.Get(URLMOVIE)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search movie failed: %s", resp.Status)
	}
	var result Movie
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
func main() {
	res, err := SearchMovie()
	if err == nil {
		fmt.Printf("Title: %s\n", res.Title)
		fmt.Printf("Poster: %s\n", res.Poster)
		DownloadPoster(res.Poster, "poster.jpg")
	}
}
func DownloadPoster(url string, name string) {
	img, _ := os.Create(name)
	defer img.Close()
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	b, _ := io.Copy(img, resp.Body)
	fmt.Println("Download success")
	fmt.Println("file size", b)

}
