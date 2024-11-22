package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {
	baseURL := "https://www.plantuml.com/plantuml"
	// baseURL := "http://localhost:8080"

	text, err := os.ReadFile("./misc/sample.pu")
	if err != nil {
		log.Fatal(err)
	}
	result, err := Encode(text)
	if err != nil {
		log.Fatal(err)
	}
	u, err := url.JoinPath(baseURL, "txt", result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
}
