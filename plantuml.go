package main

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"io"
)

// PlantUMLEncoding is the PlantUML's own base64 encoding.
// PlantUML's own Base64 match 0 ~ 63 to 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_
var PlantUMLEncoding = base64.NewEncoding("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_")

// Encode encode PlantUML diagram text description to string.
func Encode(data []byte) (string, error) {
	// 1. compress with Deflate algorithm or Brotli algorithm.
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

	// 2. encode with PlantUML's own Base64.
	return PlantUMLEncoding.EncodeToString(buffer.Bytes()), nil
}

// Decode decode encoded string to PlantUML diagram text description.
func Decode(data string) ([]byte, error) {
	// 1. decode with PlantUML's own base64.
	b, err := PlantUMLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	// 2. decompress with Deflate algorithm or Brotli algorithm.
	// TODO: support Brotli algorithm
	reader := flate.NewReader(bytes.NewReader(b))
	defer reader.Close()

	var buffer bytes.Buffer
	if _, err := io.Copy(&buffer, reader); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
