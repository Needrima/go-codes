package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func check(msg string, err error) {
	if err != nil {
		panic(fmt.Sprintf("%s: %w", msg, err))
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	check("Connection", err)
	defer conn.Close()

	fmt.Print("Enter something and press enter:")

	write(conn)

}

func write(conn net.Conn) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		conn.Write([]byte(input))
	}
}
