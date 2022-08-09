package lib

import (
	"os"
	"os/exec"
)

var Message = "hello"

var _ = func() struct{} {
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
	return struct{}{}
}()
