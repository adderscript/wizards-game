package main

import (
	"log"

	"wizards/src"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(src.NewGame()); err != nil {
		log.Fatal(err)
	}
}
