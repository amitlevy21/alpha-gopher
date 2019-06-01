package internal

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetAllUsers() ([]string, error) {
	users, err := exec.Command("cut", "-d:", "-f1", "/etc/passwd").Output()
	usersArr := strings.Split(string(users), "\n")
	return usersArr, err
}

func AddUser(userName string) error {
	err := exec.Command("useradd", userName).Run()
	return err
}

func DeleteUser(userName string) error {
	err := exec.Command("userdel", userName).Run()
	return err
}

func ChangeUserPassword(userName string, password string) error {
	err := exec.Command("bash", "-c", fmt.Sprintf("echo \"%s:%s\" | chpasswd", userName, password)).Run()
	return err
}
