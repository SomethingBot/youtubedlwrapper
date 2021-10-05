package youtubedlwrapper

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

type mockCommand struct {
	complete   bool
	path       string
	args       []string
	stdin      io.Reader
	stdinData  string
	stdout     io.Writer
	stdoutData string
	stderr     io.Writer
	stderrData string
	startErr   error
	waitErr    error
	runErr     error
}

func (mockCommand *mockCommand) Start() error {
	return mockCommand.startErr
}

func (mockCommand *mockCommand) Run() error {
	var err error
	if mockCommand.stdin != nil {
		_, err = mockCommand.stdin.Read([]byte(mockCommand.stdinData))
		if err != nil {
			return err
		}
	}

	_, err = mockCommand.stdout.Write([]byte(mockCommand.stdoutData))
	if err != nil {
		return err
	}
	_, err = mockCommand.stderr.Write([]byte(mockCommand.stderrData))
	if err != nil {
		return err
	}
	return mockCommand.runErr
}

func (mockCommand *mockCommand) Wait() error {
	return mockCommand.waitErr
}

func (mockCommand *mockCommand) Output() ([]byte, error) {
	panic("implement me")
}

func (mockCommand *mockCommand) CombinedOutput() ([]byte, error) {
	panic("implement me")
}

func (mockCommand *mockCommand) StdinPipe() (io.WriteCloser, error) {
	panic("implement me")
}

func (mockCommand *mockCommand) StdoutPipe() (io.ReadCloser, error) {
	panic("implement me")
}

func (mockCommand *mockCommand) StderrPipe() (io.ReadCloser, error) {
	panic("implement me")
}

func (mockCommand *mockCommand) SetStdin(stdin io.Reader) {
	mockCommand.stdin = stdin
}

func (mockCommand *mockCommand) SetStdout(stdout io.Writer) {
	mockCommand.stdout = stdout
}

func (mockCommand *mockCommand) SetStderr(stderr io.Writer) {
	mockCommand.stderr = stderr
}

func (mockCommand *mockCommand) String() string {
	builder := new(strings.Builder)
	builder.WriteString(mockCommand.path)
	return builder.String()
}

func TestYoutubeDLError_mockCommand_Run(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		cmd           mockCommand
		expectedErr   error
		runBeforeFunc func(mockCommand *mockCommand)
	}{
		{
			"std1",
			mockCommand{
				stdinData:  "stdin1",
				stdoutData: "stdout1",
				stderrData: "stderr1",
				startErr:   nil,
				waitErr:    nil,
				runErr:     nil,
			},
			nil,
			func(mockCommand *mockCommand) {

			},
		},
	}

	var err error
	for _, test := range tests {
		if test.runBeforeFunc != nil {
			test.runBeforeFunc(&test.cmd)
		}

		var stdoutBuffer, stderrBuffer bytes.Buffer
		test.cmd.SetStdout(&stdoutBuffer)
		test.cmd.SetStderr(&stderrBuffer)

		err = test.cmd.Run()
		if err != test.expectedErr {
			t.Errorf("test (%v), expected error (%v), got error (%v)\n", test.name, test.expectedErr, err)
		}

		stdout, err := io.ReadAll(&stdoutBuffer)
		if err != nil {
			t.Errorf("test (%v), got error (%v) while reading stdoutBuffer\n", test.name, err)
		}
		if stdoutString := string(stdout); stdoutString != test.cmd.stdoutData {
			t.Errorf("test (%v), cmd.stdout Reader returned stdout (%v), expected (%v)", test.name, stdoutString, test.cmd.stdoutData)
		}

		stderr, err := io.ReadAll(&stderrBuffer)
		if err != nil {
			t.Errorf("test (%v), got error (%v) while reading stderrBuffer\n", test.name, err)
		}
		if stderrString := string(stderr); stderrString != test.cmd.stderrData {
			t.Errorf("test (%v), cmd.stderr Reader returned stderr (%v), expected (%v)", test.name, stderrString, test.cmd.stderrData)
		}
	}
}

type commandMocker struct {
	stdinData  string
	stdoutData string
	stderrData string
	startErr   error
	waitErr    error
	runErr     error
}

func (commandMocker *commandMocker) makeMockCommand(name string, arg ...string) Cmd {
	cmd := &mockCommand{
		path:       name,
		args:       arg,
		stdinData:  commandMocker.stdinData,
		stdoutData: commandMocker.stdoutData,
		stderrData: commandMocker.stderrData,
		startErr:   commandMocker.startErr,
		waitErr:    commandMocker.waitErr,
		runErr:     commandMocker.runErr,
	}
	return cmd
}
