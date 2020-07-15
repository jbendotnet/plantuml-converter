package converter

import (
	"encoding/hex"
	"fmt"
)

// generates a link to the rendered png image for given input
func GenerateLink(input string) string {
	hash := hex.EncodeToString([]byte(input))
	return fmt.Sprintf("%s/plantuml/png/~h%s", PlantUmlServerUrl, hash)
}

// set the markdown link
func (p *PlantUmlBlock) GenerateMarkdownLink() {
	p.markdownLink = GenerateLink(p.content)
}
