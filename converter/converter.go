package converter

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/signavio/plantuml-converter/cmd"
)

type PlantUml struct {
	files         []PlantUmlFile
	scanDirectory string
}

type PlantUmlFile struct {
	filePath       string
	fileContent    string
	updatedContent string
	blocks         []PlantUmlBlock
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

func (p *PlantUml) Length() int {
	return len(p.files)
}

// filter files list for given pattern (cmd.FilePattern) that should be converted
// also reads the file content into PlantUmlFile.fileContent
func (p *PlantUml) SetFiles() {
	var files []PlantUmlFile
	// find files matching the pattern
	// set fileContent and filePath for each PlantUmlFile
	err := filepath.Walk(p.scanDirectory, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		if strings.HasSuffix(path, ".md") {
			fmt.Printf("adding %s\n", path)
			plantFile := PlantUmlFile{
				filePath:    path,
				fileContent: "",
			}
			files = append(files, plantFile)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	p.files = files
}

// adding links and set PlantUmlFile.updatedContent
func (f *PlantUmlFile) SetUpdatedContent() {
	// you can always update the markdown file because the hash will be the same
	// if the content does not change
}

// parse the plant uml blocks from a file
func (f *PlantUmlFile) SetBlocks() error {
	var blocks []PlantUmlBlock

	//TODO Has to be removed if f.fileContent is available
	if len(f.fileContent) == 0 {
		// read lines from file
		bytesRead, _ := ioutil.ReadFile(f.filePath)
		fileContent := string(bytesRead)
		f.fileContent = fileContent
	}

	lines := strings.Split(f.fileContent, "\n")

	var hasStart bool = false

	var myBlock PlantUmlBlock

	for i := 0; i < len(lines); i++ {
		vline := lines[i]
		if strings.TrimSpace(vline) == "@startuml" {
			hasStart = true
		} else if strings.TrimSpace(vline) == "@enduml" {
			if !hasStart {
				return errors.New("Failed to parse blocks.")
			}
			myBlock.lineNumber = i + 1
			blocks = append(blocks, myBlock)
			myBlock = PlantUmlBlock{}
			hasStart = false
		} else if hasStart {
			myBlock.content = myBlock.content + vline
		}
	}
	f.blocks = blocks
	return nil
}
