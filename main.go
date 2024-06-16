package main

import (
	"bufio"
	"fmt"
	"os"
)

type Piece struct {
	class uint8 // e.g. Q K N p... or empty
	color string // Red or Blue
}

func main() {
	board := [8][8]Piece{
		{{'r',"Red"},{'N',"Red"},{'B',"Red"},{'K',"Red"},{'Q',"Red"},{'B',"Red"},{'N',"Red"},{'r',"Red"}},
		{{'p',"Red"},{'p',"Red"},{'p',"Red"},{'p',"Red"},{'p',"Red"},{'p',"Red"},{'p',"Red"},{'p',"Red"}},
		{{},{},{},{},{},{},{},{}}, 
		{{},{},{},{},{},{},{},{}},
		{{},{},{},{},{},{},{},{}},
		{{},{},{},{},{},{},{},{}},
		{{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"}},
		{{'r',"Blue"},{'N',"Blue"},{'B',"Blue"},{'K',"Blue"},{'Q',"Blue"},{'B',"Blue"},{'N',"Blue"},{'r',"Blue"}},
	}
	
	// Game loop
	message := ""
	turn := "Blue" // false for blue, true for red 
	var err error 
	for {
		printBoard(&board)
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("\n%s\n\n",message)
		fmt.Printf("%s's turn: ", turn)
		move, _ := reader.ReadString('\n')
		err = ProcessMove(&board, move, turn)
		if err != nil {
			message = err.Error()
		} else {
			message = ""
			// Change the turn
			if turn == "Blue" {
				turn = "Red"
			} else {
				turn = "Blue"
			}
		}
	}
}