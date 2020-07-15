package converter

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const Max_Block_Length = 2000 // max length of url in browser

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
	// start line of the block
	startNumber int
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
