package converter

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/signavio/plantuml-converter/cmd"
)

type PlantUmlConverter interface {
	GenerateLink() string
	GetFiles() []PlantUmlFile
}

type PlantUml struct {
	files []PlantUmlFile
}

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

// parse the plant uml blocks from a file
func (f *PlantUmlFile) SetBlocks() (string, error) {
	var blocks []PlantUmlBlock
	f.blocks = blocks
	return "", errors.New("Not implemented")
}

// set file list for given pattern (cmd.FilePattern) that should be converted
func (p *PlantUml) SetFiles() {
	var files []PlantUmlFile
	// find files matching the pattern
	p.files = files
}

// update markdown file
func (f *PlantUmlFile) Update() {
	// you can always update the markdown file because the hash will be the same
	// if the content does not change
}
