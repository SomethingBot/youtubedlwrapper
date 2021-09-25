package youtubedlwrapper

import "os/exec"

func execCmd(name string, arg ...string) *exec.Cmd {
	return exec.Command(name, arg...)
}
