package game

import "fmt"

type piece struct {
	class uint8 // e.g. Q K N p... or empty
	color string // Red or Blue
}

type game struct {
	board [8][8]piece
	turn string
	gameOver bool
}

func StartGame() (g *game) {
	board := [8][8]piece{
		{{'r',"Red"},{'N',"Red"},{'B',"Red"},{'Q',"Red"},{'K',"Red"},{'B',"Red"},{'N',"Red"},{'r',"Red"}},
		{{'p',"Red"},{'p',"Red"},{'p',"Red"},{'p',"Red"},{'p',"Red"},{'p',"Red"},{'p',"Red"},{'p',"Red"}},
		{{},{},{},{},{},{},{},{}}, 
		{{},{},{},{},{},{},{},{}},
		{{},{},{},{},{},{},{},{}},
		{{},{},{},{},{},{},{},{}},
		{{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"}},
		{{'r',"Blue"},{'N',"Blue"},{'B',"Blue"},{'Q',"Blue"},{'K',"Blue"},{'B',"Blue"},{'N',"Blue"},{'r',"Blue"}},
	}
	turn := "Blue"
	gameOver := false
	return &game{board,turn,gameOver}
}

func (g game) IsGameOver() (bool) {
	return g.gameOver
}

func (g game) GetTurn() (string) {
	return g.turn
}

func (g game) GetBoard() ([8][8]piece) {
	return g.board
}

func (g game) String() string {
	return fmt.Sprintf("%v\nturn=%v\ngameOver=%v", g.board, g.turn, g.gameOver)
}

// Parses the move string, validates that it's possible (TODO), and makes the move.
func (g *game) ProcessMove(mv string) (err error) {
	//RemoveSpecialCharacters(&moveStr)
	move, err := parseMove(mv)
	if err != nil {
		return err
	}
	err = move.process(g)
	if err != nil {
		return err
	} 
	// No errors, change turns
	if g.turn == "Blue" {
		g.turn = "Red"
	} else {
		g.turn = "Blue"
	}
	return nil
}