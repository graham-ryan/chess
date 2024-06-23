package game

import (
	"fmt"
	"testing"
)

func compareBoards(arr1, arr2 [8][8]piece) bool {
  return fmt.Sprintf("%v",arr1) == fmt.Sprintf("%v",arr2)
}

func TestScandanavian(t *testing.T) {
	moves := [...]string{"e4","d5","ed5"}
	g := StartGame()
	for _, mv := range moves {
		err := g.ProcessMove(mv)
		if err != nil {
			t.Fatalf("Got error processing move %v", err)
		}
	}

	have := g.GetBoard()

	want := [8][8]piece{
		{{'R',"Red"},{'N',"Red"},{'B',"Red"},{'Q',"Red"},{'K',"Red"},{'B',"Red"},{'N',"Red"},{'R',"Red"}},
		{{'p',"Red"},{'p',"Red"},{'p',"Red"},{},{'p',"Red"},{'p',"Red"},{'p',"Red"},{'p',"Red"}},
		{{},{},{},{},{},{},{},{}}, 
		{{},{},{},{'p',"Blue"},{},{},{},{}},
		{{},{},{},{},{},{},{},{}},
		{{},{},{},{},{},{},{},{}},
		{{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"}},
		{{'R',"Blue"},{'N',"Blue"},{'B',"Blue"},{'Q',"Blue"},{'K',"Blue"},{'B',"Blue"},{'N',"Blue"},{'R',"Blue"}},
	}

	if !compareBoards(have,want) {
		t.Fatalf("Have:\n %v\n\n Want:\n %v", GetBoardString(have), GetBoardString(want))
	}
}

func TestScholarsMate(t *testing.T) {
	moves := [...]string{"e4", "e5", "Bc4", "Nc6", "Qh5", "Nf6", "Qf7"}
	g := StartGame()
	for _, mv := range moves {
		err := g.ProcessMove(mv)
		if err != nil {
			t.Fatalf("Got error processing move %v", err)
		}
	}

	have := g.GetBoard()

	want := [8][8]piece{
		{{'R',"Red"},{},{'B',"Red"},{'Q',"Red"},{'K',"Red"},{'B',"Red"},{},{'R',"Red"}},
		{{'p',"Red"},{'p',"Red"},{'p',"Red"},{'p',"Red"},{},{'Q',"Blue"},{'p',"Red"},{'p',"Red"}},
		{{},{},{'N',"Red"},{},{},{'N',"Red"},{},{}}, 
		{{},{},{},{},{'p',"Red"},{},{},{}},
		{{},{},{'B',"Blue"},{},{'p',"Blue"},{},{},{}},
		{{},{},{},{},{},{},{},{}},
		{{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"}},
		{{'R',"Blue"},{'N',"Blue"},{'B',"Blue"},{},{'K',"Blue"},{},{'N',"Blue"},{'R',"Blue"}},
	}

	if !compareBoards(have,want) {
		t.Fatalf("Have:\n %v\n\n Want:\n %v", GetBoardString(have), GetBoardString(want))
	}
}