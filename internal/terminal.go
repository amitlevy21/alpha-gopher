package internal

import (
	"os"
	"os/exec"
	"strings"
)

var currentDir = os.Getenv("PROJECT_PATH")

func ExecCommand(command []string) (string, error) {
	if command[0] == "cd" {
		if len(command) > 1 {
			currentDir = command[1]
		} else {
			currentDir = "~"
		}
		return currentDir, nil
	}
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Dir = currentDir
	combinedStd, err := cmd.CombinedOutput()
	s := string(combinedStd)
	sTrimmed := strings.TrimSuffix(s, "\n")
	return sTrimmed, err
}
