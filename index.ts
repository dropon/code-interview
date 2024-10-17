import { sleep } from "bun";
import { newBoard } from "./board";
import { newPawn } from "./piece/pawn";

const shutdown = () => {
  console.log("Shutting down...");
};

process.on("SIGKILL", shutdown);
process.on("SIGTERM", shutdown);

const game = newBoard(8);
game.placePiece(newPawn(1, 0, "white"));

while (true) {
  game.print();
  await sleep(5 * 1000);
}
