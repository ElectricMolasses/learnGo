package main

import (
	"github.com/nsf/termbox-go"
)

var players [2]string
var currentTurn int

func main() {
	err := termbox.Init()

	if err != nil {
		panic(err)
	}
	defer termbox.Close() // Will close termbox after main ends.

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	running := true

	players[0] = "Player1"
	players[1] = "Player2"

	currentTurn = 0

	// 2D slice for board
	board := [][]rune{
		[]rune{'‗', '‗', '‗'},
		[]rune{'‗', '‗', '‗'},
		[]rune{'‗', '‗', '‗'},
	}

	printBoard(board)

	for running {
		const coldef = termbox.ColorDefault
		termbox.Flush()

		// Poll event blocks the buffer so we don't need to manually
		// time the main loop.
		ev := termbox.PollEvent()

		if currentTurn == 0 {
			playTurn('X', ev)
			board[ev.MouseX][ev.MouseY] = 'X'
		} else {
			playTurn('O', ev)
			board[ev.MouseX][ev.MouseY] = 'Y'
		}

		checkVictory(board)

		if ev.Key == termbox.KeyEsc {
			running = false
		}
	}
}

func nextTurn(col1, col2 termbox.Attribute) {
	if currentTurn == 0 {
		currentTurn++
	} else {
		currentTurn--
	}
	for i := 0; i < len(players[currentTurn]); i++ {
		// 3 14
		termbox.SetCell(14+i, 3, rune(players[currentTurn][i]), col1, col2)
		// 10 is the made up character limit for names
		if i < 10 {
			for j := 1; j <= 10; j++ {
				termbox.SetCell(14+i+j, 3, ' ', col1, col2)
			}
		}
	}
}

func playTurn(symbol rune, ev termbox.Event) {
	coldef := termbox.ColorDefault
	if ev.Key == termbox.MouseLeft {
		x := ev.MouseX
		y := ev.MouseY
		// Check that click is on valid space
		if x < 3 && y < 3 {
			termbox.SetCell(x, y, symbol, coldef, coldef)

			nextTurn(coldef, coldef)
		}
	}
}

func printBoard(board [][]rune) {
	coldef := termbox.ColorDefault
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			termbox.SetCell(i, j, '‗', coldef, coldef)
		}
	}
	currentTurnText := []rune("Current Turn: Player1")
	for i := 0; i < len(currentTurnText); i++ {
		termbox.SetCell(i, 3, currentTurnText[i], coldef, coldef)
	}
}

func checkVictory(board [][]rune) {
	for i := 0; i < len(board); i++ {
		// xWin := true
		// yWin := true
		for j := 0; j < len(board); j++ {
			if board[i][j] != 'X' {
				// xWin = false
			}
		}
	}
}
