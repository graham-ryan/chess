package main

import (
	"bufio"
	chess "chess/src/game"
	"fmt"
	"os"
)

func main() {
	// Start a new game
	g := chess.StartGame()

	// Game loop
	message := ""
	var err error 
	for {
		g.PrintBoard()
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("\n%s\n\n",message)
		fmt.Printf("%s's turn: ", g.GetTurn())
		move, _ := reader.ReadString('\n')
		err = g.ProcessMove(move)
		if err != nil {
			message = err.Error()
		} else {
			message = ""
		}
	}
}