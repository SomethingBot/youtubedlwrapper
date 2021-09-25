package youtubedlwrapper

//makeMockCommand that upon Run or Start outputs parameter strings to respective std io.Writer/io.Reader
//func makeMockCommand(stdin, stdout, stderr string) *exec.Cmd {
//	return &exec.Cmd{}
//}
//
//func TestYoutubeDLError_makeFakeCommand(t *testing.T) {
//	t.Parallel()
//	tests := []struct {
//		stdin  string
//		stdout string
//		stderr string
//	}{
//		{
//			"stdin1",
//			"stdout1",
//			"stderr1",
//		},
//		{
//			"stdin2",
//			"stdout2",
//			"stderr2",
//		},
//	}
//
//	var cmd *exec.Cmd
//	for _, test := range tests {
//		cmd = makeMockCommand(test.stdin, test.stdout, test.stderr)
//		t.Logf("cmd: (%v)", cmd)
//	}
//}
