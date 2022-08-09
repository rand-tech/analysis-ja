package lib

import (
	"os"
	"os/exec"
)

var Message = "hello"

func init() {
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
