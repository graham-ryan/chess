package game

import (
	"fmt"
	"strings"
	"time"
)

type FailedParseError struct{}

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
	chars := [...]string{"x", "d", "=", "+", "#"}
	for _, char := range chars {
		_, *move, _ = strings.Cut(*move, char)
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
	piece   rune
	fromCol *int
	fromRow *int
	toCol   int
	toRow   int
}

type pawnMove struct {
	col int
	row int
	pawnPromotionPiece *rune
}

type pawnTakes struct {
	fromCol int
	toCol   int
	toRow   int
	pawnPromotionPiece *rune
}

type kingside struct{}

type queenside struct{}

// Determines if the move string is valid algebraic notation and returns a struct.
// Algebraic notation (e.g. e4, Nh3, Qb6) has the following grammar:
func parseMove(mv string) (moveType, error) {
	notation_pieces := []rune{'Q', 'K', 'N', 'B', 'R'}

	if mv == "O-O" {
		return kingside{}, nil

	} else if mv == "O-O-O" {
		return queenside{}, nil

	} else if contains(notation_pieces, rune(mv[0])) {
		return parseNormalMove(mv)

	} else {
		return parsePawnMove(mv)	
	}
}

// Parsing a regular non-pawn move
// Assumes that the first char in move is in the set [Q|K|N|B|R]
func parseNormalMove(mv string) (moveType, error) {
	fmt.Printf("mv: %v", mv)
	time.Sleep(time.Second * 5)
	if len(mv) == 3 {
		// Simple like Nf3, Be4
		toCol, err := convertInputToArrayCoordinate(mv[1], "col")
		if err != nil {
			return nil, err
		}
		toRow, err := convertInputToArrayCoordinate(mv[2], "row")
		if err != nil {
			return nil, err
		}
		return normalMove{rune(mv[0]), nil, nil, toCol, toRow}, nil

	} else if len(mv) == 4 {
		// Second character could be a row or a column
		if mv[1] >= 97 && mv[1] <=104 {
			// Column
			fromCol, err := convertInputToArrayCoordinate(mv[1], "col")
			if err != nil {
				return nil, err
			}
			toCol, err := convertInputToArrayCoordinate(mv[2], "col")
			if err != nil {
				return nil, err
			}
			toRow, err := convertInputToArrayCoordinate(mv[3], "row")
			if err != nil {
				return nil, err
			}
			return normalMove{rune(mv[0]), &fromCol, nil, toCol, toRow}, nil

		} else {
			// Row
			fromRow, err := convertInputToArrayCoordinate(mv[1], "row")
			if err != nil {
				return nil, err
			}
			toCol, err := convertInputToArrayCoordinate(mv[2], "col")
			if err != nil {
				return nil, err
			}
			toRow, err := convertInputToArrayCoordinate(mv[3], "row")
			if err != nil {
				return nil, err
			}
			return normalMove{rune(mv[0]), nil, &fromRow, toCol, toRow}, nil

		}

	} else if len(mv) == 5 {
		// Second character is column, third is row
		fromCol, err := convertInputToArrayCoordinate(mv[1], "col")
		if err != nil {
			return nil, err
		}
		fromRow, err := convertInputToArrayCoordinate(mv[2], "row")
		if err != nil {
			return nil, err
		}
		toCol, err := convertInputToArrayCoordinate(mv[3], "col")
		if err != nil {
			return nil, err
		}
		toRow, err := convertInputToArrayCoordinate(mv[4], "row")
		if err != nil {
			return nil, err
		}
		return normalMove{rune(mv[0]), &fromCol, &fromRow, toCol, toRow}, nil

	} else {
		return nil, &FailedParseError{}
	}
}

// Parsing a regular non-pawn move
// Assumes that the first char in move is a column
func parsePawnMove(mv string) (moveType, error) {
	pawn_promotion_pieces := []rune{'Q','N','B','R'}

	if len(mv) == 2 {
		col, err := convertInputToArrayCoordinate(mv[0], "col")
		if err != nil {
			return nil, err
		}
		row, err := convertInputToArrayCoordinate(mv[1], "row")
		if err != nil {
			return nil, err
		}
		return pawnMove{col, row, nil}, nil

	} else if len(mv) == 3 {
		// Check if move contains a pawn promotion piece
		if contains(pawn_promotion_pieces, rune(mv[2])) {
			// Pawn Takes
			col, err := convertInputToArrayCoordinate(mv[0], "col")
			if err != nil {
				return nil, err
			}
			row, err := convertInputToArrayCoordinate(mv[1], "row")
			if err != nil {
				return nil, err
			}
			promotion_piece := rune(mv[2])
			return pawnMove{col, row, &promotion_piece}, nil

		} else {
			// Pawn Takes
			fromCol, err := convertInputToArrayCoordinate(mv[0], "col")
			if err != nil {
				return nil, err
			}
			toCol, err := convertInputToArrayCoordinate(mv[1], "col")
			if err != nil {
				return nil, err
			}
			toRow, err := convertInputToArrayCoordinate(mv[2], "row")
			if err != nil {
				return nil, err
			}
			return pawnTakes{fromCol, toCol, toRow, nil}, nil
		}

	} else if len(mv) == 4 {
		// Pawn takes and promotes
		fromCol, err := convertInputToArrayCoordinate(mv[0], "col")
		if err != nil {
			return nil, err
		}
		toCol, err := convertInputToArrayCoordinate(mv[1], "col")
		if err != nil {
			return nil, err
		}
		toRow, err := convertInputToArrayCoordinate(mv[2], "row")
		if err != nil {
			return nil, err
		}
		// Ensure promotion piece is in set of promotion pieces
		promotion_piece := rune(mv[3])
		if !contains(pawn_promotion_pieces,promotion_piece) {
			return nil, &FailedParseError{}
		}
		return pawnTakes{fromCol, toCol, toRow, &promotion_piece}, nil

	} else {
		return nil, &FailedParseError{}
	}
}