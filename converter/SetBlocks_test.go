package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SucceededSetBlocks(t *testing.T) {

	inputFile := PlantUmlFile{
		filePath: "./mdmocks/testBlockSucceeded.md",
		blocks:   []PlantUmlBlock{},
	}

	err := inputFile.SetBlocks()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(inputFile.blocks), 3)
}

func Test_FailedSetBlocks(t *testing.T) {

	inputFile := PlantUmlFile{
		filePath: "./mdmocks/testBlockFailed.md",
		blocks:   []PlantUmlBlock{},
	}

	err := inputFile.SetBlocks()

	assert.EqualError(t, err, "Failed to parse blocks")
	assert.Equal(t, len(inputFile.blocks), 0)
}

func Test_FailedSetBlocksToLong(t *testing.T) {

	inputFile := PlantUmlFile{
		filePath: "./mdmocks/testBlockFailedDueToLong.md",
		blocks:   []PlantUmlBlock{},
	}

	err := inputFile.SetBlocks()

	assert.EqualError(t, err, "Failed due to big blocks")

}
