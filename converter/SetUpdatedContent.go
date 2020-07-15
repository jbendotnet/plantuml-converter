package converter

import (
	"strings"
)

// check whether a end of a block of the file is the current line number
func isLineNumberEndOfBlock(f *PlantUmlFile, lineNumber int) (bool, *PlantUmlBlock) {
	for i := 0; i < len(f.blocks); i++ {
		block := f.blocks[i]
		if block.lineNumber == lineNumber {
			return true, &block
		}
	}
	return false, nil
}

func appendLineContent(f *PlantUmlFile, lineContent string) {
	f.updatedContent = f.updatedContent + lineContent + "\n"
}

func appendMarkdownLink(f *PlantUmlFile, block *PlantUmlBlock) {
	f.updatedContent = f.updatedContent + "![](" + block.markdownLink + ")\n"
}

// adding links and set PlantUmlFile.updatedContent
func (f *PlantUmlFile) SetUpdatedContent() {
	// you can always update the markdown file because the hash will be the same
	// if the content does not change
	lines := strings.Split(f.fileContent, "\n")
	for i := 0; i <= len(lines); i++ {
		isLineNumberEndOfBlock, block := isLineNumberEndOfBlock(f, i)
		if isLineNumberEndOfBlock {
			// insert the markdown
			appendMarkdownLink(f, block)

			isLinePartOfFile := i < len(lines)
			isNotPlantUMLUrl := isLinePartOfFile && strings.HasPrefix(lines[i], "![]("+PlantUmlServerUrl) == false
			if isNotPlantUMLUrl {
				appendLineContent(f, lines[i])
			}
		} else {
			appendLineContent(f, lines[i])
		}
	}

}
