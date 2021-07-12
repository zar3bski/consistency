package media_file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMediaConstructor(t *testing.T) {
	file := NewMediaFile("../../test/data_test/collection1/01 - Some-title.mp3")
	assert.Equal(t, file.Name, "01 - Some-title.mp3", "they should be equal")
}
