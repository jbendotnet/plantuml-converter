package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SetUpdatedContent(t *testing.T) {

	inputFilePath := "./mdmocks/testUpdatedContent.md"
	inputBlocks := []PlantUmlBlock{}

	inputFile := PlantUmlFile{}

	inputFile.filePath = inputFilePath
	inputFile.blocks = inputBlocks

	inputFile.SetBlocks()

	inputFile.SetUpdatedContent()

	// check whether links was pasted into the updatedContent
	assert.Contains(t, inputFile.updatedContent, "http://www.plantuml.com/plantuml/png/~h0a3a48656c6c6f20776f726c643b0a3a54686973206973206f6e20646566696e6564206f6e0a7365766572616c202a2a6c696e65732a2a3b0a")
	assert.Contains(t, inputFile.updatedContent, "http://www.plantuml.com/plantuml/png/~h0a73746172740a0a69662028477261706876697a20696e7374616c6c65643f29207468656e2028796573290a3a70726f6365737320616c6c5c6e6469616772616d733b0a656c736520286e6f290a3a70726f63657373206f6e6c790a2a2a73657175656e63652a2a20616e64202a2a61637469766974792a2a206469616772616d733b0a656e6469660a0a73746f700a0a")
	// check also whether old linkds were replaced
	assert.NotContains(t, inputFile.updatedContent, "http://www.plantuml.com/plantuml/png/oldurl")
}

func Test_SucceededShortSetBlocksWithSetUpdatedContent(t *testing.T) {

	inputFilePath := "./mdmocks/testBlockSucceededSingle.md"
	inputBlocks := []PlantUmlBlock{}

	inputFile := PlantUmlFile{}

	inputFile.filePath = inputFilePath
	inputFile.blocks = inputBlocks

	err := inputFile.SetBlocks()
	inputFile.SetUpdatedContent()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(inputFile.blocks), 1)
	assert.Equal(t, inputFile.updatedContent, "@startuml\n:Hello world;\n@enduml\n![](http://www.plantuml.com/plantuml/png/~h3a48656c6c6f20776f726c643b0a)\n")
}
