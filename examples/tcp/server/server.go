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

// TCPPlayer - The Implementation of TCP protocol for chess game
type TCPPlayer struct {
	turn chess.Side
	conn net.Conn
	io.Reader
	io.Writer
}

// ShowTurn displays the message about who is playing
func (tc TCPPlayer) ShowTurn(b *chess.Board) {
	tw := map[chess.Side]string{
		chess.WHITE: "White",
		chess.BLACK: "Black",
	}

	if tc.turn == b.Turn() {
		tc.conn.Write([]byte(tw[b.Turn()] + "'s turn: "))
	} else {
		tc.conn.Write([]byte("Waiting for " + tw[b.Turn()] + "'s play.\n"))
	}
}

// RenderBoard render the board in CLI
func (tc TCPPlayer) RenderBoard(b *chess.Board) {
	tc.conn.Write([]byte(b.String()))
}

// FetchMove returns the move from your STDIN
func (tc TCPPlayer) FetchMove() (chess.Move, error) {
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
		return chess.NewMove(-1, -1, -1, -1), nil
	}
	return chess.NewMove(nums[0]-1, nums[1]-'a', nums[2]-1, nums[3]-'a'), nil
}

// ErrorMessage shows the message about what's going wrong
func (tc TCPPlayer) ErrorMessage(b *chess.Board) {
	tc.conn.Write([]byte("Wait a minute. There's something wrong with your move!\n"))
}

func main() {
	fmt.Println("Launching server...")

	ln, _ := net.Listen("tcp", ":8081")

	b := chess.NewBoard()

	opened := make([]bool, 2)
	conns := make([]net.Conn, 2)

	for {

		for i := range conns {
			if opened[i] == false {
				conns[i], _ = ln.Accept()
				opened[i] = true
				defer conns[i].Close()
			}
		}

		tc1 := TCPPlayer{conn: conns[0], turn: chess.WHITE}
		tc2 := TCPPlayer{conn: conns[1], turn: chess.BLACK}

		idx := b.AdvanceLooping(tc1, tc2)
		opened[idx] = false
	}
}
