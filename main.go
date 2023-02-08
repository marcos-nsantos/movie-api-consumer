package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var apiKeyFlag = flag.String("key", "", "API Key for IMDB API")

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

	fmt.Println(convertBytesToString(body))
}

func convertBytesToString(data []byte) string {
	return string(data[:])
}
