package media_file

type MediaFile struct {
	Path          string
	FileExtension string
	Size          int16
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
