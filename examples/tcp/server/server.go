package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"

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
func (tc TCPCallback) FetchMove() (chess.Move, error) {
	nums := make([]int, 4)
	var err error
	text, err := bufio.NewReader(tc.conn).ReadString('\n')

	if err != nil {
		return chess.Move{}, err
	}

	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)

	for i := 0; i < 4; i++ {
		var s string

		if !scanner.Scan() {
			break
		}

		s = string(scanner.Bytes())

		if i%2 == 0 {
			nums[i], err = strconv.Atoi(strings.TrimSpace(s))

			if err != nil {
				return chess.Move{}, err
			}
		} else {
			s := []byte(strings.TrimSpace(s))
			nums[i] = int(s[0])
		}
	}

	return chess.NewMove(nums[0]-1, nums[1]-'a', nums[2]-1, nums[3]-'a'), nil
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
