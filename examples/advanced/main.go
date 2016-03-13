package main

import (
	"fmt"

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

	fmt.Printf("%s's turn:", tw[b.Turn()])
}

// RenderBoard render the board in CLI
func (cc CliCallback) RenderBoard(b *chess.Board) {
	fmt.Println(b)
}

// FetchMove returns the move from your STDIN
func (cc CliCallback) FetchMove() (int, int, int, int) {
	var i, j, k, l int
	fmt.Scanf("%d %c %d %c\n", &i, &j, &k, &l)
	return i, j, k, l
}

// ErrorMessage shows the message about what's going wrong
func (cc CliCallback) ErrorMessage(b *chess.Board) {
	fmt.Println("Wait a minute. There's something wrong with your move!")
}

func main() {
	b := chess.NewBoard()
	b.AdvanceLooping(CliCallback{})
}
