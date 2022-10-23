package file

import (
	"fmt"
	"os"
)

func ReadFile() string {
	dat, err := os.ReadFile("data/basic.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(dat))
	return string(dat)
}


func ReadPrepareFile() string {
	dat, err := os.ReadFile("data/prepare.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(dat))
	return string(dat)
}
