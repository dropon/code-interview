package piece

type Color string

const (
	ColorWhite Color = "white"
	ColorBlack Color = "black"
)

type Piece interface {
	Symbol() string
	Color() Color
	Position() (int, int)
	GetMoves() [][]int
}

// BasePiece is embedded in concrete piece types.
type BasePiece struct {
	symbol string
	color  Color
	x, y   int
}

func (p *BasePiece) Symbol() string       { return p.symbol }
func (p *BasePiece) Color() Color         { return p.color }
func (p *BasePiece) Position() (int, int) { return p.x, p.y }
