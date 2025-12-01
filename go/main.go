// file: main.go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dropon/code-interview/board"
	"github.com/dropon/code-interview/piece"
)

func shutdown() {
	fmt.Println("Shutting down...")
}

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-signals
		shutdown()
		os.Exit(0)
	}()

	game := board.NewBoard(8)
	game.PlacePiece(piece.NewPawn(1, 0, piece.ColorWhite))

	for {
		game.Print()
		time.Sleep(5 * time.Second)
	}
}
