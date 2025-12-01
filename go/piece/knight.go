// file: knight.go
package piece

type Knight struct {
	BasePiece
}

func NewKnight(x, y int, color Color) *Knight {
	symbol := "♘"
	if color == ColorBlack {
		symbol = "♞"
	}

	return &Knight{
		BasePiece: BasePiece{
			symbol: symbol,
			color:  color,
			x:      x,
			y:      y,
		},
	}
}

func (k *Knight) GetMoves() [][]int {
	// not implemented
	return [][]int{}
}
