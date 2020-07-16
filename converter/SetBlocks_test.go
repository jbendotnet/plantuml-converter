package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SucceededSetBlocks(t *testing.T) {

	inputFilePath := "./mdmocks/testBlockSucceeded.md"
	inputBlocks := []PlantUmlBlock{}

	inputFile := PlantUmlFile{}

	inputFile.filePath = inputFilePath
	inputFile.blocks = inputBlocks

	err := inputFile.SetBlocks()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(inputFile.blocks), 3)
}

func Test_FailedSetBlocks(t *testing.T) {

	inputFilePath := "./mdmocks/testBlockFailed.md"
	inputBlocks := []PlantUmlBlock{}

	inputFile := PlantUmlFile{}

	inputFile.filePath = inputFilePath
	inputFile.blocks = inputBlocks

	err := inputFile.SetBlocks()

	assert.EqualError(t, err, "Failed to parse blocks")
	assert.Equal(t, len(inputFile.blocks), 0)
}

func Test_FailedSetBlocksToLong(t *testing.T) {

	inputFilePath := "./mdmocks/testBlockFailedDueToLong.md"
	inputBlocks := []PlantUmlBlock{}

	inputFile := PlantUmlFile{}

	inputFile.filePath = inputFilePath
	inputFile.blocks = inputBlocks

	err := inputFile.SetBlocks()

	assert.EqualError(t, err, "Failed due to big blocks")

}
