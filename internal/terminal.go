package internal

import (
	"os/exec"
	"strings"
)

func ExecCommand(command []string) (string, error) {
	combinedStd, err := exec.Command(command[0], command[1:]...).CombinedOutput()
	s := string(combinedStd)
	sTrimmed := strings.TrimSuffix(s, "\n")
	return sTrimmed, err
}
