package file

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	HOSTS_FILE_NAME = "hosts.txt"
)

func GetHostsFileContent() string {
	data, err := os.ReadFile(filepath.Join("data", HOSTS_FILE_NAME))
	if err != nil {
		fmt.Println("GetHostsFileContent", err)
		return ""
	}
	return string(data)
}

func SaveHostsFileContent(buffer []byte) bool {
	err := os.WriteFile(filepath.Join("data", HOSTS_FILE_NAME), buffer, os.ModePerm)
	return err == nil
}

func ReadPrepareFile() string {
	dat, err := os.ReadFile("data/prepare.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(dat))
	return string(dat)
}
