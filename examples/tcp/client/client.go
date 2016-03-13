package main

import "net"

import "fmt"
import "bufio"

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")

	message, _ := bufio.NewReader(conn).ReadString('@')
	fmt.Println(message[:len(message)-1])

	for {
		var i, j, k, l int
		fmt.Scanf("%d %c %d %c\n", &i, &j, &k, &l)
		fmt.Print("Text to send: ")

		fmt.Fprintf(conn, "%d %c %d %c\n", i, j, k, l)

		message, _ := bufio.NewReader(conn).ReadString('@')
		fmt.Println(message[:len(message)-1])
	}
}
