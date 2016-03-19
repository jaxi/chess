package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jaxi/chess"
)

// CliCallback - The Implementation of CLI protocol for chess game
type CliCallback struct{}

// ShowTurn displays the message about who is playing
func (cc CliCallback) ShowTurn(b *chess.Board) {
	tw := map[chess.Side]string{
		chess.WHITE: "White",
		chess.BLACK: "Black",
	}

	fmt.Printf("%s's turn: ", tw[b.Turn()])
}

// RenderBoard render the board in CLI
func (cc CliCallback) RenderBoard(b *chess.Board) {
	fmt.Println(b)
}

// FetchMove returns the move from your STDIN
func (cc CliCallback) FetchMove() (move chess.Move, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("Invalid Input")
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	nums := make([]int, 4)

	for i := 0; i < 4 && scanner.Scan(); i++ {
		s := scanner.Text()
		if i%2 == 0 {
			nums[i], err = strconv.Atoi(strings.TrimSpace(s))
		} else {
			if len(s) == 1 {
				nums[i] = int(([]byte(s))[0])
			}
		}
	}
	move = chess.NewMove(nums[0]-1, nums[1]-'a', nums[2]-1, nums[3]-'a')
	return
}

// ErrorMessage shows the message about what's going wrong
func (cc CliCallback) ErrorMessage(b *chess.Board) {
	fmt.Println("Wait a minute. There's something wrong with your move!")
}

func main() {
	b := chess.NewBoard()
	b.AdvanceLooping(CliCallback{})
}
