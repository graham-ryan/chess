package main

import (
	"fmt"
	"time"
)

type Piece struct {
	class uint8 // e.g. Q K N p... or empty
	color string // Red or Blue
}

func getForward(turn string) int {
	if turn == "Blue" {
		return 1
	} else {
		return -1
	}
}

type PawnNotFoundError struct {}
func (PawnNotFoundError) Error() string {return "Can't find the pawn you want to move"}

type SomethingInTheWayError struct {}
func (SomethingInTheWayError) Error() string {return "There's something in the way of the piece you want to move"}

func (p pawnMoveNode) processMove(board *[8][8]Piece, turn string) error {
	// Find the pawn that must move
	y := int(p.column) - 96
	x := int(p.row) - 48
	fmt.Println(x)
	fmt.Println(y)
	time.Sleep(5*time.Second)
	forward := getForward(turn)
	

	if board[y-forward][x].class == 'p' {

		board[y][x] = board[y-forward][x]
		board[y-forward][x] = Piece{}
	} else if board[y-forward*2][x].class == 'p' {

	} else {
		return PawnNotFoundError{}
	}

	return nil
}