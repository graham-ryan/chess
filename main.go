package main

import (
	//"errors"
	"bufio"
	"fmt"
	"os"
	//"os"
)

type Piece struct {
	class uint8 // e.g. Q K N p... or empty
	color string // Red or Blue
}


func main() {
	board  := [8][8]Piece{
		{{'r',"Red"},{'N',"Red"},{'B',"Red"},{'K',"Red"},{'Q',"Red"},{'B',"Red"},{'N',"Red"},{'r',"Red"}},
		{{'p',"Red"},{'p',"Red"},{'p',"Red"},{'p',"Red"},{'p',"Red"},{'p',"Red"},{'p',"Red"},{'p',"Red"}},
		{{},{},{},{},{},{},{},{}}, 
		{{},{},{},{},{},{},{},{}},
		{{},{},{},{},{},{},{},{}},
		{{},{},{},{},{},{},{},{}},
		{{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"}},
		{{'r',"Blue"},{'N',"Blue"},{'B',"Blue"},{'K',"Blue"},{'Q',"Blue"},{'B',"Blue"},{'N',"Blue"},{'r',"Blue"}},
	}
	
	message := "" 
	// Game loop
	for {
		printBoard(&board)
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("%s\n",message)
		fmt.Print("Move: ")
		move, _ := reader.ReadString('\n')
		fmt.Printf("%s",move)
	}
}

func printBoard(board *[8][8]Piece) {
	var tileColor string = ""
	fmt.Println("\033[21F\033[J") // Clear the screen
	fmt.Println("---------------------------------")
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if (x+y)%2 == 0 {
				tileColor = "White"
			} else {
				tileColor = "Black"
			}
			printTile(board[y][x], tileColor)
			if x == 7 {
				fmt.Print("|")
			}
		}
		fmt.Println("\n---------------------------------")
	}
}

func printTile(piece Piece, tileColor string) {
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
		case 'r':
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

	fmt.Printf("|\033[%s;%s;5m %s \033[0m", pieceColorDigit, tileColorDigit, pieceStr)

}