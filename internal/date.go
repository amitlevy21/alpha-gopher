package internal

import (
	"os/exec"
	"time"
)

func GetDate() (string, error) {
	combinedStd, err := exec.Command("date", "+%d-%b-%Y %R").CombinedOutput()
	return string(combinedStd), err
}

func SetDate(newTime time.Time) (string, error) {
	dateString := newTime.Format("2 Jan 2006 15:04:05")
	args := []string{"--set", dateString}
	combinedStd, err := exec.Command("date", args...).CombinedOutput()
	return string(combinedStd), err
}
