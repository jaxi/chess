package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/jaxi/chess"
)

// TCPCallback - The Implementation of TCP protocol for chess game
type TCPCallback struct {
	conn net.Conn
	io.Reader
	io.Writer
}

// ShowTurn displays the message about who is playing
func (tc TCPCallback) ShowTurn(b *chess.Board) {
	tw := map[chess.Side]string{
		chess.WHITE: "White",
		chess.BLACK: "Black",
	}

	tc.conn.Write([]byte(tw[b.Turn()] + "'s turn: "))
}

// RenderBoard render the board in CLI
func (tc TCPCallback) RenderBoard(b *chess.Board) {
	tc.conn.Write([]byte(b.String()))
}

// FetchMove returns the move from your STDIN
func (tc TCPCallback) FetchMove() (int, int, int, int) {
	var i, j, k, l int
	bts, _ := bufio.NewReader(tc.conn).ReadString('\n')

	message := string(bts)

	fmt.Sscanf(message, "%d %c %d %c\n", &i, &j, &k, &l)

	if j == 0 {
		os.Exit(0)
	}
	return i, j, k, l
}

// ErrorMessage shows the message about what's going wrong
func (tc TCPCallback) ErrorMessage(b *chess.Board) {
	tc.conn.Write([]byte("Wait a minute. There's something wrong with your move!\n"))
}

func main() {
	fmt.Println("Launching server...")

	ln, _ := net.Listen("tcp", ":8081")

	conn, _ := ln.Accept()
	defer conn.Close()

	b := chess.NewBoard()
	tc := TCPCallback{conn: conn}

	b.AdvanceLooping(tc)
}
