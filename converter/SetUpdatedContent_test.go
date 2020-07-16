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
	assert.Contains(t, inputFile.updatedContent, PlantUmlServerUrl+"/png/UDhYil98pSd9LoZFByf9CJGojkQo2id8B5R8B5R8prD8IKtBp4jDKSZFuofEBKijIinHqDBAoSnBBTRI0gjCLKZDIr4ulEtmaiDx56pi5BW00G00__yY7Hii")
	assert.Contains(t, inputFile.updatedContent, PlantUmlServerUrl+"/png/UDfCop6kmp080D2TKp0w7E7_eMClqGNPf46Ys3KqadlwhjdVWqYU2Q0xbllaSPpwHVL8DfDssp0FSInBOeE_CNkL26IpXpVbv-HN_8DO21Rl5z7T5X25l3_YLOYGlI4HrzHJSn7XvNlJ7I2o3u1V0000__yPgZ3B")
	assert.Contains(t, inputFile.updatedContent, PlantUmlServerUrl+"/png/UDhYil98pSd9LoZFByf9iUQo2id8B5R8B5R8prD8IKtBp4jDKSZFuofEBKijIinHqDBAoSnBBTRIikO21000__yaAHAX")
	// check also whether old linkds were replaced
	assert.NotContains(t, inputFile.updatedContent, PlantUmlServerUrl+"/png/oldurl")
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
	assert.Equal(t, inputFile.updatedContent, "@startuml\n:Hello world;\n@enduml\n![]("+PlantUmlServerUrl+"/png/UDgoyaZDoSbNACyloacnvW84003__oJI1Bm=)")
}

func Test_SucceededShortSetBlocksWithSetUpdatedContentMultipleTimes(t *testing.T) {

	inputFilePath := "./mdmocks/testBlockSucceededSingleUntouched.md"
	inputBlocks := []PlantUmlBlock{}

	inputFile := PlantUmlFile{}

	inputFile.filePath = inputFilePath
	inputFile.blocks = inputBlocks

	err := inputFile.SetBlocks()
	inputFile.SetUpdatedContent()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(inputFile.blocks), 1)
	assert.Equal(t, inputFile.updatedContent, inputFile.fileContent)
}
