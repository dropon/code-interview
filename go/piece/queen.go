package piece

type Queen struct {
	BasePiece
}

func NewQueen(x, y int, color Color) *Queen {
	symbol := "♕"
	if color == ColorBlack {
		symbol = "♛"
	}

	return &Queen{
		BasePiece: BasePiece{
			symbol: symbol,
			color:  color,
			x:      x,
			y:      y,
		},
	}
}

func (q *Queen) Moves() [][]int {
	// not implemented
	return [][]int{}
}
