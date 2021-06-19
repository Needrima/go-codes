package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Print("Input value and press enter: ")
		reader := bufio.NewReader(os.Stdin)
		value, _ := reader.ReadString('\n')

		value = strings.Trim(value, " \r\n")

		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Input not a number")
		} else {
			num = num * 2
			fmt.Println(num)
		}
	}
}
