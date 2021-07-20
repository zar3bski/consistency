package pipeline

import (
	"os"
	"path/filepath"

	"github.com/zar3bski/consistency/pkg/media_file"
)

func loadFiles(root_path string) <-chan media_file.MediaFile {
	out := make(chan media_file.MediaFile)
	filepath.Walk(root_path, func(path string, info os.FileInfo, err error) error {
		if err != nil && info.IsDir() == false {
			out <- media_file.NewMediaFile(path)
		}
		return nil
	})
	close(out)
	return out
}

func identifyType(file <-chan media_file.MediaFile) <-chan media_file.MediaFile {
	out := make(chan media_file.MediaFile)
	close(out)
	return out
}
