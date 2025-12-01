package piece

type Rook struct {
	BasePiece
}

func NewRook(x, y int, color Color) *Rook {
	symbol := "♖"
	if color == ColorBlack {
		symbol = "♜"
	}

	return &Rook{
		BasePiece: BasePiece{
			symbol: symbol,
			color:  color,
			x:      x,
			y:      y,
		},
	}
}

func (r *Rook) Moves() [][]int {
	// not implemented
	return [][]int{}
}
