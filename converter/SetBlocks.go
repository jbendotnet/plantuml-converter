package converter

import (
	"errors"
	"io/ioutil"
	"strings"
)

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
