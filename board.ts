import type { Piece } from "./piece/piece";

type Board = {
  size: number;
  board: (Piece | null)[][];
  print: Function;
  placePiece: Function;
  movePiece: Function;
};

const a = "A".charCodeAt(0);

const startFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR";

export function newBoard(n: number): Board {
  const board: (Piece | null)[][] = Array.from(Array(n), () =>
    new Array(n).fill(null),
  );

  function placePiece(piece: Piece) {
    board[piece.x][piece.y] = piece;
  }

  // TODO
  function readFEN(fenString: string) {}

  function movePiece(piece: Piece, x: number, y: number) {
    board[piece.x][piece.y] = null;
    board[x][y] = piece;
  }

  function print() {
    console.write("".padStart(3));
    for (let i = 0; i < n; i++) {
      console.write(String.fromCharCode(a + i).padStart(3));
    }

    console.write("\n");

    for (let j = 0; j < n; j++) {
      console.write((j + 1).toString().padStart(3));
      for (let i = 0; i < n; i++) {
        const p = board[i][j];
        if (!p) {
          console.write("X".padStart(3, " "));
          continue;
        }

        console.write(p!.symbol.padStart(3, " "));
      }
      console.write("\n");
    }
  }

  return {
    size: n,
    board,
    print,
    placePiece,
    movePiece,
  };
}
