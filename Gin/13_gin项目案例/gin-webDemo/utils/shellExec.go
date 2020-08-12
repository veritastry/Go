package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

const (
	shellToUse = "bash"
	quickPush  = `
		git add -A &&
		git commit -m "%v" &&
		git push;
	`
)

//ShellCmd exec a shell command with bash
func ShellCmd(command string) (string, string, error) {
	fmt.Println(command)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(shellToUse, "-c", command)
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
