package main

import (
	"fmt"
	"tic-tac-toe/structs"
)

func main() {
	boardGame := structs.NewGameBoard()
	humanPlayer := structs.NewHumanPlayer("x", "human")
	computerPlayer := structs.NewComputerPlayer("o", "computer")
	var currentPlayer structs.Player
	for !boardGame.CheckWin() && !boardGame.CheckDraw() {
		if currentPlayer == humanPlayer {
			fmt.Println("Computer turn")
			currentPlayer = computerPlayer
		} else {
			fmt.Println("Human turn")
			currentPlayer = humanPlayer
		}
		currentPlayer.Move(*boardGame)
	}
	if boardGame.CheckDraw() {
		fmt.Println("It's a draw!")
	} else {
		fmt.Printf("%v is The Winner!!\n", currentPlayer.GetName())
	}
	boardGame.PrintBoard()
}
