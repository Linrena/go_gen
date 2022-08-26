package util

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func IsDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return fi.IsDir()
}

func IsFileExists(filePath string) bool {
	fi, err := os.Stat(filePath)
	if err != nil {
		return os.IsExist(err)
	}
	return !fi.IsDir()
}

func CreateDir(filePath string) error {
	return os.MkdirAll(filePath, os.ModePerm)
}

func GetModuleNameByPath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return filepath.Base(absPath), nil
}

func WriteFile(filePath string, content []byte, override bool) error {
	f, _ := os.Stat(filePath)
	if override || f == nil {
		return ioutil.WriteFile(filePath, content, os.ModePerm) // ignore_security_alert
	}
	return nil
}

func GoImportFile(filePath string) error {
	cmd := exec.Command("goimports", "-w", filePath)
	if err := cmd.Run(); err != nil {
		return errors.New("goimports error: " + err.Error())
	}
	return nil
}
