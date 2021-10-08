package youtubedlwrapper

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

type testVideo struct {
	data          json.RawMessage
	videoMetadata VideoMetadata
}

var testVideos []testVideo

func init() {
	testDataDir := "./testdata/videometadata/"

	files, err := os.ReadDir(testDataDir)
	if err != nil {
		panic(fmt.Errorf("could not readdir (%v) error (%v)\n", testDataDir, err))
	}

	var osFile *os.File
	for _, file := range files {
		osFile, err = os.Open(testDataDir + file.Name())
		if err != nil {
			panic(fmt.Errorf("could not open file (%v), error (%v)\n", file.Name(), err))
		}

		testVideo := testVideo{}
		testVideo.data, err = io.ReadAll(osFile)
		if err != nil {
			panic(fmt.Errorf("could not read content from file (%v), error (%v)\n", osFile.Name(), err))
		}

		err = osFile.Close()
		if err != nil {
			panic(fmt.Errorf("could not close file (%v), error (%v)\n", osFile.Name(), err))
		}

		decoder := json.NewDecoder(strings.NewReader(string(testVideo.data)))
		decoder.DisallowUnknownFields()

		err = decoder.Decode(&testVideo.videoMetadata)
		if err != nil {
			panic(fmt.Errorf("file (%v), could not unmarshal testVideo.data into testVideo.VideoMetadata, error (%v)\n", osFile.Name(), err))
		}

		testVideos = append(testVideos, testVideo)
		//todo: check test.videoMetadata for completeness
	}
}

func TestYoutubeDLWrapper_GetMetaData(t *testing.T) {
	t.Parallel()

	cmdMocker := commandMocker{
		stdinData:  "",
		stdoutData: string(testVideos[0].data),
		stderrData: "",
	}

	youtubeDLWrapper := Wrapper{wrapperOptions: WrapperOptions{
		YoutubeDLBinary: "",
		cmd:             cmdMocker.makeMockCommand,
	}}

	metadata, err := youtubeDLWrapper.GetVideoMetadata("https://www.youtube.com/watch?v=" + testVideos[0].videoMetadata.ID)
	if err != nil {
		t.Error(err)
		return
	}

	if metadata.ID == "" {
		t.Error("Metadata doesnt exist after error isn't nil")
	}

	cmdMocker.runErr = &exec.ExitError{
		ProcessState: nil,
		Stderr:       []byte("status code 1"),
	}
	cmdMocker.stderrData = "status code 1"
	_, err = youtubeDLWrapper.GetVideoMetadata("https://www.youtube.com/watch?v=" + testVideos[0].videoMetadata.ID + "BADURL")
	if err != nil {
		return
	}

	t.Error("Did not receive an error from bad url passed to youtube-dl")
}
