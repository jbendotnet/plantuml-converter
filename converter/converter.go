package converter

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/signavio/plantuml-converter/cmd"
)

const Max_Block_Length = 2000 // max length of url in browser

type PlantUml struct {
	files []PlantUmlFile
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

// generates a link to the rendered png image for given input
func GenerateLink(input string) string {
	hash := hex.EncodeToString([]byte(input))
	return fmt.Sprintf("%s/plantuml/png/~h%s", cmd.PlantUmlServer, hash)
}

// set the markdown link
func (p *PlantUmlBlock) GenerateMarkdownLink() {
	p.markdownLink = GenerateLink(p.content)
}

// filter files list for given pattern (cmd.FilePattern) that should be converted
// also reads the file content into PlantUmlFile.fileContent
func (p *PlantUml) SetFiles() {
	var files []PlantUmlFile
	// find files matching the pattern
	// set fileContent and filePath for each PlantUmlFile
	p.files = files
}

// check whether a end of a block of the file is the current line number
func IsLineNUmberAnEndOfAPlantUmlBlock(f *PlantUmlFile, lineNumber int) (bool, *PlantUmlBlock) {
	for i := 0; i < len(f.blocks); i++ {
		block := f.blocks[i]
		if block.lineNumber == lineNumber {
			return true, &block
		}
	}
	return false, nil
}

// adding links and set PlantUmlFile.updatedContent
func (f *PlantUmlFile) SetUpdatedContent() {
	// you can always update the markdown file because the hash will be the same
	// if the content does not change
	lines := strings.Split(f.fileContent, "\n")
	for i := 0; i < len(lines); i++ {
		isLineOfBlock, block := IsLineNUmberAnEndOfAPlantUmlBlock(f, i)
		if isLineOfBlock {
			// insert the markdown
			f.updatedContent = f.updatedContent + "![](" + block.markdownLink + ")\n"

			if strings.HasPrefix(lines[i], "![]("+cmd.PlantUmlServer) == false {
				f.updatedContent = f.updatedContent + lines[i] + "\n"
			}
		} else {
			f.updatedContent = f.updatedContent + lines[i] + "\n"
		}

	}
	fmt.Println(f.updatedContent)

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
	var blocksize int = 0

	var myBlock PlantUmlBlock

	for i := 0; i < len(lines); i++ {

		if blocksize > Max_Block_Length {
			return errors.New("Failed due to big blocks.")
		}

		vline := lines[i]

		if strings.Contains(vline, "@startuml") {
			hasStart = true
		} else if strings.Contains(vline, "@enduml") {
			if !hasStart {
				return errors.New("Failed to parse blocks.")
			}
			myBlock.lineNumber = i + 1
			myBlock.GenerateMarkdownLink()

			blocks = append(blocks, myBlock)
			myBlock = PlantUmlBlock{}
			hasStart = false
			blocksize = 0
		} else if hasStart {
			myBlock.content = myBlock.content + vline + "\n"
			blocksize = blocksize + len(vline)
		}
	}
	f.blocks = blocks
	return nil
}
