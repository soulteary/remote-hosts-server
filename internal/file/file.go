package file

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	HOSTS_FILE_NAME         = "hosts.txt"
	PREPARE_HOSTS_FILE_NAME = "prepare.txt"
)

func GetHostsFileContent(mode string) string {
	fileName := ""
	if mode == "stable" {
		fileName = filepath.Join("data", HOSTS_FILE_NAME)
	} else {
		fileName = filepath.Join("data", PREPARE_HOSTS_FILE_NAME)

	}
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("GetHostsFileContent", err)
		return ""
	}
	return string(data)
}

func SaveHostsFileContent(mode string, buffer []byte) bool {
	// TODO: 判断输入内容有效性
	fileName := ""
	if mode == "stable" {
		fileName = filepath.Join("data", HOSTS_FILE_NAME)
	} else {
		fileName = filepath.Join("data", PREPARE_HOSTS_FILE_NAME)
	}
	err := os.WriteFile(fileName, buffer, os.ModePerm)
	return err == nil
}
