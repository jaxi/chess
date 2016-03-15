package main

import (
	"strings"

	"github.com/jaxi/chess"
	"github.com/nsf/termbox-go"
)

func redrawAll(board *chess.Board) {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)
	w, h := termbox.Size()

	lines := strings.Split(board.String(), "\n")

	bh := len(lines)
	bw := len(lines[0])

	sy := (h - bh) / 2
	sx := (w - bw) / 2

	for i := 0; i < bh; i++ {
		for j, rv := range lines[i] {
			termbox.SetCell(sx+j, sy+i, rv, coldef, coldef)
		}
	}
	termbox.Flush()
}

func main() {
	board := chess.NewBoard()

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	redrawAll(board)

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
		redrawAll(board)
	}
}
