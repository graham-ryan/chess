package main

import (
	"strings"
)

type FailedParseError struct {}
func (e *FailedParseError) Error() string {
	return "We parsed a type that doesn't exist, I guess? Weird."
}

// Parses the move string, validates that it's possible (TODO), and makes the move.
func ProcessMove(board *[8][8]Piece, moveStr string, turn string) (err error) {
	//RemoveSpecialCharacters(&moveStr)
	move := ParseMove(moveStr)
	return move.move.processMove(board, turn)
}

// Remove symbols not necessary to calculating the move
func RemoveSpecialCharacters(move *string) {
	chars := [...]string{"x","d","=","+","#",}
	for _, char := range chars {
		_,*move,_ = strings.Cut(*move, char)
	}
}

type moveType interface {
	processMove(board *[8][8]Piece, turn string) error
}

type normalMoveNode struct {
	piecePrime piecePrimeNode
	column byte
	row byte
}

type pawnMoveNode struct {
	column byte
	row byte
}

type pawnTakesNode struct {
	fromColumn byte
	toColumn byte
	row byte
}

type pawnPromotesNode struct {
	fromColumn byte
	toColumn byte
	row byte
	pawnPromotionPiece byte
}

type kingsideNode struct {}

type queensideNode struct {}

type piecePrimeNode struct {
	// TODO
}

type moveNode struct {
	move moveType
}

// Determines if the move string is valid algebraic notation.
//
// Algebraic notation (e.g. e4, Nh3, Qb6) has the following grammar:
// {Move} -> {NormalMove}|{PawnMove}|{PawnTakes}|{PawnPromotes}|{Kingside}|{Queenside}
// {NormalMove} -> {PiecePrime}{Column}{Row}
// {PawnMove} -> {Column}{Row}
// {PawnTakes} -> {Column}{Column}{Row}
// {PawnPromotes} -> {Column}{Column}{Row}{PawnPromotionPiece}
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
func ParseMove(move string) moveNode {
	return moveNode{pawnMoveNode{move[0],move[1]}}
}