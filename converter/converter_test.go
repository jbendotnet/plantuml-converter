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

}
