package youtubedlwrapper

import (
	"testing"
)

func TestYoutubeDLWrapper_GetMetaData(t *testing.T) {
	t.Parallel()
	youtubeDLWrapper, err := New(WrapperOptions{YoutubeDLBinary: "youtube-dl"})
	if err != nil {
		t.Error(err)
		return
	}
	_, err = youtubeDLWrapper.GetVideoMetadata("https://www.youtube.com/watch?v=rFejpH_tAHM")
	if err != nil {
		t.Error(err)
		return
	}
}
