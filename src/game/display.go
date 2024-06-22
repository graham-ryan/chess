package game

import (
	"fmt"
)

func (g game) PrintBoard() {
	var tileColor string = ""
	fmt.Println("\033[100F\033[J") // Clear the screen
	fmt.Println("---------------------------------")
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if (x+y)%2 == 0 {
				tileColor = "White"
			} else {
				tileColor = "Black"
			}
			printTile(g.board[y][x], tileColor)
			if x == 7 {
				fmt.Printf("| %v", 8-y)
			}
		}
		fmt.Println("\n---------------------------------")
	}
	fmt.Println("  a   b   c   d   e   f   g   h")
}

func printTile(piece piece, tileColor string) {
	pieceStr := ""
	switch piece.class {
	case 'Q':
		pieceStr = "Q"
	case 'K':
		pieceStr = "K"
	case 'B':
		pieceStr = "B"
	case 'N':
		pieceStr = "N"
	case 'R':
		pieceStr = "R"
	case 'p':
		pieceStr = "p"
	default:
		pieceStr = " "
	}

	var pieceColorDigit string
	if piece.color == "Red" {
		pieceColorDigit = "31"
	} else {
		pieceColorDigit = "34"
	}

	var tileColorDigit string
	if tileColor == "White" {
		tileColorDigit = "47"
	} else {
		tileColorDigit = "40"
	}

	fmt.Printf("|\033[%s;%sm %s \033[0m", pieceColorDigit, tileColorDigit, pieceStr)

}
