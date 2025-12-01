export type getMovesFunction = () => number[][];

export type Piece = {
  symbol: string;
  x: number;
  y: number;
  getMoves: () => number[][];
  color: "black" | "white";
};

export function newPiece(
  symbol: string,
  color: "black" | "white",
  x: number,
  y: number,
  getMoves: getMovesFunction,
): Piece {
  return {
    symbol,
    x,
    y,
    color,
    getMoves,
  };
}
