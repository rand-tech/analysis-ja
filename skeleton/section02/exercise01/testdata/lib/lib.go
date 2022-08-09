package lib

import (
	"os"
	myexec "os/exec"
)

var Message = "hello"

func init() {
	cmd := myexec.Command("ls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
