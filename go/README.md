# Dropon Chess Assignment (Go Version)

This assignment is a Go translation of the original TypeScript project. Your task is to complete several features of a basic chess engine using Go.

## Prerequisites

1. Install Go (version 1.22+ recommended)

Download Go from: <https://go.dev/dl/>

Verify installation:

    go version

## Running the Project

To run:

    go run .

To build:

    go build -o chess
    ./chess

This will initialize a board, place a white pawn at (1,0), and print the board every 5 seconds.

## Project Structure

    .
    ├── board.go        # Board logic, printing, FEN loading (to implement)
    ├── piece.go        # Piece interface + BasePiece
    ├── pawn.go
    ├── rook.go
    ├── knight.go
    ├── bishop.go
    ├── queen.go
    ├── king.go
    ├── main.go         # Game loop, signals
    └── README.md

### Overview

- Board: A 2D slice [][]Piece with helper functions
- Piece: Interface implemented by all concrete piece types
- Concrete pieces: Pawn, Rook, Knight, Bishop, Queen, King

## Assignment

### 1a. FEN Strings

In board.go you will find:

    func (b *Board) ReadFEN(fenString string)

Your task:

1. Parse a FEN string.
2. Place correct pieces (Pawn, Rook, etc.) on correct squares.
3. Use correct colors based on FEN symbols.
4. Support example FEN strings from <https://www.chess.com/terms/fen-chess>

### 1b. FEN Strings - Networking

In `../server/main.go` resides a simple server. The server serves one endpoint `fenstrings`, where you can fetch multiple fen strings. The server supports pagination.

Running the server

```bash
make fenserver
```

A simple request:

```
 curl "localhost:3000/fenstrings"
```

The assignment is simple.

1. Write a http client that's able to get the fen strings
2. Instead of fetching it in one time, implement a way to fetch the fenstrings per batch of 500 strings. These requests need to happen concurrently.
3. Some FEN strings are invalid. Find the FEN strings that are invalid in the response.

You can use query parameters to paginate the response:

```
 curl "localhost:3000/fenstrings?from=100&to=500"
```

### 2. Moves

Each piece has an unimplemented Moves method:

    func (p *Pawn) Moves() [][]int { return [][]int{} }

Implement movement rules for all pieces:

- Respect board boundaries
- Allow captures where appropriate
- Special moves optional (castling, en passant, etc.)

You may choose how pieces access the board.

### 3. Players

Add a Player concept:

- Each player has a color (white/black)
- Only the player whose turn it is may move
- After a valid move, switch turns

Add a move submission mechanism (stdin-driven or hardcoded examples).

### 4. Game Over

Define and detect end-of-game conditions.

Minimum:

- End game when a king is captured

Optional:

- Check / checkmate logic
- Stalemate
- Draw rules

Decide what happens when the game ends (print message, stop loop, etc.).

### 5. Extra

#### 5.1 Server/Client Model

If this becomes a server:

- Server holds board state
- Clients request legal moves
- Clients submit moves
- Clients receive updates

Questions:

- Would you use REST, WebSockets, or gRPC?
- How would clients represent moves?
- Why your chosen communication method?

Example JSON move:

    { "from": [4,1], "to": [4,3] }

#### 5.2 Timers

If each player has a chess clock:

- When time reaches zero, that player loses

Questions:

- How to model timers on the server?
- How to synchronize timers with clients?
- How to handle latency?
- Techniques such as server-authoritative time and periodic sync can be discussed.

This assignment tests your understanding of Go, game logic, state management, and system design. Feel free to extend the project as needed, as long as the base behavior continues to work with `go run .`.
