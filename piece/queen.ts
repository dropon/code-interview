import { newPiece, type getMovesFunction } from "./piece";

export function newQueen(x: number, y: number, color: "black" | "white") {
  let symbol = "♕";
  if (color === "black") {
    symbol = "♛";
  }

  const getMoves: getMovesFunction = () => {
    const moves: number[][] = [];

    return moves;
  };

  return newPiece(symbol, color, x, y, getMoves);
}
