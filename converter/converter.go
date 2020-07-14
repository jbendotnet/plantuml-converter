package converter

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/signavio/plantuml-converter/cmd"
)

type PlantUmlFile struct {
	filePath string
	blocks   []PlantUmlBlock
}

type PlantUmlBlock struct {
	// to last line number of content block
	// import to remember where to put the link
	lineNumber int
	// plant uml block
	content string
	// generated markdown link
	markdownLink string
}

// generates a link to the rendered png image for given input
func GenerateLink(input string) string {
	hash := hex.EncodeToString([]byte(input))
	return fmt.Sprintf("%s/plantuml/png/~h%s", cmd.PlantUmlServer, hash)
}

// set the markdown link
func (p *PlantUmlBlock) GenerateMarkdownLink() {
	p.markdownLink = GenerateLink(p.content)
}

// parse a given input and return a list of plantUML blocks
func ParsePlantUml(input string) (string, error) {
	return "", errors.New("Not implemented")
}

// get a file list for given pattern (cmd.FilePattern) that should be converted
func GetFiles() []PlantUmlFile {
	var files []PlantUmlFile
	return files
}

// update markdown file and add markdown link if it does not exist
func (f *PlantUmlFile) Update() {

}
