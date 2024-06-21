package game

import (
	"fmt"
	"time"
)

func getForward(turn string) int {
	if turn == "Blue" {
		return -1
	} else {
		return 1
	}
}

type PawnNotFoundError struct {}
func (PawnNotFoundError) Error() string {return "Can't find the pawn you want to move"}

type PawnNothingToTakeError struct{}
func (PawnNothingToTakeError) Error() string {return "Pawn has nothing to take on that square"}

type SomethingInTheWayError struct {}
func (SomethingInTheWayError) Error() string {return "There's something in the way of the piece you want to move"}

type OutOfRangeError struct {}
func (OutOfRangeError) Error() string {return "You indicated a row or column that doesn't exist"}

func (mv pawnMove) process(g *game) error {
	forward := getForward(g.turn)
	
	if g.board[mv.row-forward][mv.col].class == 'p' && g.board[mv.row-forward][mv.col].color == g.turn {
		// Pawn moves one sqaure
		g.board[mv.row][mv.col] = g.board[mv.row-forward][mv.col]
		g.board[mv.row-forward][mv.col] = piece{}
	} else if g.board[mv.row-forward*2][mv.col].class == 'p' && g.board[mv.row-forward*2][mv.col].color == g.turn {
		// Pawn moves two squares
		g.board[mv.row][mv.col] = g.board[mv.row-forward*2][mv.col]
		g.board[mv.row-forward*2][mv.col] = piece{}
	} else {
		return PawnNotFoundError{}
	}

	return nil
}

func (mv pawnTakes) process(g *game) error {
	forward := getForward(g.turn)

	// TODO En Pessant
	if g.board[mv.toRow][mv.toCol].class != 0 && g.board[mv.toRow][mv.toCol].color != g.turn {
		// We found the piece we want to take
		// Is there a pawn on the from column?
		if g.board[mv.toRow-forward][mv.fromCol].color == g.turn {
			g.board[mv.toRow][mv.toCol] = g.board[mv.toRow-forward][mv.fromCol]
			g.board[mv.toRow-forward][mv.fromCol] = piece{}
		} else {
			return PawnNotFoundError{}
		}
	} else {
		return PawnNothingToTakeError{}
	}

	return nil
}

// Kingside castles, if possible
func (mv kingside) process(g *game) error {
	return nil
}

// Queenside castles, if possible
func (mv queenside) process(g *game) error {
	return nil
}

func (mv normalMove) process(g *game) error {
	fmt.Println(mv)
	time.Sleep(10*time.Second)
	return nil
}