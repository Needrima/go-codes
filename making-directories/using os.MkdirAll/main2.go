package main

import (
	"fmt"
	"os"
)

func main() {
	//func Mkdir(path string, perm FileMode) error
	if err := os.MkdirAll("dir/subdir", 0700); err != nil {
		fmt.Println(err.Error())
	}
}
