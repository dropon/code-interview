package piece

type Bishop struct {
	BasePiece
}

func NewBishop(x, y int, color Color) *Bishop {
	symbol := "♗"
	if color == ColorBlack {
		symbol = "♝"
	}

	return &Bishop{
		BasePiece: BasePiece{
			symbol: symbol,
			color:  color,
			x:      x,
			y:      y,
		},
	}
}

func (b *Bishop) Moves() [][]int {
	// not implemented
	return [][]int{}
}
