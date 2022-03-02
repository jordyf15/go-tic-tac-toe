package structs

type Player interface {
	Move(gameBoard GameBoard)
	GetName() string
}
