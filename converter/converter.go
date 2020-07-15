package converter

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const Max_Block_Length = 2000 // max length of url in browser

type PlantUml struct {
	files         []PlantUmlFile
	scanDirectory string
	pattern       string
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

func (p *PlantUml) GetPlantFileByPath(path string) *PlantUmlFile {
	for _, plantFile := range p.files {
		if plantFile.filePath == path {
			return &plantFile
		}
	}
	return nil
}

// filter files list for given pattern (cmd.FilePattern) that should be converted
// also reads the file content into PlantUmlFile.fileContent
// it's possible also to apply the recursive with e.g. **/*.md
func (p *PlantUml) SetFiles() {
	var files []PlantUmlFile
	pattern := fmt.Sprintf("%s%c%s", p.scanDirectory, os.PathSeparator, p.pattern)
	fmt.Println(pattern)
	matches, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatal(err)
	}
	for _, match := range matches {
		fmt.Println("found match " + match)
		data, err := ioutil.ReadFile(match)
		if err != nil {
			fmt.Printf("Could not read file %s\n", match)
		} else {
			plantFile := PlantUmlFile{
				filePath:    match,
				fileContent: string(data),
			}
			files = append(files, plantFile)
		}
	}
	p.files = files
}
