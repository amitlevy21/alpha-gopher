package internal

import (
	"fmt"
	"os/exec"
)

func GetAllUsers() (string, error) {
	users, err := exec.Command("bash", "-c", "/go/src/github.com/amitlevy21/alpha-gopher/internal/get_user_data.sh").CombinedOutput()
	return string(users), err
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
