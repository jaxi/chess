package main

import (
	"io"
	"net"
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
		var i, j, k, l int
		fmt.Scanf("%d %c %d %c\n", &i, &j, &k, &l)
		fmt.Printf("\n")

		fmt.Fprintf(conn, "%d %c %d %c\n", i, j, k, l)
	}
}
