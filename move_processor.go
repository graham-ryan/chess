package main

// Parses the move string, validates that it's possible (TODO), and makes the move.
func ProcessMove(board *[8][8]Piece, moveStr string, turn string) (err error) {
	//RemoveSpecialCharacters(&moveStr)
	move, err := ParseMove(moveStr)
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

func (move pawnMove) process(board *[8][8]Piece, turn string) error {
	// Find the pawn that must move
	// Convert ASCII representation of row and col to our 2d array
	col := int(move.col) - 97
	row := 7 - (int(move.row) - 49)
	//fmt.Printf("col: %v, row: %v",col,row)
	//time.Sleep(1*time.Second)

	// col and row must be in range [0-7]
	if col < 0 || col > 7 || row < 0 || row > 7 {
		return OutOfRangeError{}
	} 
	forward := getForward(turn)
	
	if board[row-forward][col].class == 'p' && board[row-forward][col].color == turn {
		// Pawn moves one sqaure
		board[row][col] = board[row-forward][col]
		board[row-forward][col] = Piece{}
	} else if board[row-forward*2][col].class == 'p' && board[row-forward*2][col].color == turn {
		// Pawn moves two squares
		board[row][col] = board[row-forward*2][col]
		board[row-forward*2][col] = Piece{}
	} else {
		return PawnNotFoundError{}
	}

	return nil
}

func (move pawnTakes) process(board *[8][8]Piece, turn string) error {
	// Find the pawn that must move
	// Convert ASCII representation of row and col to our 2d array
	fromCol := int(move.fromCol) - 97
	toCol := int(move.toCol) - 97
	row := 7 - (int(move.row) - 49)

	// col and row must be in range [0-7]
	if fromCol < 0 || fromCol > 7 || toCol < 0 || toCol > 7 || row < 0 || row > 7 {
		return OutOfRangeError{}
	} 
	forward := getForward(turn)

	// TODO En Pessant
	if board[row][toCol].class != 0 && board[row][toCol].color != turn {
		// We found the piece we want to take
		// Is there a pawn on the from column?
		if board[row-forward][fromCol].color == turn {
			board[row][toCol] = board[row-forward][fromCol]
			board[row-forward][fromCol] = Piece{}
		} else {
			return PawnNothingToTakeError{}
		}
	} else {
		return PawnNothingToTakeError{}
	}

	return nil
}

func (p kingside) process(board *[8][8]Piece, turn string) error {
	return nil
}

func (p queenside) process(board *[8][8]Piece, turn string) error {
	return nil
}

