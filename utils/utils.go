package utils

import (
	"errors"
	"io/ioutil"
	"os/exec"
)

func RunCmdAndWait(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	resp, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}
	errB, err := ioutil.ReadAll(stderr)
	if err != nil {
		return "", nil
	}

	err = cmd.Wait()
	if err != nil {
		// in case of error, capture the exact message
		if len(errB) > 0 {
			return "", errors.New(string(errB))
		}
		return "", err
	}

	return string(resp), nil
}

func RunCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)

	_, err := cmd.StdoutPipe()
	if err != nil {
		return  err
	}

	err = cmd.Start()
	if err != nil {
		return  err
	}

	return  nil
}
