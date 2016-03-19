package main

import (
	"strings"

	"github.com/jaxi/chess"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type terminalUICallback struct {
}

func (tuc terminalUICallback) ShowTurn(b *chess.Board) {

}

const coldef = termbox.ColorDefault

func (tuc terminalUICallback) RenderBoard(b *chess.Board) {
	termbox.Clear(coldef, coldef)

	boardPrint(b)

	_, h := termbox.Size()
	tbPrint(0, h-1, coldef, coldef, "Press ESC to exit...")
	termbox.Flush()
}

func (tuc terminalUICallback) FetchMove() (chess.Move, error) {
	return chess.Move{}, nil
}

func (tuc terminalUICallback) ErrorMessage(b *chess.Board) {

}

func tbPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func boardPrint(b *chess.Board) {
	w, h := termbox.Size()

	lines := strings.Split(b.TerminalString(), "\n")

	bh := len(lines)
	bw := len(lines[0])

	sy := (h - bh) / 2
	sx := (w - bw) / 2

	for i := 0; i < bh; i++ {
		for j, rv := range lines[i] {
			termbox.SetCell(sx+j, sy+i, rv, coldef, coldef)
		}
	}
}
func main() {
	board := chess.NewBoard()

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	go func() {
		board.AdvanceLooping(terminalUICallback{})
	}()

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
