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
	execCmd         func(name string, arg ...string) *exec.Cmd
}

type Wrapper struct {
	wrapperOptions WrapperOptions
}

func New(wrapperOptions WrapperOptions) (wrapper Wrapper, err error) {
	if _, err = exec.LookPath(wrapperOptions.YoutubeDLBinary); err != nil {
		return
	}
	wrapper.wrapperOptions = wrapperOptions
	wrapper.wrapperOptions.execCmd = execCmd
	return wrapper, nil
}

func (wrapper *Wrapper) GetVideoMetadata(url string) (videoMetadata VideoMetadata, err error) {
	cmd := wrapper.wrapperOptions.execCmd(wrapper.wrapperOptions.YoutubeDLBinary, "--dump-single-json", url)

	var stdoutBuffer bytes.Buffer
	cmd.Stdout = &stdoutBuffer

	var stderrBuffer bytes.Buffer
	cmd.Stderr = &stderrBuffer

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
		fmt.Printf("youtubedlresponse: (%v)\n", stdoutBuffer.String())
		return videoMetadata, err
	}

	return
}
