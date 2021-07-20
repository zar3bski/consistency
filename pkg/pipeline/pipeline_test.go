package pipeline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {
	p := loadFiles("../../test/data_test/collection1")
	out := identifyType(p)
	first_file := <-out
	assert.Equal(t, "01 - Some-title.mp3", first_file.Name, "they should be equal")
}
