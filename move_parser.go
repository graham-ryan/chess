package main

import (
	"fmt"
	"strings"
)

type FailedParseError struct {}
func (e *FailedParseError) Error() string {
	return "We parsed a type that doesn't exist, I guess? Weird."
}

// Remove symbols not necessary to calculating the move
func RemoveSpecialCharacters(move *string) {
	chars := [...]string{"x","d","=","+","#",}
	for _, char := range chars {
		_,*move,_ = strings.Cut(*move, char)
	}
}

type moveType interface {
	process(board *[8][8]Piece, turn string) error
}

type normalMove struct {
	piecePrime piecePrime
	col byte
	row byte
}

type pawnMove struct {
	col byte
	row byte
}

type pawnTakes struct {
	fromCol byte
	toCol byte
	row byte
}

type pawnPromotes struct {
	fromCol byte
	toCol byte
	row byte
	pawnPromotionPiece byte
}

type kingside struct {}

type queenside struct {}

type piecePrime struct {
	// TODO
}

type column struct {
	val byte
}

// Determines if the move string is valid algebraic notation.
//
// Algebraic notation (e.g. e4, Nh3, Qb6) has the following grammar:
// {Move} -> {NormalMove}|{PawnMove}|{PawnTakes}|{PawnPromotes}|{Kingside}|{Queenside}
// {NormalMove} -> {PiecePrime}{Column}{Row}
// {PawnMove} -> {Column}{Row}
// {PawnTakes} -> {Column}{Column}{Row}
// {PawnPromotes} -> {Column}{Column}{Row}{PawnPromotionPiece} TODO ambiguous if takes or regular move
// {Kingside} -> 0-0
// {Queenside} -> 0-0-0
// {PiecePrime} -> {Piece}
// {PiecePrime} -> {Piece}{Column}
// {PiecePrime} -> {Piece}{Row}
// {PiecePrime} -> {Piece}{Column}{Row}
// {Piece} -> Q|K|N|B|R  
// {PawnPromotionPiece} -> Q|N|B|R
// {Column} -> a|b|c|d|e|f|g|h
// {Row} -> 1|2|3|4|5|6|7|8
func ParseMove(move string) (moveType, error) {
	//notation_pieces := [...]string{"Q","K","N","B","R"}
	//pawn_promotion_pieces := [...]string{"Q","N","B","R"}

	fmt.Print(len(move))
	if move == "O-O" {
		return kingside{}, nil
	} else if move == "O-O-O" {
		return queenside{}, nil
	} else if len(move) == 3 {
		//if notation_pieces move[0]
		return pawnMove{move[0],move[1]}, nil
	} else if len(move) == 4 {
		return pawnTakes{move[0],move[1],move[2]}, nil
	} else if len(move) == 5 { 
		return pawnMove{move[0],move[1]}, nil
	} else {
		return nil, &FailedParseError{}
	}
}