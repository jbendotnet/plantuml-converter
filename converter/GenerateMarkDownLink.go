package converter

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
)

const base64_alpha = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"

// generates a link to the rendered png image for given input
func GenerateLink(input string) string {
	hash := Encode([]byte(input))
	return fmt.Sprintf("%s/png/%s", PlantUmlServerUrl, hash)
}

func Encode(raw []byte) string {
	compressed := deflate(raw)
	return base64_custom(compressed)
}

func base64_custom(input []byte) string {
	alpha := base64.NewEncoding(base64_alpha)
	writer := bytes.NewBufferString("")
	encoder := base64.NewEncoder(alpha, writer)
	encoder.Write(input)
	encoder.Close()
	return writer.String()
}

// see https://github.com/yogendra/plantuml-go/blob/master/plantuml-go.go
func deflate(input []byte) []byte {
	var b bytes.Buffer
	w, _ := zlib.NewWriterLevel(&b, zlib.BestCompression)
	w.Write(input)
	w.Close()
	return b.Bytes()
}

// set the markdown link
func (p *PlantUmlBlock) GenerateMarkdownLink() {
	p.markdownLink = GenerateLink(p.content)
}
