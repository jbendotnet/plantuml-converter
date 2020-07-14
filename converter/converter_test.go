package converter

import (
	"fmt"
	"github.com/signavio/plantuml-converter/cmd"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
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
		writeEmptyFile(file)
		defer os.Remove(file)
	}
	uml.SetFiles()
	assert.Equal(t, expected, uml.files)
}

func writeEmptyFile(filename string) {
	data := []byte("")
	ioutil.WriteFile(filename, data, os.ModePerm)
}
