package internal

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

type Device struct {
	Name       string `json:"name"`
	Size       string `json:"size"`
	MountPoint string `json:"mountPoint"`
}

func GetDevices() ([]*Device, error) {
	lsblkCmd := "lsblk -o name,size,mountpoint -n -l | uniq"
	combinedStd, err := exec.Command("bash", "-c", lsblkCmd).CombinedOutput()
	devices := []*Device{}
	scanner := bufio.NewScanner(strings.NewReader(string(combinedStd)))
	for scanner.Scan() {
		dev_spec := strings.Fields(scanner.Text())
		mountPoint := ""
		if len(dev_spec) > 2 {
			mountPoint = dev_spec[2]
		}
		dev := &Device{Name: dev_spec[0], Size: dev_spec[1], MountPoint: mountPoint}
		devices = append(devices, dev)
	}

	return devices, err
}

func Backup(path string) (string, error) {
	source := path
	destionation := "/backup"
	rsync := "rsync -aAXv --delete --exclude=/dev/* --exclude=/proc/* --exclude=/sys/*" +
		" --exclude=/tmp/* --exclude=/run/* --exclude=/mnt/* --exclude=/media/* --exclude=\"swapfile\"" +
		" --exclude=\"lost+found\" --exclude=\".cache\""
	rsync = fmt.Sprintf("%s %s %s", rsync, source, destionation)
	combinedStd, err := exec.Command("bash", "-c", rsync).CombinedOutput()

	return string(combinedStd), err
}
