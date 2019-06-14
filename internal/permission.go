package internal

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"
)

func runCommandAsUser(requestedUser, command string) ([]byte, error) {
	u, err := user.Lookup(requestedUser)
	uid, err := strconv.Atoi(u.Uid)
	gid, err := strconv.Atoi(u.Gid)
	cmd := exec.Command(command)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
	combinedStd, err := cmd.CombinedOutput()
	return combinedStd, err
}

func GetOwner(path string) (string, error) {
	combinedStd, err := exec.Command("stat", "-c", "%U", path).CombinedOutput()
	return string(combinedStd), err
}

func ChangeOwnership(path, newOwner string) error {
	u, err := user.Lookup(newOwner)
	uid, err := strconv.Atoi(u.Uid)
	err = os.Chown(path, uid, -1)
	return err
}

func GetCurrentPermissions(path string) (string, error) {
	info, err := os.Stat(path)
	mode := info.Mode()
	return fmt.Sprintf("%o", mode), err
}

func ChangePermissions(userName, path string, permissionNumber uint32) error {
	err := os.Chmod(path, os.FileMode(permissionNumber))
	return err
}
