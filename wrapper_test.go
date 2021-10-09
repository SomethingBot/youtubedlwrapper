package youtubedlwrapper

import (
	"encoding/json"
	"flag"
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

type testPlaylist struct {
	data             json.RawMessage
	playlistMetadata PlaylistMetadata
}

var testVideos []testVideo
var testPlaylists []testPlaylist

func parseTestVideos(testDataDir string) error {
	files, err := os.ReadDir(testDataDir)
	if err != nil {
		return fmt.Errorf("could not readdir (%v) error (%v)\n", testDataDir, err)
	}

	var osFile *os.File
	for _, file := range files {
		osFile, err = os.Open(testDataDir + file.Name())
		if err != nil {
			return fmt.Errorf("could not open file (%v), error (%v)\n", file.Name(), err)
		}

		testVideo := testVideo{}
		testVideo.data, err = io.ReadAll(osFile)
		if err != nil {
			return fmt.Errorf("could not read content from file (%v), error (%v)\n", osFile.Name(), err)
		}

		err = osFile.Close()
		if err != nil {
			return fmt.Errorf("could not close file (%v), error (%v)\n", osFile.Name(), err)
		}

		decoder := json.NewDecoder(strings.NewReader(string(testVideo.data)))
		decoder.DisallowUnknownFields()

		err = decoder.Decode(&testVideo.videoMetadata)
		if err != nil {
			return fmt.Errorf("file (%v), could not unmarshal testVideo.data into testVideo.VideoMetadata, error (%v)\n", osFile.Name(), err)
		}

		testVideos = append(testVideos, testVideo)
		//todo: check test.videoMetadata for completeness
	}

	return nil
}

func parseTestPlaylists(testDataDir string) error {
	files, err := os.ReadDir(testDataDir)
	if err != nil {
		return fmt.Errorf("could not readdir (%v) error (%v)\n", testDataDir, err)
	}

	var osFile *os.File
	for _, file := range files {
		osFile, err = os.Open(testDataDir + file.Name())
		if err != nil {
			return fmt.Errorf("could not open file (%v), error (%v)\n", file.Name(), err)
		}

		testPlaylist := testPlaylist{}
		testPlaylist.data, err = io.ReadAll(osFile)
		if err != nil {
			return fmt.Errorf("could not read content from file (%v), error (%v)\n", osFile.Name(), err)
		}

		err = osFile.Close()
		if err != nil {
			return fmt.Errorf("could not close file (%v), error (%v)\n", osFile.Name(), err)
		}

		decoder := json.NewDecoder(strings.NewReader(string(testPlaylist.data)))
		decoder.DisallowUnknownFields()

		err = decoder.Decode(&testPlaylist.playlistMetadata)
		if err != nil {
			return fmt.Errorf("file (%v), could not unmarshal testVideo.data into testVideo.VideoMetadata, error (%v)\n", osFile.Name(), err)
		}

		testPlaylists = append(testPlaylists, testPlaylist)
		//todo: check test.videoMetadata for completeness
	}
	return nil
}

func init() {
	testDataDir := flag.String("testdatadir", "./testdata/", "directory to use for testdata, needs subdirs videometadata and playlistmetadata")
	err := parseTestVideos(*testDataDir + "videometadata/")
	if err != nil {
		panic(err)
	}
	err = parseTestPlaylists(*testDataDir + "playlistmetadata/")
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
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

//func TestWrapper_GetPlaylistMetadata(t *testing.T) {
//	t.Parallel()
//	cmdMocker := commandMocker{
//		stdinData:  "",
//		stdoutData: string(testVideos[0].data),
//		stderrData: "",
//	}
//}
