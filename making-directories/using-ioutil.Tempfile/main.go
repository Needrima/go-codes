package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// func TempDir(dir, pattern string) (name string, err error)
	name, err := ioutil.TempDir(".", "New_dir-*")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Created directory %s\n", name)
	}
}
