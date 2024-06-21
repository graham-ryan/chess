package game

import (
	"fmt"
	"strings"
	"time"
)

type FailedParseError struct {}
func (e *FailedParseError) Error() string {
	return "We parsed a type that doesn't exist, I guess? Weird."
}

// Convert ASCII representation of row to our board 2d array
func convertInputToArrayCoordinate(coordinate byte, rowOrCol string) (int, error) {
	val := 0
	if rowOrCol == "row" {
		val = 7 - (int(coordinate) - 49)
	} else if rowOrCol == "col" {
		val = int(coordinate) - 97
	} 

	// cols and rows must be in range [0-7]
	if val < 0 || val > 7 {
		return val, OutOfRangeError{}
	} 

	return val, nil 
}

// Remove symbols not necessary to calculating the move
func RemoveSpecialCharacters(move *string) {
	chars := [...]string{"x","d","=","+","#",}
	for _, char := range chars {
		_,*move,_ = strings.Cut(*move, char)
	}
}

func contains(arr []rune, txt rune) bool {
	for _, v := range arr {
		if v == txt {
			return true
		}
	}
	return false
}

type moveType interface {
	process(g *game) error
}

type normalMove struct {
	piece rune
	fromCol *int
	fromRow *int
	toCol int
	toRow int
}

type pawnMove struct {
	col int
	row int
}

type pawnTakes struct {
	fromCol int
	toCol int
	toRow int
}

type pawnPromotes struct {
	fromCol int
	toCol int
	toRow int
	pawnPromotionPiece int
}

type kingside struct {}

type queenside struct {}

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
func parseMove(move string) (moveType, error) {
	notation_pieces := []rune{'Q','K','N','B','R'}
	//pawn_promotion_pieces := [...]string{"Q","N","B","R"}

	// TODO pawn promotion???
	if move == "O-O" {
		return kingside{}, nil

	} else if move == "O-O-O" {
		return queenside{}, nil

	} else if contains(notation_pieces, rune(move[0])) {
		return parseNormalMove(move)

	} else if len(move) == 3 {
		col, err := convertInputToArrayCoordinate(move[0], "col")
		if err != nil {
			return nil, err
		}
		row, err := convertInputToArrayCoordinate(move[1], "row")
		if err != nil {
			return nil, err
		}
		return pawnMove{col,row}, nil

	} else if len(move) == 4 {
		fromCol, err := convertInputToArrayCoordinate(move[0], "col")
		if err != nil {
			return nil, err
		}
		toCol, err := convertInputToArrayCoordinate(move[1], "col")
		if err != nil {
			return nil, err
		}
		toRow, err := convertInputToArrayCoordinate(move[2], "row")
		if err != nil {
			return nil, err
		}
		return pawnTakes{fromCol,toCol,toRow}, nil

	} else {
		return nil, &FailedParseError{}
	}
}

// Parsing a regular non-pawn move
// Assumes that the first char in move is in the set [Q|K|N|B|R]
func parseNormalMove(move string) (moveType, error) {
	fmt.Printf("mv: %v", move)
	time.Sleep(time.Second*5)
	if len(move) == 4 {
		// Simple like Nf3, Be4
		toCol, err := convertInputToArrayCoordinate(move[2], "col")
		if err != nil {
			return nil, err
		}
		toRow, err := convertInputToArrayCoordinate(move[3], "row")
		if err != nil {
			return nil, err
		}
		return normalMove{rune(move[0]),nil,nil,toCol,toRow}, nil

	} else if len(move) == 5 {
		// Second character could be a row or a column
		return nil, &FailedParseError{}

	} else if len(move) == 6 {
		// Second character is column, third is row
		return nil, &FailedParseError{}

	} else {
		return nil, &FailedParseError{}
	}
}