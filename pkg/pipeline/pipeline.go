package pipeline

import "github.com/zar3bski/consistency/pkg/media_file"
import "path/filepath"

func loadFiles(root_path string) <-chan media_file.MediaFile {
	out := make(chan media_file.MediaFile)
	go func() {
		for path := range filepath.Walk(root_path){
			out<- media_file.NewMediaFile(path)
		}
		
	}
	return out
}

func identifyType(file media_file.MediaFile) <-chan media_file.MediaFile {
	out := make(chan media_file.MediaFile)
	return out
}
