package converter

import (
	"encoding/hex"
	"fmt"
	"github.com/signavio/plantuml-converter/cmd"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GenerateLink(t *testing.T) {
	input := "Bob -> Alice : hello"
	output := GenerateLink(input)
	decoded, err := hex.DecodeString(output)
	assert.NoError(t, err)
	assert.Equal(t, input, decoded)
	assert.Contains(t, output, cmd.PlantUmlServer)
	fmt.Println(output)
}
