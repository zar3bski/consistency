package media_file

import (
	"log"
	"os"
	"path/filepath"
)

type MediaFile struct {
	Name          string
	Path          string
	FileExtension string
	Size          int64
}

func NewMediaFile(path string) MediaFile {
	extension := filepath.Ext(path)
	fi, err := os.Stat(path)
	if err != nil {
		log.Fatalf("File %s could not be analyzed: %v\n", path, err)
	}
	return MediaFile{Name: fi.Name(), Path: path, FileExtension: extension, Size: fi.Size()}
}

type MusicFile struct {
	MediaFile
	Artist       string
	Album        string
	Year         int8
	TrackNb      int8
	Codec        string
	SamplingRate int16
}
