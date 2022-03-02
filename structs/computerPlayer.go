package structs

import (
	"math/rand"
)

type ComputerPlayer struct {
	marker string
	name   string
}

func NewComputerPlayer(marker string, name string) *ComputerPlayer {
	return &ComputerPlayer{
		marker: marker,
		name:   name,
	}
}

func (cp ComputerPlayer) GetMarker() string {
	return cp.marker
}

func (cp ComputerPlayer) GetName() string {
	return cp.name
}

func (cp ComputerPlayer) Move(gameboard GameBoard) {
	validGrids := getAllValidMoves(gameboard.GetBoardGrids())
	randomValidMoves := rand.Intn(len(validGrids))
	randomMove := validGrids[randomValidMoves]
	gameboard.MarkGrid(randomMove+1, cp.marker)
	gameboard.PrintBoard()
}

func getAllValidMoves(boardGrids []string) []int {
	possibleMoves := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	validMoves := []int{}
	for _, v := range possibleMoves {
		if boardGrids[v] != "x" && boardGrids[v] != "o" {
			validMoves = append(validMoves, v)
		}
	}
	return validMoves
}
