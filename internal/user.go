package internal

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetAllUsers() ([]string, error) {
	users, err := exec.Command("cut", "-d:", "-f1", "/etc/passwd").CombinedOutput()
	usersArr := strings.Split(string(users), "\n")
	return usersArr, err
}

func AddUser(userName string) (string, error) {
	std, err := exec.Command("useradd", userName).CombinedOutput()
	return string(std), err
}

func DeleteUser(userName string) (string, error) {
	std, err := exec.Command("userdel", userName).CombinedOutput()
	return string(std), err
}

func ChangeUserPassword(userName string, password string) (string, error) {
	std, err := exec.Command("bash", "-c", fmt.Sprintf("echo \"%s:%s\" | chpasswd", userName, password)).CombinedOutput()
	return string(std), err
}
