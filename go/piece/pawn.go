// file: pawn.go
package piece

type Pawn struct {
	BasePiece
}

func NewPawn(x, y int, color Color) *Pawn {
	symbol := "♙"
	if color == ColorBlack {
		symbol = "♟"
	}

	return &Pawn{
		BasePiece: BasePiece{
			symbol: symbol,
			color:  color,
			x:      x,
			y:      y,
		},
	}
}

func (p *Pawn) GetMoves() [][]int {
	// not implemented (matches original TS)
	return [][]int{}
}
