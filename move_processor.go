package main

import (
	"fmt"
	"time"
)

// Parses the move string, validates that it's possible (TODO), and makes the move.
func ProcessMove(board *[8][8]Piece, moveStr string, turn string) (err error) {
	//RemoveSpecialCharacters(&moveStr)
	move, err := parseMove(moveStr)
	if err != nil {
		return err
	}
	return move.process(board, turn)
}

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

func (mv pawnMove) process(board *[8][8]Piece, turn string) error {
	forward := getForward(turn)
	
	if board[mv.row-forward][mv.col].class == 'p' && board[mv.row-forward][mv.col].color == turn {
		// Pawn moves one sqaure
		board[mv.row][mv.col] = board[mv.row-forward][mv.col]
		board[mv.row-forward][mv.col] = Piece{}
	} else if board[mv.row-forward*2][mv.col].class == 'p' && board[mv.row-forward*2][mv.col].color == turn {
		// Pawn moves two squares
		board[mv.row][mv.col] = board[mv.row-forward*2][mv.col]
		board[mv.row-forward*2][mv.col] = Piece{}
	} else {
		return PawnNotFoundError{}
	}

	return nil
}

func (mv pawnTakes) process(board *[8][8]Piece, turn string) error {
	forward := getForward(turn)

	// TODO En Pessant
	if board[mv.toRow][mv.toCol].class != 0 && board[mv.toRow][mv.toCol].color != turn {
		// We found the piece we want to take
		// Is there a pawn on the from column?
		if board[mv.toRow-forward][mv.fromCol].color == turn {
			board[mv.toRow][mv.toCol] = board[mv.toRow-forward][mv.fromCol]
			board[mv.toRow-forward][mv.fromCol] = Piece{}
		} else {
			return PawnNothingToTakeError{}
		}
	} else {
		return PawnNothingToTakeError{}
	}

	return nil
}

// Kingside castles, if possible
func (mv kingside) process(board *[8][8]Piece, turn string) error {
	return nil
}

// Queenside castles, if possible
func (mv queenside) process(board *[8][8]Piece, turn string) error {
	return nil
}

func (mv normalMove) process(board *[8][8]Piece, turn string) error {
	fmt.Println(mv)
	time.Sleep(10*time.Second)
	return nil
}
