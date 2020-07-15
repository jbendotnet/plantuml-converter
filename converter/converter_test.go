package converter

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateLink(t *testing.T) {
	input := "Bob -> Alice : hello"
	output := GenerateLink(input)
	assert.Contains(t, output, PlantUmlServerUrl)
	fmt.Println(output)
}

type testFile struct {
	path    string
	content string
}

func Test_SetFiles(t *testing.T) {
	uml := PlantUml{
		ScanDirectory: ".",
		Pattern:       "*.md",
	}
	files := []testFile{
		{
			path:    "readme.md",
			content: "hello world",
		},
		{
			path:    "other-readme.md",
			content: "squirrel",
		},
	}
	expected := []PlantUmlFile{
		{
			filePath:    "other-readme.md",
			fileContent: "squirrel",
			blocks:      nil,
		},
		{
			filePath:    "readme.md",
			fileContent: "hello world",
			blocks:      nil,
		},
	}
	for _, file := range files {
		testFile := fmt.Sprintf("%s%c%s", uml.ScanDirectory, os.PathSeparator, file.path)
		err := writeFile(testFile, file.content)
		assert.NoError(t, err)
		defer os.Remove(file.path)
	}
	uml.SetFiles()
	assert.Equal(t, 2, len(uml.files))
	for _, expectedFile := range expected {
		assert.Equal(t, expectedFile.filePath, uml.GetPlantFileByPath(expectedFile.filePath).filePath)
		assert.Equal(t, expectedFile.fileContent, uml.GetPlantFileByPath(expectedFile.filePath).fileContent)
	}
}

type TestCase struct {
	input              PlantUmlFile
	template, expected string
}

func writeFile(filename string, content string) error {
	data := []byte(content)
	return ioutil.WriteFile(filename, data, os.ModePerm)
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
