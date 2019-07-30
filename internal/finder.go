package internal

import (
	"os/exec"
	"strings"
)

func Find(path string, args []string) ([]string, error) {
	args = append([]string{path}, args...)
	std, err := exec.Command("find", args...).CombinedOutput()
	sTrimmed := strings.TrimSuffix(string(std), "\n")
	arr := strings.Split(sTrimmed, "\n")
	return arr, err
}
