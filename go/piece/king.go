// file: king.go
package piece

type King struct {
	BasePiece
}

func NewKing(x, y int, color Color) *King {
	symbol := "♔"
	if color == ColorBlack {
		symbol = "♚"
	}

	return &King{
		BasePiece: BasePiece{
			symbol: symbol,
			color:  color,
			x:      x,
			y:      y,
		},
	}
}

func (k *King) GetMoves() [][]int {
	// not implemented
	return [][]int{}
}
