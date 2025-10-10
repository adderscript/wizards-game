package main

import (
	"log"

	"wizards/src/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game.New()); err != nil {
		log.Fatal(err)
	}
}
