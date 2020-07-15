package converter

import (
	"strings"
)

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

			if strings.HasPrefix(lines[i], "![]("+PlantUmlServerUrl) == false {
				f.updatedContent = f.updatedContent + lines[i] + "\n"
			}
		} else {
			f.updatedContent = f.updatedContent + lines[i] + "\n"
		}

	}

}
