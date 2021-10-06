package youtubedlwrapper

import (
	"io"
	"os/exec"
)

type execCommand interface {
	Start() error
	Run() error
	Wait() error
	Output() ([]byte, error)
	CombinedOutput() ([]byte, error)
	StdinPipe() (io.WriteCloser, error)
	StdoutPipe() (io.ReadCloser, error)
	StderrPipe() (io.ReadCloser, error)
	SetStdin(stdin io.Reader)
	SetStdout(stdout io.Writer)
	SetStderr(stderr io.Writer)
	String() string
}

type standardCmd struct {
	*exec.Cmd
}

func newStandardCmd(name string, arg ...string) execCommand {
	return &standardCmd{exec.Command(name, arg...)}
}

func (standardCmd *standardCmd) SetStdin(stdin io.Reader) {
	standardCmd.Stdin = stdin
}

func (standardCmd *standardCmd) SetStdout(stdout io.Writer) {
	standardCmd.Stdout = stdout
}

func (standardCmd *standardCmd) SetStderr(stderr io.Writer) {
	standardCmd.Stderr = stderr
}
