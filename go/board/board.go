package board

import (
	"fmt"

	"github.com/dropon/code-interview/piece"
)

type Board struct {
	Size int
	Grid [][]piece.Piece
}

const startFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"

func NewBoard(n int) *Board {
	grid := make([][]piece.Piece, n)
	for i := range grid {
		grid[i] = make([]piece.Piece, n)
	}

	return &Board{
		Size: n,
		Grid: grid,
	}
}

func (b *Board) PlacePiece(p piece.Piece) {
	x, y := p.Position()
	b.Grid[x][y] = p
}

func (b *Board) ReadFEN(fenString string) {
	// intentionally left unimplemented
}

func (b *Board) MovePiece(p piece.Piece, x, y int) {
	px, py := p.Position()
	b.Grid[px][py] = nil
	b.Grid[x][y] = p
	// note: not updating internal (x,y) on the piece
	// which mirrors your original TS (it also didn't update piece.x / piece.y)
}

func (b *Board) Print() {
	fmt.Printf("%3s", "")
	for i := 0; i < b.Size; i++ {
		fmt.Printf("%3c", 'A'+rune(i))
	}
	fmt.Println()

	for j := 0; j < b.Size; j++ {
		fmt.Printf("%3d", j+1)
		for i := 0; i < b.Size; i++ {
			p := b.Grid[i][j]
			if p == nil {
				fmt.Printf("%3s", "X")
				continue
			}
			fmt.Printf("%3s", p.Symbol())
		}
		fmt.Println()
	}
}
