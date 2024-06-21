package game

import (
	"fmt"
	"testing"
)

func compareBoards(arr1, arr2 [8][8]piece) bool {
  return fmt.Sprintf("%v",arr1) == fmt.Sprintf("%v",arr2)
}

func TestScandanavian(t *testing.T) {
	moves := [...]string{"e4\n","d5\n","ed5\n"}
	g := StartGame()
	for _, mv := range moves {
		err := g.ProcessMove(mv)
		if err != nil {
			t.Fatalf("Got error processing move %v", err)
		}
	}

	have := g.GetBoard()

	want := [8][8]piece{
		{{'r',"Red"},{'N',"Red"},{'B',"Red"},{'Q',"Red"},{'K',"Red"},{'B',"Red"},{'N',"Red"},{'r',"Red"}},
		{{'p',"Red"},{'p',"Red"},{'p',"Red"},{},{'p',"Red"},{'p',"Red"},{'p',"Red"},{'p',"Red"}},
		{{},{},{},{},{},{},{},{}}, 
		{{},{},{},{'p',"Blue"},{},{},{},{}},
		{{},{},{},{},{},{},{},{}},
		{{},{},{},{},{},{},{},{}},
		{{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"},{},{'p',"Blue"},{'p',"Blue"},{'p',"Blue"}},
		{{'r',"Blue"},{'N',"Blue"},{'B',"Blue"},{'Q',"Blue"},{'K',"Blue"},{'B',"Blue"},{'N',"Blue"},{'r',"Blue"}},
	}

	if !compareBoards(have,want) {
		t.Fatalf("Board = %v,\n\n want this %v", have, want)
	}
}