package youtubedlwrapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
)

type WrapperOptions struct {
	YoutubeDLBinary string
	cmd             func(name string, arg ...string) execCommand
}

type Wrapper struct {
	wrapperOptions WrapperOptions
}

//New returns a wrapper with specified WrapperOptions set
func New(wrapperOptions WrapperOptions) (wrapper Wrapper, err error) {
	if _, err = exec.LookPath(wrapperOptions.YoutubeDLBinary); err != nil {
		return
	}
	wrapper.wrapperOptions = wrapperOptions
	wrapper.wrapperOptions.cmd = newStandardCmd
	return wrapper, nil
}

func (wrapper *Wrapper) GetVideoMetadata(url string) (videoMetadata VideoMetadata, err error) {
	cmd := wrapper.wrapperOptions.cmd(wrapper.wrapperOptions.YoutubeDLBinary, "--dump-single-json", url)

	var stdoutBuffer bytes.Buffer
	cmd.SetStdout(&stdoutBuffer)

	var stderrBuffer bytes.Buffer
	cmd.SetStderr(&stderrBuffer)

	switch err = cmd.Run(); err.(type) {
	case *exec.ExitError:
		youtubeDLError, err := io.ReadAll(&stderrBuffer)
		if err != nil {
			return videoMetadata, err
		}

		if youtubeDLErrorString := string(youtubeDLError); youtubeDLErrorString != "" {
			return videoMetadata, YoutubeDLError{error: youtubeDLErrorString}
		}
		break
	case nil:
		break
	default:
		return videoMetadata, err
	}

	youtubeDLOutput := stdoutBuffer.String()
	if len(youtubeDLOutput) == 0 {
		return videoMetadata, ErrNoYoutubeDLOutput
	}

	if err = json.Unmarshal(stdoutBuffer.Bytes(), &videoMetadata); err != nil {
		return videoMetadata, err
	}

	return
}

//GetPlaylistMetadata using wrapperOptions provided in New
func (wrapper *Wrapper) GetPlaylistMetadata(url string) (playlistMetadata PlaylistMetadata, err error) {
	cmd := wrapper.wrapperOptions.cmd(wrapper.wrapperOptions.YoutubeDLBinary, "--dump-single-json", url)

	var stdoutBuffer bytes.Buffer
	cmd.SetStdout(&stdoutBuffer)

	var stderrBuffer bytes.Buffer
	cmd.SetStderr(&stderrBuffer)

	switch err = cmd.Run(); err.(type) {
	case *exec.ExitError:
		youtubeDLError, err := io.ReadAll(&stderrBuffer)
		if err != nil {
			return playlistMetadata, err
		}

		if youtubeDLErrorString := string(youtubeDLError); youtubeDLErrorString != "" {
			return playlistMetadata, YoutubeDLError{error: youtubeDLErrorString}
		}
		break
	case nil:
		break
	default:
		return playlistMetadata, err
	}

	youtubeDLOutput := stdoutBuffer.String()
	if len(youtubeDLOutput) == 0 {
		return playlistMetadata, ErrNoYoutubeDLOutput
	}

	if err = json.Unmarshal(stdoutBuffer.Bytes(), &playlistMetadata); err != nil {
		fmt.Printf("youtubedlresponse: (%v)\n", stdoutBuffer.String())
		return playlistMetadata, err
	}

	return

}
