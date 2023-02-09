package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var apiKeyFlag = flag.String("key", "", "API Key for IMDB API")

type Movie struct {
	Items []struct {
		Title string `json:"title"`
		Image string `json:"image"`
	} `json:"items"`
}

func main() {
	flag.Parse()
	apiKey := *apiKeyFlag

	if apiKey == "" {
		fmt.Println("Please provide an API Key")
		os.Exit(1)
	}
	api := fmt.Sprintf("https://imdb-api.com/en/API/Top250Movies/%s", apiKey)

	resp, err := http.Get(api)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	movies, err := getMoviesTitleAndImage(body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, movie := range movies.Items {
		fmt.Println("Title:", movie.Title)
		fmt.Println("Image:", movie.Image)
		fmt.Println()
	}
}

func getMoviesTitleAndImage(data []byte) (Movie, error) {
	var movies Movie
	if err := json.Unmarshal(data, &movies); err != nil {
		fmt.Println(err)
		return Movie{}, errors.New("unable to unmarshal data")
	}
	return movies, nil
}
