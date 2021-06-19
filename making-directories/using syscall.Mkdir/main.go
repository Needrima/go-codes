package main

import (
	"syscall"
)

func main() {
	// func Mkdir(path string, mode uint32) (err error)
	if err := syscall.Mkdir("/Users/Needrima/Documents/gowoworkspace/src/go-practice/making-directories/new_dir", 0700); err == nil {
		println("Directory made")
	}
}
