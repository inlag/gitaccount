package main

import (
	"fmt"
	"os/exec"
)

func SetGitConfig(user User) error {
	str := fmt.Sprintf("git config --local user.name \"%s\"", user.Name)
	_, err := execute(str)
	if err != nil {
		return err
	}

	_, err = execute(fmt.Sprintf("git config --local user.email \"%s\"", user.Email))
	if err != nil {
		return err
	}

	return nil
}

func execute(command string) (bool, error) {
	const shell = "/bin/bash"
	_, err := exec.Command(shell, "-c", command).Output()
	if err != nil {
		return false, err
	}
	return true, nil
}
