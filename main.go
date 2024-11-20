package main

import (
	"bytes"
	"compress/flate"
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
	// 1. UTF-8 encode.
	data := []byte(input)

	// 2. compress with Deflate algorithm or Brotli algorithm.
	// TODO: support Brotli algorithm
	var buffer bytes.Buffer
	writer, err := flate.NewWriter(&buffer, flate.DefaultCompression)
	if err != nil {
		return "", err
	}
	if _, err := writer.Write(data); err != nil {
		return "", err
	}
	writer.Close()

	// 3. encode with Base64 like algorithm.
	// Base64 like algorithm: Match 0 ~ 63 value to 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_
	return encode64(buffer.Bytes()), nil
}

// encode64 encodes the given byte slice into a custom Base64-like string.
func encode64(data []byte) string {
	var result string
	for i := 0; i < len(data); i += 3 {
		if i+2 == len(data) {
			result += append3bytes(int(data[i]), int(data[i+1]), 0)
		} else if i+1 == len(data) {
			result += append3bytes(int(data[i]), 0, 0)
		} else {
			result += append3bytes(int(data[i]), int(data[i+1]), int(data[i+2]))
		}
	}
	return result
}

// append3bytes encodes 3 bytes into 4 characters.
func append3bytes(b1, b2, b3 int) string {
	c1 := b1 >> 2
	c2 := ((b1 & 0x3) << 4) | (b2 >> 4)
	c3 := ((b2 & 0xF) << 2) | (b3 >> 6)
	c4 := b3 & 0x3F

	return string([]rune{
		encode6bit(c1 & 0x3F),
		encode6bit(c2 & 0x3F),
		encode6bit(c3 & 0x3F),
		encode6bit(c4 & 0x3F),
	})
}

// encode6bit encodes a 6-bit value into a character.
func encode6bit(b int) rune {
	if b < 10 {
		return rune('0' + b)
	}
	b -= 10
	if b < 26 {
		return rune('A' + b)
	}
	b -= 26
	if b < 26 {
		return rune('a' + b)
	}
	b -= 26
	if b == 0 {
		return '-'
	}
	if b == 1 {
		return '_'
	}
	return '?'
}
