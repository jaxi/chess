package chess

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Board is where the chess is played on.
type Board struct {
	squares [][]Square
	turn    Side
}

// Turn shows which side is playing now
func (b *Board) Turn() Side {
	return b.turn
}

// NewBoard creates a new board that has pieces been placed
func NewBoard() *Board {
	sq := make([][]Square, 8)

	for i := 0; i < 8; i++ {
		sq[i] = make([]Square, 8)
		for j := 0; j < 8; j++ {
			sq[i][j] = EmptySquare{}
		}
	}

	for i := 0; i < 8; i++ {
		sq[1][i] = Pawn{Piece{PAWN, WHITE}, &Movable{}}
		sq[6][i] = Pawn{Piece{PAWN, BLACK}, &Movable{}}
	}

	for _, v := range []Side{WHITE, BLACK} {
		var i int
		if v == WHITE {
			i = 0
		} else {
			i = 7
		}
		sq[i][0] = Rook{Piece{ROOK, v}, &Movable{}}
		sq[i][1] = Knight{Piece{KNIGHT, v}}
		sq[i][2] = Bishop{Piece{BISHOP, v}}
		sq[i][3] = Queen{Piece{QUEEN, v}}
		sq[i][4] = King{Piece{KING, v}, &Movable{}}
		sq[i][5] = Bishop{Piece{BISHOP, v}}
		sq[i][6] = Knight{Piece{KNIGHT, v}}
		sq[i][7] = Rook{Piece{ROOK, v}, &Movable{}}
	}

	return &Board{
		squares: sq,
		turn:    WHITE,
	}
}

func printBar() string {
	var nbf bytes.Buffer
	nbf.WriteString("   ")
	for i := 0; i < 8; i++ {
		nbf.WriteString("+---")
	}
	nbf.WriteString("+\n")
	return nbf.String()
}

func printNum() string {
	var nbf bytes.Buffer

	nbf.WriteString("   ")
	for i := 0; i < 8; i++ {
		nbf.WriteString("  ")
		nbf.WriteString(string('a' + i))
		nbf.WriteString(" ")
	}
	nbf.WriteString(" \n")
	return nbf.String()
}

// String - aka. turn the board into a string
func (b *Board) String() string {

	var bf bytes.Buffer
	bar := printBar()
	numBar := printNum()

	bf.WriteString(numBar)
	for i := 7; i >= 0; i-- {
		bf.WriteString(bar)
		bf.WriteString(" " + strconv.Itoa(i+1) + " ")
		for j := 0; j < 8; j++ {
			bf.WriteString("|" + b.squares[i][j].String())
		}
		bf.WriteString("| " + strconv.Itoa(i+1) + " ")
		bf.WriteString("\n")
	}
	bf.WriteString(bar)
	bf.WriteString(numBar)

	return bf.String()
}

// Move the piece on the board
func (b *Board) Move(pos1, pos2 Position) bool {
	if !pos1.IsValid() || !pos2.IsValid() {
		return false
	}

	p := b.squares[pos1.x][pos1.y]

	if p.Side() != b.turn {
		return false
	}

	return p.Move(b, pos1, pos2)
}

// Looping the game
func (b *Board) Looping() {
	tw := map[Side]string{
		WHITE: "White",
		BLACK: "Black",
	}
	fmt.Println(b)

	for {
		fmt.Printf("%s's turn:", tw[b.turn])
		var i, j, k, l int
		fmt.Scanf("%d %c %d %c\n", &i, &j, &k, &l)
		if b.Move(Position{i - 1, j - 'a'}, Position{k - 1, l - 'a'}) {
			b.turn = b.turn%2 + 1
			fmt.Println(b)
		} else {
			fmt.Println("The move is not quite right, please try again.")
		}
	}
}

// Move - The Movement in one turn
type Move struct {
	pos1 Position
	pos2 Position
}

// NewMove create a new Move struct
func NewMove(i, j, k, l int) Move {
	return Move{
		Position{i, j},
		Position{k, l},
	}
}

// Callback is the standard protocol can connect to the game
type Callback interface {
	ShowTurn(b *Board)
	RenderBoard(b *Board)
	FetchMove() (Move, error)
	ErrorMessage(b *Board)
}

// AdvanceLooping loop the gamve with more options
func (b *Board) AdvanceLooping(c Callback) {
	c.RenderBoard(b)

	for {
		var move Move
		var err error

		for {
			c.ShowTurn(b)
			move, err = c.FetchMove()
			if err != nil {
				if err == io.EOF {
					os.Exit(0)
				}
				fmt.Println("read error:", err)
				c.ErrorMessage(b)
			} else {
				break
			}
		}

		if b.Move(move.pos1, move.pos2) {
			b.turn = b.turn%2 + 1
			c.RenderBoard(b)
		} else {
			c.ErrorMessage(b)
		}
	}
}
