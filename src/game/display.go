package game

import (
	"fmt"
	"strings"
)

func GetBoardString(board [8][8]piece) string {
	var b strings.Builder
	var tileColor string = ""
	fmt.Fprintf(&b, "---------------------------------\n")
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if (x+y)%2 == 0 {
				tileColor = "White"
			} else {
				tileColor = "Black"
			}
			fmt.Fprintf(&b, "%v", getTile(board[y][x], tileColor))
			if x == 7 {
				fmt.Fprintf(&b,"| %v", 8-y)
			}
		}
		fmt.Fprintf(&b, "\n---------------------------------\n")
	}
	fmt.Fprintf(&b, "  a   b   c   d   e   f   g   h")
	return b.String()
}

func getTile(piece piece, tileColor string) string {
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

	return fmt.Sprintf("|\033[%s;%sm %s \033[0m", pieceColorDigit, tileColorDigit, pieceStr)
}
