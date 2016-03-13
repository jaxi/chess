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
func (cc CliCallback) FetchMove() (chess.Move, error) {
	nums := make([]int, 4)
	var err error
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
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
func (cc CliCallback) ErrorMessage(b *chess.Board) {
	fmt.Println("Wait a minute. There's something wrong with your move!")
}

func main() {
	b := chess.NewBoard()
	b.AdvanceLooping(CliCallback{})
}
