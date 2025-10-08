package engine

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
	GetTag() string
	Update()
	Draw(screen *ebiten.Image)
}
