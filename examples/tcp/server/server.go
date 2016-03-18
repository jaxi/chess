package main

import (
	"bufio"
	"errors"
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
	scanner := bufio.NewScanner(tc.conn)
	scanner.Split(bufio.ScanWords)
	nums := make([]int, 4)
	i := 0
	var err error

	for i = 0; i < 4 && scanner.Scan(); i++ {
		s := scanner.Text()
		if i%2 == 0 {
			nums[i], err = strconv.Atoi(strings.TrimSpace(s))
			if err != nil {
				return chess.Move{}, errors.New("Invalid input")
			}
		} else {
			if len(s) == 1 {
				nums[i] = int(([]byte(s))[0])
			} else {
				return chess.Move{}, errors.New("Invalid input")
			}
		}
	}

	if i == 0 {
		return chess.Move{}, nil
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

	b := chess.NewBoard()

	for {
		conn, _ := ln.Accept()
		defer conn.Close()
		tc := TCPCallback{conn: conn}

		b.AdvanceLooping(tc)
	}

}
