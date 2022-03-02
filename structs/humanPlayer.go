package structs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type HumanPlayer struct {
	marker string
	name   string
}

func NewHumanPlayer(marker string, name string) *HumanPlayer {
	return &HumanPlayer{
		marker: marker,
		name:   name,
	}
}

func (hp HumanPlayer) GetMarker() string {
	return hp.marker
}

func (hp HumanPlayer) GetName() string {
	return hp.name
}

func (hp HumanPlayer) Move(gameBoard GameBoard) {
	gameBoard.PrintBoard()
	isValidMove := false
	var input int
	for !isValidMove {
		input = playerInput()
		isValidMove = gameBoard.CheckLegalGrid(input)
		if !isValidMove {
			fmt.Println("Please choose an empty grid!")
		}
	}
	gameBoard.MarkGrid(input, hp.marker)
}

func playerInput() int {
	scanner := bufio.NewScanner(os.Stdin)
	isValid := false
	var result int
	for !isValid {
		fmt.Print("Input [1-9]: ")
		scanner.Scan()
		input := scanner.Text()
		validInput, err := strconv.Atoi(input)
		if err != nil || validInput < 1 || validInput > 9 {
			fmt.Println("Please input between 1 and 9")
		} else {
			isValid = true
			result = validInput
		}
	}
	return result
}
