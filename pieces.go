package main

// PieceKind is the type representing for the pieces on the board
type PieceKind int

const (
	// EMPTY_SQUARE aka Empty
	EMPTY_SQUARE PieceKind = 0 + iota
	// PAWN (♙, ♟)
	PAWN
	// ROOK (♖, ♜)
	ROOK
	// KNIGHT (♘, ♞)
	KNIGHT
	// BISHOP (♗, ♝)
	BISHOP
	// QUEEN (♕, ♛)
	QUEEN
	// KING (♔, ♚)
	KING
)

// Position is the place where a piece stand
type Position struct {
	x int
	y int
}

// IsValid returns if a position is on the board or not
func (p Position) IsValid() bool {
	return p.x >= 0 && p.x < 8 && p.y >= 0 && p.y < 8
}

// Equal returns if two positions are in the same place
func (p Position) Equal(p2 Position) bool {
	return p.x == p2.x && p.y == p2.y
}

// Side can be either white or black
type Side int

const (
	// EMPTY is where neither white nor black piece stand on it
	EMPTY Side = 0 + iota
	// WHITE is the side that start the game second
	WHITE
	// BLACK move after the WHITE side is moved at the beginning
	BLACK
)

// Square represents for one of the 64 squares on the board
type Square interface {
	String() string
	PieceKind() PieceKind
	Side() Side
	Move(b *Board, pos1, pos2 Position) bool
}

// The Piece stand at wherever on the board
type Piece struct {
	pk PieceKind
	sd Side
}

func (p Piece) String() string {
	return "   "
}

// Side returns the side of the piece
func (p Piece) Side() Side {
	return p.sd
}

// Movable is used to attach to the piece that need recording if
// it's been moved or not
type Movable struct {
	moved bool
}

// isMoved check if a square is moved or not
func (m Movable) isMoved() bool {
	return m.moved
}

// Pawn - the weakest
type Pawn struct {
	Piece
	*Movable
}

func (p Pawn) String() string {
	if p.sd == WHITE {
		return " ♙ "
	}
	return " ♟ "
}

// PieceKind of Pawn
func (p Pawn) PieceKind() PieceKind { return PAWN }

// Move like a Pawn
func (p Pawn) Move(b *Board, pos1, pos2 Position) bool {
	s1 := b.squares[pos1.x][pos2.y]

	// Allowed move count
	amc := 1
	if p.moved == false {
		amc = 2
	}

	dir := 1
	if p.Side() == BLACK {
		dir = -1
	}

	for i := 1; i <= amc; i++ {
		dx := dir * i
		if !(Position{pos1.x + dx, pos2.y}).IsValid() {
			break
		}

		if pos2.Equal(Position{pos1.x + dx, pos2.y}) {
			s2 := b.squares[pos2.x][pos2.y]
			if s1.Side() == s2.Side() {
				return false
			}
			b.squares[pos2.x][pos2.y] = b.squares[pos1.x][pos1.y]
			movedPawn := b.squares[pos2.x][pos2.y].(Pawn)
			movedPawn.moved = true
			b.squares[pos1.x][pos1.y] = EmptySquare{}
			return true
		}
	}

	ds := [][]int{{1, 1}, {1, -1}}
	for _, d := range ds {
		x, y := pos1.x+d[0], pos1.y+d[1]
		if !(Position{x, y}).IsValid() {
			break
		}

		if pos2.Equal(Position{x, y}) {
			s2 := b.squares[pos2.x][pos2.y]
			if s1.Side() == s2.Side() {
				return false
			}
			b.squares[pos2.x][pos2.y] = b.squares[pos1.x][pos1.y]
			b.squares[pos1.x][pos1.y] = EmptySquare{}
			return true
		}
	}
	return false
}

// Rook - aka. the Tower
type Rook struct {
	Piece
	*Movable
}

func (r Rook) String() string {
	if r.sd == WHITE {
		return " ♖ "
	}
	return " ♜ "
}

// PieceKind of Rook
func (r Rook) PieceKind() PieceKind { return ROOK }

// Move like a Rook
func (r Rook) Move(b *Board, pos1, pos2 Position) bool {
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, dir := range dirs {
		dx, dy := dir[0], dir[1]
		for i, j := pos1.x+dx, pos1.y+dy; (Position{i, j}.IsValid()); i, j = i+dx, j+dy {
			p2 := b.squares[i][j]
			// blocked by it's own piece
			if r.Side() == p2.Side() {
				break
			}
			if pos2.Equal(Position{i, j}) {
				b.squares[pos2.x][pos2.y] = b.squares[pos1.x][pos1.y]
				b.squares[pos1.x][pos1.y] = EmptySquare{}

				movedRook := b.squares[pos2.x][pos2.y].(Rook)
				movedRook.moved = true
				return true
			}
		}
	}
	return false
}

// Knight - Piece that moves in the weird way
type Knight struct {
	Piece
}

func (k Knight) String() string {
	if k.sd == WHITE {
		return " ♘ "
	}
	return " ♞ "
}

// PieceKind of Knight
func (k Knight) PieceKind() PieceKind { return KNIGHT }

// Move like a Knight
func (k Knight) Move(b *Board, pos1, pos2 Position) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			for t := 0; t <= 1; t++ {
				dirs := []int{i, j}
				dirs[t] *= 2

				x, y := pos1.x+dirs[0], pos1.y+dirs[1]
				if pos2.Equal(Position{x, y}) {
					p2 := b.squares[pos2.x][pos2.y]

					if p2.Side() == k.Side() {
						return false
					}
					b.squares[pos2.x][pos2.y] = b.squares[pos1.x][pos1.y]
					b.squares[pos1.x][pos1.y] = EmptySquare{}
					return true
				}
			}
		}
	}
	return false
}

// Bishop - The bishop
type Bishop struct {
	Piece
}

func (b Bishop) String() string {
	if b.sd == WHITE {
		return " ♗ "
	}
	return " ♝ "
}

// PieceKind of Bishop
func (b Bishop) PieceKind() PieceKind { return BISHOP }

// Move like a Bishop
func (b Bishop) Move(bd *Board, pos1, pos2 Position) bool {
	dirs := [][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	for _, dir := range dirs {
		dx, dy := dir[0], dir[1]
		for i, j := pos1.x+dx, pos1.y+dy; (Position{i, j}.IsValid()); i, j = i+dx, j+dy {
			p2 := bd.squares[i][j]
			// blocked by it's own piece
			if b.Side() == p2.Side() {
				break
			}
			if pos2.Equal(Position{i, j}) {
				bd.squares[pos2.x][pos2.y] = bd.squares[pos1.x][pos1.y]
				bd.squares[pos1.x][pos1.y] = EmptySquare{}
				return true
			}
		}
	}
	return false
}

// Queen - The strongest piece on the board
type Queen struct {
	Piece
}

func (q Queen) String() string {
	if q.sd == WHITE {
		return " ♕ "
	}
	return " ♛ "
}

// PieceKind of Queen
func (q Queen) PieceKind() PieceKind { return QUEEN }

// Move like the Queen
func (q Queen) Move(b *Board, pos1, pos2 Position) bool {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				break
			}
			for i, j := pos1.x+dx, pos1.y+dy; (Position{i, j}.IsValid()); i, j = i+dx, j+dy {
				p2 := b.squares[i][j]
				// blocked by it's own piece
				if q.Side() == p2.Side() {
					break
				}
				if pos2.Equal(Position{i, j}) {
					b.squares[pos2.x][pos2.y] = b.squares[pos1.x][pos1.y]
					b.squares[pos1.x][pos1.y] = EmptySquare{}
					return true
				}
			}
		}
	}
	return false
}

// King - The piece to protect
type King struct {
	Piece
	*Movable
}

func (k King) String() string {
	if k.sd == WHITE {
		return " ♔ "
	}
	return " ♚ "
}

// PieceKind of KING
func (k King) PieceKind() PieceKind { return KING }

// Move in the King's way
func (k King) Move(b *Board, pos1, pos2 Position) bool {
	s1 := b.squares[pos1.x][pos2.y]

	diff := (pos1.x - pos2.x) * (pos1.x - pos2.x)
	if diff >= -1 && diff <= 1 && !pos1.Equal(pos2) {
		s2 := b.squares[pos2.x][pos2.y]
		if s1.Side() == s2.Side() {
			return false
		}
		b.squares[pos2.x][pos2.y] = b.squares[pos1.x][pos1.y]
		b.squares[pos1.x][pos1.y] = EmptySquare{}
		movedKing := b.squares[pos2.x][pos2.y].(King)
		movedKing.moved = true
		return true
	}
	return false
}

// EmptySquare - aka no piece on the board
type EmptySquare struct {
	Piece
}

// PieceKind of KING
func (es EmptySquare) PieceKind() PieceKind { return EMPTY_SQUARE }

// Move always returns false if the squre is empty
func (es EmptySquare) Move(b *Board, pos1, pos2 Position) bool { return false }
