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

	inputFilePath := "./testBlockSucceeded.md"
	inputBlocks := []PlantUmlBlock{}

	inputFile := PlantUmlFile{}

	inputFile.filePath = inputFilePath
	inputFile.blocks = inputBlocks

	err := inputFile.SetBlocks()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(inputFile.blocks), 2)
}
func Test_FailedSetBlocks(t *testing.T) {

	inputFilePath := "./testBlockFailed.md"
	inputBlocks := []PlantUmlBlock{}

	inputFile := PlantUmlFile{}

	inputFile.filePath = inputFilePath
	inputFile.blocks = inputBlocks

	err := inputFile.SetBlocks()

	assert.EqualError(t, err, "Failed to parse blocks.")
	assert.Equal(t, len(inputFile.blocks), 0)
}
func Test_SetUpdatedContent(t *testing.T) {

	inputFilePath := "./testUpdatedContent.md"
	inputBlocks := []PlantUmlBlock{}

	inputFile := PlantUmlFile{}

	inputFile.filePath = inputFilePath
	inputFile.blocks = inputBlocks

	inputFile.SetBlocks()

	inputFile.SetUpdatedContent()

	// check whether links was pasted into the updatedContent

}
