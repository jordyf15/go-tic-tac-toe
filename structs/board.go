package structs

import "fmt"

type GameBoard struct {
	boardGrids []string
}

func NewGameBoard() *GameBoard {
	return &GameBoard{
		boardGrids: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
	}
}

func (gb GameBoard) PrintBoard() {
	fmt.Printf("%v %v %v\n%v %v %v\n%v %v %v\n", gb.boardGrids[0], gb.boardGrids[1], gb.boardGrids[2],
		gb.boardGrids[3], gb.boardGrids[4], gb.boardGrids[5],
		gb.boardGrids[6], gb.boardGrids[7], gb.boardGrids[8])
}

func (gb GameBoard) GetBoardGrids() []string {
	return gb.boardGrids
}

func (gb GameBoard) MarkGrid(grid int, mark string) {
	gridIndex := grid - 1
	gb.boardGrids[gridIndex] = mark
}

func (gb GameBoard) CheckLegalGrid(grid int) bool {
	gridIndex := grid - 1
	return gb.boardGrids[gridIndex] != "x" && gb.boardGrids[gridIndex] != "o"
}

func (gb GameBoard) CheckWin() bool {
	for i := 0; i < 3; i++ {
		if gb.boardGrids[i*3] == gb.boardGrids[(i*3)+1] &&
			gb.boardGrids[(i*3)+1] == gb.boardGrids[(i*3)+2] {
			return true
		}
		if gb.boardGrids[0+i] == gb.boardGrids[3+i] &&
			gb.boardGrids[3+i] == gb.boardGrids[6+i] {
			return true
		}
	}
	if gb.boardGrids[0] == gb.boardGrids[4] &&
		gb.boardGrids[4] == gb.boardGrids[8] {
		return true
	}
	if gb.boardGrids[2] == gb.boardGrids[4] &&
		gb.boardGrids[4] == gb.boardGrids[6] {
		return true
	}
	return false
}

func (gb GameBoard) CheckDraw() bool {
	emptyGrids := 0
	for _, v := range gb.boardGrids {
		if v != "x" && v != "o" {
			emptyGrids++
		}
	}
	return emptyGrids <= 0
}

func (gb GameBoard) RestartBoard() {
	gb.boardGrids = []string{"1", "2", "3",
		"4", "5", "6",
		"7", "8", "9"}
}
