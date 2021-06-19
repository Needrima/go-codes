package main

import (
	"fmt"
	"os"
)

func main() {
	//func Mkdir(name string, perm FileMode) error
	if err := os.Mkdir("dir", 0700); err != nil {
		fmt.Println(err.Error())
	}
}
