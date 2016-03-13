package main

import (
	"bufio"
	"io"
	"net"
	"os"
)

import "fmt"

func streamingOutput(conn net.Conn) {
	tmp := make([]byte, 256)
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		fmt.Printf(string(tmp[:n]))
	}
}

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")

	go streamingOutput(conn)

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, string(text))
	}
}
