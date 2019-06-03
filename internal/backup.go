package internal

import (
	"os/exec"
	"strings"
)

type Device struct {
	name       string
	size       string
	mountPoint string
}

const maxDevices = 20

func GetDevices() ([]Device, error) {
	lsblkCmd := "lsblk -o name,size,mountpoint -n -l /dev/sd* | uniq"
	combinedStd, err := exec.Command("bash", "-c", lsblkCmd).CombinedOutput()
	devices := make([]Device, maxDevices)
	for device := range combinedStd {
		dev_spec := strings.Split(string(device), "")
		dev := Device{name: dev_spec[0], size: dev_spec[1], mountPoint: dev_spec[2]}
		devices = append(devices, dev)
	}

	return devices, err
}
