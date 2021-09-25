package youtubedlwrapper

import "fmt"

var ErrNoYoutubeDLOutput = fmt.Errorf("youtubedlwrapper: youtube-dl returned no output")

//YoutubeDLError is specifically an error returned on stderr by youtube-dl
type YoutubeDLError struct {
	error string
}

func (youtubeDLError YoutubeDLError) Error() string {
	return youtubeDLError.error
}
