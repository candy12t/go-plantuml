package plantuml

import (
	"compress/flate"
	"encoding/base64"
	"io"
)

// PlantUMLEncoding is the PlantUML's own base64 encoding.
// PlantUML's own Base64 match 0 ~ 63 to 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_
var PlantUMLEncoding = base64.NewEncoding("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_")

// Encode encode PlantUML diagram text description to string.
func Encode(dist io.Writer, src io.Reader) error {
	b64w := base64.NewEncoder(PlantUMLEncoding, dist)
	defer b64w.Close()

	fw, err := flate.NewWriter(b64w, flate.DefaultCompression)
	defer fw.Close()
	if err != nil {
		return err
	}

	if _, err := io.Copy(fw, src); err != nil {
		return err
	}
	return nil
}

// Decode decode encoded string to PlantUML diagram text description.
func Decode(dist io.Writer, src io.Reader) error {
	b64r := base64.NewDecoder(PlantUMLEncoding, src)

	frc := flate.NewReader(b64r)
	defer frc.Close()

	if _, err := io.Copy(dist, frc); err != nil {
		return err
	}
	return nil
}
