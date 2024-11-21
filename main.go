package main

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"os"
)

const PlantUMLEncoder = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"

func main() {
	baseURL := "https://www.plantuml.com/plantuml"
	// baseURL := "http://localhost:8080"

	text, err := os.ReadFile("./misc/sample.pu")
	if err != nil {
		log.Fatal(err)
	}
	result, err := EncodePlantUML(text)
	if err != nil {
		log.Fatal(err)
	}
	u, err := url.JoinPath(baseURL, "txt", result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
}

func EncodePlantUML(input []byte) (string, error) {
	// 1. compress with Deflate algorithm or Brotli algorithm.
	// TODO: support Brotli algorithm
	var buffer bytes.Buffer
	writer, err := flate.NewWriter(&buffer, flate.DefaultCompression)
	if err != nil {
		return "", err
	}
	if _, err := writer.Write(input); err != nil {
		return "", err
	}
	writer.Close()

	// 2. encode with Base64 like algorithm.
	// Base64 like algorithm: Match 0 ~ 63 value to 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_
	return base64.NewEncoding(PlantUMLEncoder).EncodeToString(buffer.Bytes()), nil
}
