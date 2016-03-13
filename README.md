# Chess

A chess game that _just works™_.


# How to use it

Run the code below

```go
package main

import "github.com/jaxi/chess"

func main() {
	b := chess.NewBoard()
	b.Looping()
}
```

And then run the code above and it pretty much looks like this (let me know if it doesn't)

```
➜  examples git:(master) go run main.go
     a   b   c   d   e   f   g   h
   +---+---+---+---+---+---+---+---+
 8 | ♜ | ♞ | ♝ | ♛ | ♚ | ♝ | ♞ | ♜ | 8
   +---+---+---+---+---+---+---+---+
 7 | ♟ | ♟ | ♟ | ♟ | ♟ | ♟ | ♟ | ♟ | 7
   +---+---+---+---+---+---+---+---+
 6 |   |   |   |   |   |   |   |   | 6
   +---+---+---+---+---+---+---+---+
 5 |   |   |   |   |   |   |   |   | 5
   +---+---+---+---+---+---+---+---+
 4 |   |   |   |   |   |   |   |   | 4
   +---+---+---+---+---+---+---+---+
 3 |   |   |   |   |   |   |   |   | 3
   +---+---+---+---+---+---+---+---+
 2 | ♙ | ♙ | ♙ | ♙ | ♙ | ♙ | ♙ | ♙ | 2
   +---+---+---+---+---+---+---+---+
 1 | ♖ | ♘ | ♗ | ♕ | ♔ | ♗ | ♘ | ♖ | 1
   +---+---+---+---+---+---+---+---+
     a   b   c   d   e   f   g   h
```
