package internal

import (
	"os/exec"
	"strings"
)

func Copy(src, dst string, opts []string) (string, error) {
	args := []string{}
	if opts != nil {
		args = opts
	}

	args = append(args, src, dst)
	combinedStd, err := exec.Command("cp", args...).CombinedOutput()
	return string(combinedStd), err
}

func Move(src, dst string) (string, error) {
	combinedStd, err := exec.Command("mv", src, dst).CombinedOutput()
	return string(combinedStd), err
}

func List(path string, opts []string) ([]string, error) {
	args := []string{}
	if opts != nil {
		args = opts
	}

	args = append(args, path)
	combinedStd, err := exec.Command("ls", args...).CombinedOutput()
	s := string(combinedStd)
	s = strings.TrimSuffix(s, "\n")
	entries := strings.Split(s, "\n")
	return entries, err
}

func BuildTree() (string, error) {
	combinedStd, err := exec.Command("tree", "-JDug", "/").CombinedOutput()
	return string(combinedStd), err
}
