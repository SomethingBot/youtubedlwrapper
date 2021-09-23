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

	_, err = youtubeDLWrapper.GetVideoMetadata("https://www.youtube.com/watch?v=lfW5CF0Nsis")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = youtubeDLWrapper.GetVideoMetadata("https://www.youtube.com/watch?v=lfW5CF0NsiBADURL")
	if err != nil {
		return
	}

	t.Error("Did not receive an error from bad url passed to youtube-dl")
}
