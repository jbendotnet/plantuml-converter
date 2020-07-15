package converter

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/signavio/plantuml-converter/cmd"
	"github.com/stretchr/testify/assert"
)

func Test_GenerateLink(t *testing.T) {
	input := "Bob -> Alice : hello"
	output := GenerateLink(input)
	assert.Contains(t, output, cmd.PlantUmlServer)
	fmt.Println(output)
}

func Test_SetFiles(t *testing.T) {
	// remove this once the implementation is done to let the tests run on ci
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	uml := PlantUml{}
	files := []string{"bla.txt", "some.html", "readme.md", "other-readme.md"}
	expected := []PlantUmlFile{
		{
			filePath: "readme.md",
			blocks:   nil,
		},
		{
			filePath: "other-readme.md",
			blocks:   nil,
		},
	}
	for _, file := range files {
		writeFile(file, "")
		defer os.Remove(file)
	}
	uml.SetFiles()
	assert.Equal(t, expected, uml.files)
}

type TestCase struct {
	input    PlantUmlFile
	template string
	expected string
}

func TestPlantUmlFile_Update(t *testing.T) {
	// remove this once the implementation is done to let the tests run on ci
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	exampleBlock := `
@startuml
Bob -> Alice : hello
@enduml
`
	testCases := []TestCase{
		{
			input: PlantUmlFile{
				filePath: "testExampleOneBlock.md",
				blocks: []PlantUmlBlock{
					{
						content: fmt.Sprintf(updateExampleOneBlock(), ""),
					},
				},
			},
			template: updateExampleOneBlock(),
			expected: fmt.Sprintf(updateExampleOneBlock(), GenerateLink(exampleBlock)),
		},
	}
	for _, testCase := range testCases {
		testCase.input.SetUpdatedContent()
		assert.Equal(t, testCase.expected, testCase.input.updatedContent)
	}
}

func writeFile(filename string, content string) {
	data := []byte(content)
	ioutil.WriteFile(filename, data, os.ModePerm)
}

func updateExampleOneBlock() string {
	return `
## heading 1
* adasd
* asdd
<!--
@startuml
Bob -> Alice : hello
@enduml
-->
%s
something else
`
}

func updateExampleTwoBlocks() string {
	return `
## heading 1
* adasd
* asdd
<!--
@startuml
Bob -> Alice : hello
@enduml
-->
%s
something else
<!--
@startuml
Peter -> Gertrud : bye
@enduml
%s
-->
`
}

func updateExampleZeroBlocks() string {
	return `
## heading 1
* adasd
* asdd
<!--
hidden content
-->
`
}

func Test_SucceededSetBlocks(t *testing.T) {

	inputFilePath := "./mdmocks/testBlockSucceeded.md"
	inputBlocks := []PlantUmlBlock{}

	inputFile := PlantUmlFile{}

	inputFile.filePath = inputFilePath
	inputFile.blocks = inputBlocks

	err := inputFile.SetBlocks()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(inputFile.blocks), 2)
}
func Test_FailedSetBlocks(t *testing.T) {

	inputFilePath := "./mdmocks/testBlockFailed.md"
	inputBlocks := []PlantUmlBlock{}

	inputFile := PlantUmlFile{}

	inputFile.filePath = inputFilePath
	inputFile.blocks = inputBlocks

	err := inputFile.SetBlocks()

	assert.EqualError(t, err, "Failed to parse blocks.")
	assert.Equal(t, len(inputFile.blocks), 0)
}

func Test_FailedSetBlocksToLong(t *testing.T) {

	inputFilePath := "./mdmocks/testBlockFailedDueToLong.md"
	inputBlocks := []PlantUmlBlock{}

	inputFile := PlantUmlFile{}

	inputFile.filePath = inputFilePath
	inputFile.blocks = inputBlocks

	err := inputFile.SetBlocks()

	assert.EqualError(t, err, "Failed due to big blocks.")

}

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
