package pipeline

import "testing"

func TestPipeline(t *testing.T) {
	p := identifyType(loadFile("test/data_test/collection1/01 - Some-title.mp3"))
}
