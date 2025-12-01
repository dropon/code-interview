# Assignment Dropon


## Prerequisites

1. Install bun 

```bash
# Linux/macOS:
curl -fsSL https://bun.sh/install | bash

# Windows:
powershell -c "irm bun.sh/install.ps1 | iex"
```

## Running the project

```bash
bun run index.ts
```

## Given

- An empty chess board with some methods to move pieces and print the board
- A `piece` folder that contains all the different pieces in chess. 
The `piece.ts` file contains structure of a piece.

## Assignment

### 1. FEN strings

There's a function inside `board.ts` called `readFEN`. This function reads
a FEN string and places the pieces on the board according to that position.

More info can be found here: https://www.chess.com/terms/fen-chess

1. Read a FEN string and place the corresponding pieces on the board. There are
a few examples on the page above and they should work.

### 2. Moves 

Each piece has a `getMoves` method, which checks what moves are valid
in the current position. Implement the `getMoves` function for every 
piece.

### 3. Players

Of course, now there are no players. Add players into the mix. Each player
should be able to play a move when it's their turn.


### 4. Game over

When is the game over? What should happen when the game is over?


### 5. Extra

1. Let's say you want to make a server out of this project. The internal
state is kept on the server, players can check valid moves and they of course also
can play moves. Let's say the board building part is done. How would
you communicate between server/client and why?

2. Let's say you want to implement a timer for both players. This typically is 
a clock of a certain amount of secods. When it expires, the game is also
lost. Is it difficult to synchronize the clocks over a network? Are
you familiar with ways you could achieve that?

